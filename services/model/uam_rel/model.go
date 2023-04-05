package uamrel

import (
	"context"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RelModel struct {
	tableUserGroup string
	tableUserRole  string
	tableGroupRole string
	tableRolePerm  string
	db             *gorm.DB
}

func NewRelModel(db *gorm.DB) *RelModel {
	return &RelModel{
		db:             db,
		tableUserGroup: TableRelUserGroup,
		tableUserRole:  TableRelUserRole,
		tableGroupRole: TableRelGroupRole,
		tableRolePerm:  TableRelRolePerm,
	}
}

// FindUidsByGroupId 根据组ID查询组内所有用户UID列表
func (m *RelModel) FindUidsByGroupId(ctx context.Context, clientId, groupId int64) ([]int64, error) {
	var (
		err     error
		records []map[string]interface{}
	)
	if groupId == 0 {
		return nil, nil
	}
	db := m.db.Table(m.tableUserGroup)
	if err = db.Where("`client_id` = ? AND `group_id` = ?",
		clientId, groupId).Distinct("`uid`").Find(&records).Error; err != nil {
		return nil, err
	}
	uids := make([]int64, len(records))
	for i, r := range records {
		uids[i] = int64(r["uid"].(int32))
	}
	return uids, err
}

// FindUidsByRoleId 根据角色ID查询组内所有用户UID列表
func (m *RelModel) FindUidsByRoleId(ctx context.Context, clientId, roleId int64) ([]int64, error) {
	var (
		err     error
		records []map[string]interface{}
	)
	if roleId == 0 {
		return nil, nil
	}
	db := m.db.Table(m.tableUserRole)
	if err = db.Where("`client_id` = ? AND `role_id` = ?",
		clientId, roleId).Distinct("`uid`").Find(&records).Error; err != nil {
		return nil, err
	}
	uids := make([]int64, len(records))
	for i, r := range records {
		uids[i] = int64(r["uid"].(int32))
	}
	return uids, err
}

// FindGroupIdsByUid 查询用户关联组ID列表
func (m *RelModel) FindGroupIdsByUid(ctx context.Context, clientId, uid int64) ([]int64, error) {
	var (
		err     error
		records []RelUserGroup
	)
	db := m.db.Table(m.tableUserGroup)
	if err = db.Where("`client_id` = ? AND `uid` = ?", clientId, uid).Find(&records).Error; err != nil {
		return nil, err
	}
	groupIds := make([]int64, len(records))
	for i, r := range records {
		groupIds[i] = r.GroupId
	}
	return groupIds, nil
}

// FindRoleIdsByUid 查询用户关联角色ID列表
func (m *RelModel) FindRoleIdsByUid(ctx context.Context, clientId, uid int64) ([]int64, error) {
	var (
		err     error
		records []RelUserRole
	)
	db := m.db.Table(m.tableUserRole)
	if err = db.Where("`client_id` = ? AND `uid` = ?", clientId, uid).Find(&records).Error; err != nil {
		return nil, err
	}
	roleIds := make([]int64, len(records))
	for i, r := range records {
		roleIds[i] = r.RoleId
	}
	return roleIds, nil
}

// FindRoleIdsByGroupId 查询组关联的所有角色ID列表
func (m *RelModel) FindRoleIdsByGroupId(ctx context.Context, groupId int64) ([]int64, error) {
	var (
		err     error
		records []map[string]interface{}
	)
	db := m.db.Table(m.tableGroupRole)
	if err = db.Select("role_id").Where("`group_id` = ?", groupId).Find(&records).Error; err != nil {
		return nil, err
	}
	roleIds := make([]int64, len(records))
	for i, r := range records {
		roleIds[i] = int64(r["role_id"].(int32))
	}
	return roleIds, nil
}

// FindRoleIdsByGroupIds 查询多组关联的所有角色ID去重列表
func (m *RelModel) FindRoleIdsByGroupIds(ctx context.Context, groupIds []int64) ([]int64, error) {
	var (
		err     error
		records []map[string]interface{}
	)
	if len(groupIds) == 0 {
		return nil, nil
	}
	db := m.db.Table(m.tableGroupRole)
	if err = db.Where("`group_id` IN ?", groupIds).Distinct("role_id").Find(&records).Error; err != nil {
		return nil, err
	}
	roleIds := make([]int64, len(records))
	for i, r := range records {
		roleIds[i] = int64(r["role_id"].(int32))
	}
	return roleIds, nil
}

// FindPermIdsByRoleIds 查询角色关联所有权限
func (m *RelModel) FindPermIdsByRoleIds(ctx context.Context, roleIds []int64) ([]int64, error) {
	var (
		err     error
		records []map[string]interface{}
	)
	if len(roleIds) == 0 {
		return nil, nil
	}
	db := m.db.Table(m.tableRolePerm)
	if err = db.Where("`role_id` IN ?", roleIds).Distinct("permission_id").Find(&records).Error; err != nil {
		return nil, err
	}
	permIds := make([]int64, len(records))
	for i, r := range records {
		permIds[i] = int64(r["permission_id"].(int32))
	}
	return permIds, nil
}

// AddUserGroupByGroupIds 将用户添加至组
func (m *RelModel) AddUserGroupByGroupIds(ctx context.Context, db *gorm.DB, client_id, uid int64, groupIds []int64) error {
	if len(groupIds) == 0 {
		return nil
	}
	if db == nil {
		db = m.db.Table(m.tableUserGroup)
	}
	records := make([]*RelUserGroup, len(groupIds))
	for i, groupId := range groupIds {
		records[i] = &RelUserGroup{
			ClientId: client_id,
			Uid:      uid,
			GroupId:  groupId,
		}
	}
	return db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "uidx_user_group"}},
		DoNothing: true,
	}).Create(&records).Error

}

// RemoveUserGroupByGroupIds 将用户移出组
func (m *RelModel) RemoveUserGroupByGroupIds(ctx context.Context, db *gorm.DB, client_id, uid int64, groupIds []int64) error {
	if len(groupIds) == 0 {
		return nil
	}
	if db == nil {
		db = m.db.Table(m.tableUserGroup)
	}
	return db.Where("`client_id` = ? AND `uid` = ? AND `group_id` IN ?",
		client_id, uid, groupIds).Delete(&RelUserGroup{}).Error
}

// RemoveUserGroupByGroupId 移除组关联的所有用户关系
func (m *RelModel) RemoveUserGroupByGroupId(ctx context.Context, db *gorm.DB, groupId int64) error {
	return db.Where("`group_id` = ?", groupId).Delete(&RelUserGroup{}).Error
}

// AddUserRoleByRoleIds 为用户添加角色
func (m *RelModel) AddUserRoleByRoleIds(ctx context.Context, db *gorm.DB, client_id, uid int64, roleIds []int64) error {
	if len(roleIds) == 0 {
		return nil
	}
	if db == nil {
		db = m.db.Table(m.tableUserRole)
	}
	records := make([]*RelUserRole, len(roleIds))
	for i, roleId := range roleIds {
		records[i] = &RelUserRole{
			ClientId: client_id,
			Uid:      uid,
			RoleId:   roleId,
		}
	}
	return db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "uidx_user_role"}},
		DoNothing: true,
	}).Create(&records).Error
}

// RemoveUserRoleByRoleIds 将用户移出角色
func (m *RelModel) RemoveUserRoleByRoleIds(ctx context.Context, db *gorm.DB, clientId, uid int64, roleIds []int64) error {
	if len(roleIds) == 0 {
		return nil
	}
	if db == nil {
		db = m.db.Table(m.tableUserRole)
	}
	return db.Where("`client_id` = ? AND `uid` = ? AND `role_id` IN ?",
		clientId, uid, roleIds).Delete(&RelUserRole{}).Error
}

// RemoveUserRoleByRoleId 移除角色与所有用户的关联关系
func (m *RelModel) RemoveUserRoleByRoleId(ctx context.Context, db *gorm.DB, roleId int64) error {
	return db.Where("`role_id` = ?", roleId).Delete(&RelUserRole{}).Error
}

// AddGroupRoleByRoleIds 将角色添加至组
func (m *RelModel) AddGroupRoleByRoleIds(ctx context.Context, db *gorm.DB, groupId int64, roleIds []int64) error {
	if len(roleIds) == 0 {
		return nil
	}
	records := make([]*RelGroupRole, len(roleIds))
	for i, roleId := range roleIds {
		records[i] = &RelGroupRole{
			GroupId: groupId,
			RoleId:  roleId,
		}
	}
	return db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "uidx_group_role"}},
		DoNothing: true,
	}).Create(&records).Error
}

// RemoveGroupRoleByRoleIds 将角色移出组
func (m *RelModel) RemoveGroupRoleByRoleIds(ctx context.Context, db *gorm.DB, groupId int64, roleIds []int64) error {
	if len(roleIds) == 0 {
		return nil
	}
	return db.Where("`group_id` = ? AND `role_id` IN ?", groupId, roleIds).Delete(&RelGroupRole{}).Error
}

// RemoveGroupRoleByRoleId 将角色从所有关联组中移除
func (m *RelModel) RemoveGroupRoleByRoleId(ctx context.Context, db *gorm.DB, roleId int64) error {
	return db.Where("`role_id` = ?", roleId).Delete(&RelGroupRole{}).Error
}

// RemoveGroupRoleByGroupId 将组关联的所有角色移除
func (m *RelModel) RemoveGroupRoleByGroupId(ctx context.Context, db *gorm.DB, groupId int64) error {
	return db.Where("`group_id` = ?", groupId).Delete(&RelGroupRole{}).Error
}

// AddRolePermByPermIds 将权限关联至角色
func (m *RelModel) AddRolePermByPermIds(ctx context.Context, db *gorm.DB, roleId int64, permIds []int64) error {
	if len(permIds) == 0 {
		return nil
	}
	records := make([]*RelRolePermission, len(permIds))
	for i, permId := range permIds {
		records[i] = &RelRolePermission{
			RoleId:       roleId,
			PermissionId: permId,
		}
	}
	fmt.Println("records: ", records)
	return db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "uidx_role_permission"}},
		DoNothing: true,
	}).Create(&records).Error
}

// RemoveRolePermByPermId 将指定权限从关联所有角色上移除
func (m *RelModel) RemoveRolePermByPermId(ctx context.Context, db *gorm.DB, permId int64) error {
	return db.Where("`permission_id` = ?", permId).Delete(&RelRolePermission{}).Error
}

// RemoveRolePermByPermIds 将指定权限从角色上移除
func (m *RelModel) RemoveRolePermByPermIds(ctx context.Context, db *gorm.DB, roleId int64, permIds []int64) error {
	if len(permIds) == 0 {
		return nil
	}
	return db.Where("`role_id` = ? AND `permission_id` IN ?", roleId, permIds).Delete(&RelRolePermission{}).Error
}

// RemoveRolePermByRoleId 删除角色对应权限关联关系
func (m *RelModel) RemoveRolePermByRoleId(ctx context.Context, db *gorm.DB, roleId int64) error {
	return db.Where("`role_id` = ?", roleId).Delete(&RelRolePermission{}).Error
}
