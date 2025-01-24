// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameFacility = "facility"

// Facility mapped from table <facility>
type Facility struct {
	InstanceID  string `gorm:"column:instance_id;primaryKey" json:"instance_id"`
	MetaID      string `gorm:"column:meta_id" json:"meta_id"`
	UID         string `gorm:"column:uid" json:"uid"`
	VersionID   string `gorm:"column:version_id" json:"version_id"`
	Identifier  string `gorm:"column:identifier" json:"identifier"`
	Description string `gorm:"column:description" json:"description"`
	Title       string `gorm:"column:title" json:"title"`
	Type        string `gorm:"column:type" json:"type"`
	Keywords    string `gorm:"column:keywords" json:"keywords"`
}

// TableName Facility's table name
func (*Facility) TableName() string {
	return TableNameFacility
}
