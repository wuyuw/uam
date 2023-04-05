package uamclient

import "time"

const TableUamClient = "uam_client"

type UamClient struct {
	Id         int64     `gorm:"column:id;primaryKey;autoIncrement"`
	Name       string    `gorm:"column:name"`
	AppCode    string    `gorm:"column:app_code"`
	PrivateKey string    `gorm:"column:private_key"`
	Department string    `gorm:"column:department"` // 所属部门
	Maintainer string    `gorm:"column:maintainer"` // 对接人
	Status     int64     `gorm:"column:status"`     //  '状态: 0-正常, 1-禁用, 2-删除'
	Type       int64     `gorm:"column:type"`       // 客户端类型: 1-普通客户端, 2-系统客户端
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime"`
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime"`
}

func (UamClient) TableName() string {
	return TableUamClient
}
