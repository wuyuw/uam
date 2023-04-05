package uampermission

import (
	"context"
	"fmt"
	"uam/services/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PermissionModel struct {
	table string
	db    *gorm.DB
}

func NewPermissionModel(db *gorm.DB) *PermissionModel {
	return &PermissionModel{
		db:    db,
		table: TableUamPermission,
	}
}

// FindOne 根据id查询权限
func (m *PermissionModel) FindOne(ctx context.Context, id int64) (*UamPermission, error) {
	var (
		err  error
		perm UamPermission
	)
	db := m.db.Table(m.table)
	err = db.First(&perm, id).Error
	switch err {
	case nil:
		return &perm, nil
	case gorm.ErrRecordNotFound:
		return nil, model.ErrNotFound
	default:
		return nil, err
	}
}

// FindOneByKey 根据key查询权限
func (m *PermissionModel) FindOneByKey(ctx context.Context, clientId int64, key string) (*UamPermission, error) {
	var (
		err  error
		perm UamPermission
	)
	db := m.db.Table(m.table)
	err = db.Where("`client_id` = ? AND `key` = ?", clientId, key).First(&perm).Error
	switch err {
	case nil:
		return &perm, nil
	case gorm.ErrRecordNotFound:
		return nil, model.ErrNotFound
	default:
		return nil, err
	}
}

// FindTypeOptions 获取权限类型可选项
func (m *PermissionModel) FindTypeOptions(ctx context.Context, clientId int64) ([]string, error) {
	var (
		err         error
		records     []map[string]interface{}
		typeOptions []string
	)
	typeOptions = make([]string, 0)
	db := m.db.Table(m.table)
	if err = db.Where("`client_id` = ?", clientId).Distinct("type").Find(&records).Error; err != nil {
		return nil, err
	}
	for _, record := range records {
		typeOptions = append(typeOptions, record["type"].(string))
	}
	return typeOptions, nil
}

// FindByType 查询客户端所有权限
func (m *PermissionModel) FindByType(ctx context.Context, clientId int64, permType string) ([]*UamPermission, error) {
	var (
		err      error
		permList []*UamPermission
	)
	db := m.db.Table(m.table)
	db = db.Where("`client_id` = ?", clientId)
	if permType != "" {
		db.Where("`type` = ?", permType)
	}
	if err = db.Order("`type`").Find(&permList).Error; err != nil {
		return nil, err
	}
	return permList, nil
}

// FindPageList 获取权限分页列表
func (m *PermissionModel) FindPageList(ctx context.Context,
	page, pageSize int64, clientId int64, permType, editable, search string) (*model.PageListResp, error) {
	var (
		err      error
		total    int64
		permList []*UamPermission
	)
	db := m.db.Table(m.table)
	db = db.Where("`client_id` = ?", clientId)
	if permType != "" {
		db = db.Where("`type` = ?", permType)
	}
	if editable != "" {
		db = db.Where("`editable` = ?", editable)
	}
	if search != "" {
		searchArgs := fmt.Sprintf("%%%s%%", search)
		db = db.Where("`key` LIKE ? or `name` LIKE ? or `desc` LIKE ?", searchArgs, searchArgs, searchArgs)
	}
	db = db.Order("`type`")
	if err = db.Count(&total).Error; err != nil {
		return nil, err
	}
	offset := pageSize * (page - 1)
	if total <= offset {
		page = 1
		offset = 0
	}
	if err = db.Limit(int(pageSize)).Offset(int(offset)).Find(&permList).Error; err != nil {
		return nil, err
	}
	resp := &model.PageListResp{
		Page:     page,
		PageSize: pageSize,
		Total:    total,
		List:     permList,
	}
	return resp, nil
}

// FindByIds 根据ID列表查询权限
func (m *PermissionModel) FindByIds(ctx context.Context, ids []int64) ([]*UamPermission, error) {
	var (
		err      error
		permList []*UamPermission
	)
	if len(ids) == 0 {
		return permList, nil
	}
	db := m.db.Table(m.table)
	if err = db.Where("`id` IN ?", ids).Find(&permList).Error; err != nil {
		return nil, err
	}
	return permList, nil
}

// FindByKeys 根据keys列表查询权限
func (m *PermissionModel) FindByKeys(ctx context.Context, clientId int64, keys []string) ([]*UamPermission, error) {
	var (
		err      error
		permList []*UamPermission
	)
	if len(keys) == 0 {
		return permList, nil
	}
	db := m.db.Table(m.table)
	if err = db.Where("`client_id` = ? AND `key` IN ?", clientId, keys).Find(&permList).Error; err != nil {
		return nil, err
	}
	return permList, nil
}

// FindKeysByIds 根据ID列表获取key列表
func (m *PermissionModel) FindKeysByIds(ctx context.Context, clientId int64, ids []int64) ([]string, error) {
	var (
		err     error
		records []map[string]interface{}
	)
	if len(ids) == 0 {
		return nil, nil
	}
	db := m.db.Table(m.table)
	if err = db.Where("`client_id` = ? AND `id` IN ?", clientId, ids).Distinct("key").Find(&records).Error; err != nil {
		return nil, err
	}
	keys := make([]string, len(records))
	for i, record := range records {
		keys[i] = record["key"].(string)
	}
	return keys, nil
}

// InsertOne 添加权限
func (m *PermissionModel) InsertOne(ctx context.Context, db *gorm.DB, perm *UamPermission) error {
	return db.Create(perm).Error
}

// Update 更新权限
func (m *PermissionModel) Update(ctx context.Context, perm *UamPermission) error {
	db := m.db.Table(m.table)
	return db.Save(perm).Error
}

// Delete 删除权限
func (m *PermissionModel) Delete(ctx context.Context, db *gorm.DB, id int64) error {
	return db.Delete(&UamPermission{}, id).Error
}

// UpsertPermList 批量添加权限
func (m *PermissionModel) UpsertPermList(ctx context.Context, permList []*UamPermission) error {
	if len(permList) == 0 {
		return nil
	}
	db := m.db.Table(m.table)
	return db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "client_id"}, {Name: "key"}},
		DoUpdates: clause.AssignmentColumns([]string{"type", "name", "desc", "editable"}),
	}).CreateInBatches(permList, 100).Error
}
