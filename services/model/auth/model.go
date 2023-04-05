package auth

import (
	"context"
	"uam/services/model"

	"gorm.io/gorm"
)

type AuthLocalModel struct {
	table string
	db    *gorm.DB
}

func NewAuthLocalModel(db *gorm.DB) *AuthLocalModel {
	return &AuthLocalModel{
		db:    db,
		table: TableAuthLocal,
	}
}

// InsertOne 添加一条记录
func (m *AuthLocalModel) InsertOne(ctx context.Context, db *gorm.DB, authLocal *AuthLocal) error {
	return db.Create(authLocal).Error
}

// FindOneByUsername 根据Username查询用户信息
func (m *AuthLocalModel) FindOneByUsername(ctx context.Context, username string) (*AuthLocal, error) {
	var (
		err       error
		authLocal AuthLocal
	)
	db := m.db.Table(m.table)
	err = db.Where("`username` = ?", username).First(&authLocal).Error
	switch err {
	case nil:
		return &authLocal, nil
	case gorm.ErrRecordNotFound:
		return nil, model.ErrNotFound
	default:
		return nil, err
	}
}
