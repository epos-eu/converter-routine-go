// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNamePublicationCategory = "publication_category"

// PublicationCategory mapped from table <publication_category>
type PublicationCategory struct {
	CategoryInstanceID    string `gorm:"column:category_instance_id;primaryKey" json:"category_instance_id"`
	PublicationInstanceID string `gorm:"column:publication_instance_id;primaryKey" json:"publication_instance_id"`
}

// TableName PublicationCategory's table name
func (*PublicationCategory) TableName() string {
	return TableNamePublicationCategory
}
