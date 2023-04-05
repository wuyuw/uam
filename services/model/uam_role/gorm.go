package uamrole

import "time"

const TableUamRole = "uam_role"

type UamRole struct {
	Id         int64     `gorm:"column:id;primaryKey;autoIncrement"`
	ClientId   int64     `gorm:"column:client_id"`
	Name       string    `gorm:"column:name"` // 名称
	Desc       string    `gorm:"column:desc"`
	Editable   int64     `gorm:"column:editable"` // 是否允许后台编辑
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime"`
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime"`
}

func (UamRole) TableName() string {
	return TableUamRole
}
