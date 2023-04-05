package uamrel

import "time"

const (
	TableRelUserGroup = "rel_user_group"
	TableRelUserRole  = "rel_user_role"
	TableRelGroupRole = "rel_group_role"
	TableRelRolePerm  = "rel_role_permission"
)

type RelUserGroup struct {
	Id         int64     `gorm:"column:id;primaryKey;autoIncrement"`
	ClientId   int64     `gorm:"column:client_id"`
	Uid        int64     `gorm:"column:uid"`
	GroupId    int64     `gorm:"column:group_id"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime"`
}

func (RelUserGroup) TableName() string {
	return TableRelUserGroup
}

type RelUserRole struct {
	Id         int64     `gorm:"column:id;primaryKey;autoIncrement"`
	ClientId   int64     `gorm:"column:client_id"`
	Uid        int64     `gorm:"column:uid"`
	RoleId     int64     `gorm:"column:role_id"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime"`
}

func (RelUserRole) TableName() string {
	return TableRelUserRole
}

type RelGroupRole struct {
	Id         int64     `gorm:"column:id;primaryKey;autoIncrement"`
	GroupId    int64     `gorm:"column:group_id"`
	RoleId     int64     `gorm:"column:role_id"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime"`
}

func (RelGroupRole) TableName() string {
	return TableRelGroupRole
}

type RelRolePermission struct {
	Id           int64     `gorm:"column:id;primaryKey;autoIncrement"`
	RoleId       int64     `gorm:"column:role_id"`
	PermissionId int64     `gorm:"column:permission_id"`
	CreateTime   time.Time `gorm:"column:create_time;autoCreateTime"`
}

func (RelRolePermission) TableName() string {
	return TableRelRolePerm
}
