// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameIdentifier = "identifier"

// Identifier mapped from table <identifier>
type Identifier struct {
	InstanceID string `gorm:"column:instance_id;primaryKey" json:"instance_id"`
	MetaID     string `gorm:"column:meta_id" json:"meta_id"`
	UID        string `gorm:"column:uid" json:"uid"`
	VersionID  string `gorm:"column:version_id" json:"version_id"`
	Type       string `gorm:"column:type" json:"type"`
	Value      string `gorm:"column:value" json:"value"`
}

// TableName Identifier's table name
func (*Identifier) TableName() string {
	return TableNameIdentifier
}
