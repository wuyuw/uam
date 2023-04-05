package auth

import "time"

const TableAuthLocal = "auth_local"

type AuthLocal struct {
	Id         int64     `gorm:"column:id;primaryKey;autoIncrement"`
	Uid        int64     `gorm:"column:uid"`      // uid
	Username   string    `gorm:"column:username"` // 用户名
	Password   string    `gorm:"column:password"` // 密码
	Salt       string    `gorm:"column:salt"`     // 盐
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime"`
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime"`
}

func (AuthLocal) TableName() string {
	return TableAuthLocal
}
