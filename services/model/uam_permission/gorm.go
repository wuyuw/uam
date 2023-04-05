package uampermission

import "time"

const TableUamPermission = "uam_permission"

type UamPermission struct {
	Id         int64     `gorm:"column:id;primaryKey;autoIncrement"`
	ClientId   int64     `gorm:"column:client_id"`
	Type       string    `gorm:"column:type"` // 类型
	Key        string    `gorm:"column:key"`  // 唯一标识
	Name       string    `gorm:"column:name"` // 名称
	Desc       string    `gorm:"column:desc"`
	Editable   int64     `gorm:"column:editable"` // 是否允许后台编辑
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime"`
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime"`
}

func (UamPermission) TableName() string {
	return TableUamPermission
}
