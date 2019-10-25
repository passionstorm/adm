package models

import (
	"adm/app/pkg/util"
	"github.com/markbates/pop/nulls"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"strings"
	"time"
)

type AdminUserModel struct {
	ID           int64
	Name         nulls.String      `db:"name"`
	Email        nulls.String      `db:"email"`
	PasswordHash nulls.String      `db:"password_hash"`
	InsertedAt   time.Time         `db:"inserted_at"`
	UpdatedAt    time.Time         `db:"updated_at"`
	Password     string            `db:"-"`
	Errors       map[string]string `db:"-"`
}

func AdminUser(name, email, password string) *AdminUserModel {
	return &AdminUserModel{
		Name:     nulls.NewString(name),
		Email:    nulls.NewString(downcase(email)),
		Password: password,
	}
}

func GetAdminUser(id int64) (*AdminUserModel, error) {
	sql := `
		SELECT id, name, email, password_hash
		FROM admin_users
		WHERE id=$1
	`

	adminUser := AdminUserModel{}
	err := db.Get(&adminUser, sql, id)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &adminUser, nil
}

func GetAdminByEmail(email string) (*AdminUserModel, error) {
	sql := `
		SELECT id, name, email, password_hash, inserted_at, updated_at 
		FROM admin_users
		WHERE email=$1
	`

	adminUser := AdminUserModel{}
	err := db.Get(&adminUser, sql, email)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &adminUser, nil
}

func (u *AdminUserModel) CheckAuth(password string) bool {
	hashedP := []byte(u.PasswordHash.String)
	p := []byte(password)

	err := bcrypt.CompareHashAndPassword(hashedP, p)
	if err != nil {
		return false
	}

	return true
}

func (u *AdminUserModel) Create() error {
	u.InsertedAt = time.Now()
	u.UpdatedAt = time.Now()
	valid := u.Validate()

	if !valid {
		return errors.New(util.MapToStr(u.Errors))
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.WithStack(err)
	}
	u.PasswordHash = nulls.NewString(string(hash))
	stmt, err := db.NamedExec(`
		INSERT INTO admin_users (name, email, password_hash, inserted_at, updated_at)
		VALUES (:name, :email, :password_hash, :inserted_at, :updated_at)
	`, u)
	if err != nil {
		return errors.WithStack(err)
	}
	u.ID, err = stmt.LastInsertId()
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (u *AdminUserModel) Validate() bool {
	u.Errors = make(map[string]string)
	re := regexp.MustCompile(".+@.+\\..+")

	email := strings.TrimSpace(u.Email.String)
	password := strings.TrimSpace(u.Password)
	matched := re.Match([]byte(email))

	if matched == false {
		u.Errors["Email"] = "Please enter a valid email address"
	}

	if len(password) <= 6 {
		u.Errors["Password"] = "Password must be atleast 6 characters"
	}

	return len(u.Errors) == 0
}
