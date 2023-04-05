package uamclient

import (
	"context"
	"uam/services/model"

	"gorm.io/gorm"
)

type ClientModel struct {
	table string
	db    *gorm.DB
}

func NewClientModel(db *gorm.DB) *ClientModel {
	return &ClientModel{
		db:    db,
		table: TableUamClient,
	}
}

// FindOne 根据id查询客户端
func (m *ClientModel) FindOne(ctx context.Context, id int64) (*UamClient, error) {
	var (
		err    error
		client UamClient
	)
	db := m.db.Table(m.table)
	err = db.First(&client, id).Error
	switch err {
	case nil:
		return &client, nil
	case gorm.ErrRecordNotFound:
		return nil, model.ErrNotFound
	default:
		return nil, err
	}
}

// FindAll 查询所有客户端
func (m *ClientModel) FindAll(ctx context.Context) ([]*UamClient, error) {
	var (
		err     error
		clients []*UamClient
	)
	db := m.db.Table(m.table)
	if err = db.Find(&clients).Error; err != nil {
		return nil, err
	}
	return clients, nil
}

func (m *ClientModel) FindByCodes(ctx context.Context, appCodes []string) ([]*UamClient, error) {
	var (
		err     error
		clients []*UamClient
	)
	if len(appCodes) == 0 {
		return nil, nil
	}
	db := m.db.Table(m.table)
	if err = db.Where("`app_code` IN ?", appCodes).Find(&clients).Error; err != nil {
		return nil, err
	}
	return clients, nil
}

// FindOneByAppCode 根据app_code查询客户端
func (m *ClientModel) FindOneByAppCode(ctx context.Context, appCode string) (*UamClient, error) {
	var (
		err    error
		client UamClient
	)
	db := m.db.Table(m.table)
	err = db.Where("`app_code` = ?", appCode).First(&client).Error
	switch err {
	case nil:
		return &client, nil
	case gorm.ErrRecordNotFound:
		return nil, model.ErrNotFound
	default:
		return nil, err
	}
}

// InsertOne 添加客户端
func (m *ClientModel) InsertOne(ctx context.Context, db *gorm.DB, client *UamClient) error {
	return db.Create(client).Error
}

// Update 更新客户端
func (m *ClientModel) Update(ctx context.Context, client *UamClient) error {
	db := m.db.Table(m.table)
	return db.Save(client).Error
}

// Delete 删除客户端
func (m *ClientModel) Delete(ctx context.Context, db *gorm.DB, id int64) error {
	return db.Delete(&UamClient{}, id).Error
}

// FindSysClient 查找系统客户端
func (m *ClientModel) FindSysClient(ctx context.Context) (*UamClient, error) {
	var (
		err    error
		client UamClient
	)
	db := m.db.Table(m.table)
	err = db.Where("`type` = 2").First(&client).Error
	switch err {
	case nil:
		return &client, nil
	case gorm.ErrRecordNotFound:
		return nil, model.ErrNotFound
	default:
		return nil, err
	}
}
