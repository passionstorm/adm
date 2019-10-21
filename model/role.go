package model

import (
	"adm/pkg/db"
	"adm/pkg/db/dialect"
	"strconv"
)

type RoleModel struct {
	Base
	Id        int64
	Name      string
	Slug      string
	CreatedAt string
	UpdatedAt string
}

func Role() *RoleModel {
	return &RoleModel{Base: Base{Table: "roles"}}
}

func RoleWithId(id string) *RoleModel {
	idInt, _ := strconv.Atoi(id)
	return &RoleModel{Base: Base{Table: "roles"}, Id: int64(idInt)}
}

func (t *RoleModel) Find(id interface{}) *RoleModel {
	item, _ := db.Get(t.Table).Find(id)
	return t.MapToModel(item)
}

func (t *RoleModel) New(name, slug string) *RoleModel {

	id, _ := db.Get(t.Table).Insert(dialect.H{
		"name": name,
		"slug": slug,
	})

	t.Id = id
	t.Name = name
	t.Slug = slug

	return t
}

func (t *RoleModel) Update(name, slug string) *RoleModel {

	_, _ = db.Get(t.Table).
		Where("id", "=", t.Id).
		Update(dialect.H{
			"name": name,
			"slug": slug,
		})

	t.Name = name
	t.Slug = slug

	return t
}

func (t *RoleModel) CheckPermission(permissionId string) bool {
	checkPermission, _ := db.Get("permissions").
		Where("permission_id", "=", permissionId).
		Where("role_id", "=", t.Id).
		First()
	return checkPermission != nil
}

func (t *RoleModel) DeletePermissions() {
	_ = db.Get("goadmin_role_permissions").
		Where("role_id", "=", t.Id).
		Delete()
}

func (t *RoleModel) AddPermission(permissionId string) {
	if permissionId != "" {
		if !t.CheckPermission(permissionId) {
			_, _ = db.Get("goadmin_role_permissions").
				Insert(dialect.H{
					"permission_id": permissionId,
					"role_id":       t.Id,
				})
		}
	}
}

func (t *RoleModel) MapToModel(m map[string]interface{}) *RoleModel {
	t.Id = m["id"].(int64)
	t.Name, _ = m["name"].(string)
	t.Slug, _ = m["slug"].(string)
	t.CreatedAt, _ = m["created_at"].(string)
	t.UpdatedAt, _ = m["updated_at"].(string)
	return t
}
