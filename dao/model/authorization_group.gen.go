// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameAuthorizationGroup = "authorization_group"

// AuthorizationGroup mapped from table <authorization_group>
type AuthorizationGroup struct {
	ID      string `gorm:"column:id;primaryKey" json:"id"`
	GroupID string `gorm:"column:group_id" json:"group_id"`
	MetaID  string `gorm:"column:meta_id" json:"meta_id"`
}

// TableName AuthorizationGroup's table name
func (*AuthorizationGroup) TableName() string {
	return TableNameAuthorizationGroup
}
