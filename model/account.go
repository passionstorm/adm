package model

import (
	"adm/pkg/db"
	"time"
)

type AccountModel struct {
	Base      `json:"-"`
	ID        uint       `json:"id"`
	Username  string     `json:"username"`
	Name      *string    `json:"name"`
	Email     *string    `json:"email"`
	Password  *string    `json:"password"`
	Phone     *string    `json:"phone"`
	Role      *string    `json:"role"`
	Actived   bool       `json:"actived"`
	Token     string     `json:"token";sql:"-"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"-"`
	Version   uint       `json:"-"`
}

func Account() *AccountModel {
	return &AccountModel{Base: Base{Table: "accounts"}}
}

func (t *AccountModel) All() []AccountModel {
	items, _ := db.Get(t.Table).All()
	list := make([]AccountModel, 0)
	for _, v := range items {
		list = append(list, t.Base.MapToModel(t, v).(AccountModel))
	}
	return list
}
