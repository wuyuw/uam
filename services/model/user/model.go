package user

import (
	"context"
	"fmt"
	"uam/services/model"

	"gorm.io/gorm"
)

type UserModel struct {
	table string
	db    *gorm.DB
}

func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{
		db:    db,
		table: TableUser,
	}
}

func (m *UserModel) InsertOne(ctx context.Context, db *gorm.DB, user *User) error {
	return db.Create(user).Error
}

// FindOneByUid 根据uid查询用户信息
func (m *UserModel) FindOneByUid(ctx context.Context, uid int64) (*User, error) {
	var (
		err  error
		user User
	)
	db := m.db.Table(m.table)
	err = db.Where("`uid` = ?", uid).First(&user).Error
	switch err {
	case nil:
		return &user, nil
	case gorm.ErrRecordNotFound:
		return nil, model.ErrNotFound
	default:
		return nil, err
	}
}

// FindUidsBySearch 通过用户名和昵称搜索用户
func (m *UserModel) FindUidsBySearch(ctx context.Context, search string) ([]int64, error) {
	var (
		err     error
		records []map[string]interface{}
	)
	db := m.db.Table(m.table)
	if search == "" {
		return nil, nil
	}
	searchArgs := fmt.Sprintf("%%%s%%", search)
	if err = db.Select("`uid`").Where(" nickname LIKE ?",
		searchArgs).Find(&records).Error; err != nil {
		return nil, err
	}
	uids := make([]int64, len(records))
	for i, r := range records {
		uids[i] = int64(r["uid"].(int32))
	}
	return uids, nil
}

// FindByUids 通过uid列表查询用户
func (m *UserModel) FindByUids(ctx context.Context, uids []int64) ([]*User, error) {
	var (
		err   error
		users []*User
	)
	if len(uids) == 0 {
		return nil, nil
	}
	db := m.db.Table(m.table)
	if err = db.Where("`uid` IN ?", uids).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
