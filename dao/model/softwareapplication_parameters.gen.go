// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSoftwareapplicationParameter = "softwareapplication_parameters"

// SoftwareapplicationParameter mapped from table <softwareapplication_parameters>
type SoftwareapplicationParameter struct {
	SoftwareapplicationInstanceID string `gorm:"column:softwareapplication_instance_id;primaryKey" json:"softwareapplication_instance_id"`
	ParameterInstanceID           string `gorm:"column:parameter_instance_id;primaryKey" json:"parameter_instance_id"`
}

// TableName SoftwareapplicationParameter's table name
func (*SoftwareapplicationParameter) TableName() string {
	return TableNameSoftwareapplicationParameter
}
