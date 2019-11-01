package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// User struct
type User struct {
	Id        int    `orm:"auto"`
	FirstName string `form:"first_name"`
	LastName  string `form:"last_name"`
	Email     string `form:"email"`
	Password  string `form:"password"`
	Phone     string `form:"phone"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// RandomSalt for password
type RandomSalt struct {
	PasswordSalt string
}

// LoginForm for password
type LoginForm struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

// Notes struct
type Notes struct {
	Id        int `orm:"auto"`
	UserId    int
	Comments  string `form:"comments"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func init() {
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Notes))
}
