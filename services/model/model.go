package model

import (
	"database/sql"
	"uam/tools/constants"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var ErrNotFound = sql.ErrNoRows

type PageListResp struct {
	Page     int64
	PageSize int64
	Total    int64
	List     interface{}
}

func ExecInTx(db *gorm.DB, txFunc func(tx *gorm.DB) error) error {
	var err error
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err = txFunc(tx); err != nil {
		tx.Rollback()
		return errors.Wrap(err, constants.MsgDBErr)
	}
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return errors.Wrap(err, constants.MsgDBErr)
	}
	return nil
}
