// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameFacilityCategory = "facility_category"

// FacilityCategory mapped from table <facility_category>
type FacilityCategory struct {
	CategoryInstanceID string `gorm:"column:category_instance_id;primaryKey" json:"category_instance_id"`
	FacilityInstanceID string `gorm:"column:facility_instance_id;primaryKey" json:"facility_instance_id"`
}

// TableName FacilityCategory's table name
func (*FacilityCategory) TableName() string {
	return TableNameFacilityCategory
}
