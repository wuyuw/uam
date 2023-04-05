package uamgroup

import (
	"context"
	"uam/services/model"

	"gorm.io/gorm"
)

type GroupModel struct {
	table string
	db    *gorm.DB
}

func NewGroupModel(db *gorm.DB) *GroupModel {
	return &GroupModel{
		db:    db,
		table: TableUamGroup,
	}
}

// FindOne 根据id查询组
func (m *GroupModel) FindOne(ctx context.Context, id int64) (*UamGroup, error) {
	var (
		err  error
		role UamGroup
	)
	db := m.db.Table(m.table)
	err = db.First(&role, id).Error
	switch err {
	case nil:
		return &role, nil
	case gorm.ErrRecordNotFound:
		return nil, model.ErrNotFound
	default:
		return nil, err
	}
}

// FindOneByName 根据name查询客户端组
func (m *GroupModel) FindOneByName(ctx context.Context, clientId int64, name string) (*UamGroup, error) {
	var (
		err   error
		group UamGroup
	)
	db := m.db.Table(m.table)
	err = db.Where("`client_id` = ? AND `name` = ?", clientId, name).First(&group).Error
	switch err {
	case nil:
		return &group, nil
	case gorm.ErrRecordNotFound:
		return nil, model.ErrNotFound
	default:
		return nil, err
	}
}

// FindAll 查询客户端所有组
func (m *GroupModel) FindAll(ctx context.Context, clientId int64) ([]*UamGroup, error) {
	var (
		err    error
		groups []*UamGroup
	)
	db := m.db.Table(m.table)
	if err = db.Where("`client_id` = ?", clientId).Find(&groups).Error; err != nil {
		return nil, err
	}
	return groups, nil
}

// FindByIds 根据ID列表查询组
func (m *GroupModel) FindByIds(ctx context.Context, clientId int64, ids []int64) ([]*UamGroup, error) {
	var (
		err    error
		groups []*UamGroup
	)
	db := m.db.Table(m.table)
	if err = db.Where("`client_id` = ? AND `id` IN ?", clientId, ids).Find(&groups).Error; err != nil {
		return nil, err
	}
	return groups, nil
}

// FindByEditable 根据可编辑字段筛选组
func (m *GroupModel) FindByEditable(ctx context.Context, clientId int64, editable string) ([]*UamGroup, error) {
	var (
		err    error
		groups []*UamGroup
	)
	db := m.db.Table(m.table)
	db = db.Where("`client_id` = ?", clientId)
	if editable != "" {
		db = db.Where("`editable` = ?", editable)
	}
	if err = db.Find(&groups).Error; err != nil {
		return nil, err
	}
	return groups, nil
}

// InsertOne 添加组
func (m *GroupModel) InsertOne(ctx context.Context, db *gorm.DB, group *UamGroup) error {
	if db == nil {
		db = m.db.Table(m.table)
	}
	return db.Create(group).Error
}

// Update 更新组
func (m *GroupModel) Update(ctx context.Context, group *UamGroup) error {
	db := m.db.Table(m.table)
	return db.Save(group).Error
}

// Delete 删除组
func (m *GroupModel) Delete(ctx context.Context, db *gorm.DB, id int64) error {
	return db.Delete(&UamGroup{}, id).Error
}
