package model

import "log"

type Account struct {
	Username string  `json:"username"`
	Name     *string `json:"name"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
	Phone    *string `json:"phone"`
	Role     *string `json:"role"`
	Actived  bool    `json:"actived"`
	Token    string  `json:"token";sql:"-"`
	Version  string  `json:"version"`
	Model
}

func (account *Account) GetAllUser() []Account {
	acc := []Account{}
	err := db.Select(&acc, "SELECT * FROM accounts")
	if err != nil {
		log.Fatal(err)
	}

	return acc
}
