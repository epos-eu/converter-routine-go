// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameOrganizationAffiliation = "organization_affiliation"

// OrganizationAffiliation mapped from table <organization_affiliation>
type OrganizationAffiliation struct {
	PersonInstanceID       string `gorm:"column:person_instance_id;primaryKey" json:"person_instance_id"`
	OrganizationInstanceID string `gorm:"column:organization_instance_id;primaryKey" json:"organization_instance_id"`
}

// TableName OrganizationAffiliation's table name
func (*OrganizationAffiliation) TableName() string {
	return TableNameOrganizationAffiliation
}
