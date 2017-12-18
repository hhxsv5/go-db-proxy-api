package mdb1

import (
	"github.com/go-xorm/builder"
)

type User struct {
	Id         uint32
	Cellphone  string
	InsertTime uint32 `xorm:"created"`
	ModifyTime uint32 `xorm:"updated"`
}

func CreateUser(cellphone string) *User {
	user := new(User)
	user.Cellphone = cellphone
	_, err := orm.Insert(user)
	if err != nil {
		return nil
	}
	return user
}

func GetUserById(id uint32) *User {
	var user *User
	orm.Where("id=?", id).Limit(1, 0).Get(user)
	return user
}

func GetUsersByIds(ids []uint32) *[]User {
	var users *[]User
	orm.Where(builder.In("id", ids)).Find(users)
	return users
}
