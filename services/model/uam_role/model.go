package uamrole

import (
	"context"
	"uam/services/model"

	"gorm.io/gorm"
)

type RoleModel struct {
	table string
	db    *gorm.DB
}

func NewRoleModel(db *gorm.DB) *RoleModel {
	return &RoleModel{
		db:    db,
		table: TableUamRole,
	}
}

// FindOne 根据id查询角色
func (m *RoleModel) FindOne(ctx context.Context, id int64) (*UamRole, error) {
	var (
		err  error
		role UamRole
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

// FindOneByName 根据name查询客户端角色
func (m *RoleModel) FindOneByName(ctx context.Context, clientId int64, name string) (*UamRole, error) {
	var (
		err  error
		role UamRole
	)
	db := m.db.Table(m.table)
	err = db.Where("`client_id` = ? AND `name` = ?", clientId, name).First(&role).Error
	switch err {
	case nil:
		return &role, nil
	case gorm.ErrRecordNotFound:
		return nil, model.ErrNotFound
	default:
		return nil, err
	}
}

// FindByEditable 根据是否可编辑筛选角色
func (m *RoleModel) FindByEditable(ctx context.Context, clientId int64, editable string) ([]*UamRole, error) {
	var (
		err   error
		roles []*UamRole
	)
	db := m.db.Table(m.table)
	db = db.Where("`client_id` = ?", clientId)
	if editable != "" {
		db = db.Where("`editable` = ?", editable)
	}
	if err = db.Order("editable").Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

// FindListByClientId 查询客户端所有角色
func (m *RoleModel) FindListByClientId(ctx context.Context, clientId int64) ([]*UamRole, error) {
	var (
		err   error
		roles []*UamRole
	)
	db := m.db.Table(m.table)
	if err = db.Where("`client_id` = ?", clientId).Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

// FindListByIds 根据角色ID列表查询角色
func (m *RoleModel) FindListByIds(ctx context.Context, clientId int64, roleIds []int64) ([]*UamRole, error) {
	var (
		err   error
		roles []*UamRole
	)
	if len(roleIds) == 0 {
		return nil, nil
	}
	db := m.db.Table(m.table)
	if err = db.Where("`client_id` = ? AND `id` IN ?", clientId, roleIds).Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

// InsertOne 添加角色
func (m *RoleModel) InsertOne(ctx context.Context, db *gorm.DB, role *UamRole) error {
	if db == nil {
		db = m.db.Table(m.table)
	}
	return db.Create(role).Error
}

// Update 更新角色
func (m *RoleModel) Update(ctx context.Context, role *UamRole) error {
	db := m.db.Table(m.table)
	return db.Save(role).Error
}

// Delete 删除角色
func (m *RoleModel) Delete(ctx context.Context, db *gorm.DB, id int64) error {
	return db.Delete(&UamRole{}, id).Error
}
