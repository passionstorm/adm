package model

type Account struct {
	Username string `gorm:"unique" json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token";sql:"-"`
	Model
}

func (account *Account) GetAllUser() []*Account {
	acc := make([]*Account, 0)
	db.Find(&acc)

	return acc
}

func (account *Account) Create() []*Account {
	acc := make([]*Account, 0)
	db.Find(&acc)

	return acc
}
