package db1

import (
	"github.com/go-xorm/builder"
)

type User struct {
	Id         uint64 `xorm:"id"`
	Cellphone  string `xorm:"cellphone"`
	InsertTime int64  `xorm:"insertTime bigint created"`
	ModifyTime int64  `xorm:"modifyTime bigint updated"`
}

func CreateUser(cellphone string) User {
	user := new(User)
	user.Cellphone = cellphone
	_, err := orm.Insert(user)
	if err != nil {
		panic(err)
	}
	return *user
}

func GetUserById(id uint64) User {
	var user User
	_, err := orm.Where("id=?", id).Limit(1, 0).Get(&user)
	if err != nil {
		panic(err)
	}
	return user
}

func GetUsersByIds(ids []uint64) []User {
	var users []User
	var err error
	if len(ids) > 0 {
		err = orm.Where(builder.In("id", ids)).Find(&users)
	} else {
		err = orm.Where("1").Find(&users)
	}
	if err != nil {
		panic(err)
	}
	return users
}
