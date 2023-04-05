package user

import "time"

const TableUser = "user"

type User struct {
	Id         int64     `gorm:"column:id;primaryKey;autoIncrement"`
	Uid        int64     `gorm:"column:uid"`      // uid
	Nickname   string    `gorm:"column:nickname"` // 花名
	Email      string    `gorm:"column:email"`    // 邮箱
	Phone      string    `gorm:"column:phone"`    // 手机
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime"`
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime"`
}

func (User) TableName() string {
	return TableUser
}
