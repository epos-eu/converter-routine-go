package connection

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// dbPools is a map from an environment variable (that holds a DSN)
// to a slice of *gorm.DB connections. This allows you to have multiple
// distinct DSNs/connection sets under different environment variables.
var dbPools = make(map[string][]*gorm.DB)

// Protect dbPools with a mutex if multiple goroutines might race to init
var mu sync.Mutex

// ConnectMetadata is a thin wrapper that uses the manager to fetch
// a connection for METADATA_CATALOGUE_CONNECTION_STRING.
func ConnectMetadata() (*gorm.DB, error) {
	return connectManager("METADATA_CATALOGUE_CONNECTION_STRING")
}

// ConnectConverter is a thin wrapper that uses the manager to fetch
// a connection for CONVERTER_CATALOGUE_CONNECTION_STRING.
func ConnectConverter() (*gorm.DB, error) {
	return connectManager("CONVERTER_CATALOGUE_CONNECTION_STRING")
}

// connectManager checks if we have a pool of *gorm.DB for the given
// environment variable. If not, it initializes it, then returns a *gorm.DB.
func connectManager(envVar string) (*gorm.DB, error) {
	// check if we already have a pool
	if _, exists := dbPools[envVar]; !exists || len(dbPools[envVar]) == 0 {
		// initialize a new pool of connections
		err := initializePool(envVar)
		if err != nil {
			return nil, fmt.Errorf("initialization error: %w", err)
		}
		// check if that succeeded in creating any connections
		if len(dbPools[envVar]) == 0 {
			return nil, fmt.Errorf("no database connections available for %s", envVar)
		}
	}

	// At this point, dbPools[envVar] should have at least 1 *gorm.DB
	// Try each one in turn and return the first that is reachable.
	for _, db := range dbPools[envVar] {
		sqlDB, err := db.DB()
		if err != nil {
			log.Printf("Error getting underlying *sql.DB: %v", err)
			continue
		}

		// Check connectivity
		if err := sqlDB.Ping(); err != nil {
			log.Printf("Failed to ping database for env=%s: %v", envVar, err)
			continue
		}

		// Return the first that works
		return db, nil
	}

	return nil, fmt.Errorf("all database hosts are unreachable for %s", envVar)
}

// initializePool reads the DSN from envVar, parses out the hosts, sets up
// multiple connections (one per host) and stores them in dbPools[envVar].
func initializePool(envVar string) error {
	hosts, params, err := parseMultiHostDSN(envVar)
	if err != nil {
		return fmt.Errorf("failed to parse DSN for %s: %w", envVar, err)
	}

	// GORM logger config
	logConfig := logger.Config{
		SlowThreshold:             time.Second,
		LogLevel:                  logger.Error,
		IgnoreRecordNotFoundError: false,
	}

	// Make a slice to hold the *gorm.DB for each host
	newDbs := make([]*gorm.DB, 0, len(hosts))

	for _, host := range hosts {
		currentDSN := fmt.Sprintf("postgresql://%s/%s", host, params)

		db, err := gorm.Open(postgres.New(postgres.Config{
			DriverName: "pgx",
			DSN:        currentDSN,
		}), &gorm.Config{
			Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logConfig),
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   "",
				SingularTable: true,
			},
		})
		if err != nil {
			log.Printf("Failed to connect to host %s (env=%s): %v", host, envVar, err)
			continue
		}

		newDbs = append(newDbs, db)
	}

	if len(newDbs) == 0 {
		return fmt.Errorf("failed to initialize any DB connections for %s", envVar)
	}

	// store in global map
	dbPools[envVar] = newDbs

	return nil
}

// parseMultiHostDSN fetches the DSN from the given envVar and
// splits it into (hosts, params).
func parseMultiHostDSN(envVar string) ([]string, string, error) {
	dsn, ok := os.LookupEnv(envVar)
	log.Printf("%s: %s", envVar, dsn)
	if !ok {
		return nil, "", fmt.Errorf("%s is not set", envVar)
	}

	// Remove "jdbc:" prefix if present
	dsn = strings.Replace(dsn, "jdbc:", "", 1)
	log.Println("Cleaned DSN (jdbc prefix removed):", dsn)

	// Remove unsupported parameters like targetServerType & loadBalanceHosts
	re := regexp.MustCompile(`(&?(targetServerType|loadBalanceHosts)=[^&]+)`)
	dsn = re.ReplaceAllString(dsn, "")
	log.Println("Cleaned DSN (unsupported parameters removed):", dsn)

	// Clean up trailing "?" or "&"
	dsn = regexp.MustCompile(`[?&]$`).ReplaceAllString(dsn, "")
	log.Println("Cleaned DSN (trailing ? or & removed):", dsn)

	// Must contain "//"
	hostStart := strings.Index(dsn, "//")
	if hostStart == -1 {
		return nil, "", fmt.Errorf("invalid connection string format: missing '//'")
	}

	// Extract everything after `//` (hosts and params)
	hostsAndParams := dsn[hostStart+2:]
	splitIndex := strings.Index(hostsAndParams, "/")
	if splitIndex == -1 {
		return nil, "", fmt.Errorf("invalid connection string format: missing '/' after hosts")
	}

	hosts := hostsAndParams[:splitIndex]
	params := hostsAndParams[splitIndex+1:]

	hostList := strings.Split(hosts, ",")

	log.Printf("Parsed Hosts from %s: %v", envVar, hostList)
	log.Printf("Connection Params from %s: %s", envVar, params)

	return hostList, params, nil
}
