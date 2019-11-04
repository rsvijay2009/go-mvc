package models

import (
	"github.com/astaxie/beego/orm"
)

// AddUser to create a new user
func (u *User) AddUser() int64 {
	db := orm.NewOrm()
	pk, err := db.Insert(u)

	if err != nil {
		panic(err)
	}

	return pk
}

// FindUser to find a specific user
func FindUser(email string) User {
	db := orm.NewOrm()
	var user User
	db.QueryTable(new(User)).Filter("email", email).One(&user)

	return user
}

// FindUserByID to find a specific user by Id
func FindUserByID(uid int) User {
	db := orm.NewOrm()
	var user User
	db.QueryTable(new(User)).Filter("id", uid).One(&user)
	return user
}
