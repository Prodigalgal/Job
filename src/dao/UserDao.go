package dao

import (
	"Job/src/beans"
	"Job/src/connect"
)

type UserDaoImpl struct{}

func (u UserDaoImpl) Register(user *beans.User) (int64, error) {
	insert, err := connect.Engine.Table("user").Insert(user)
	if err != nil {
		return 0, err
	}
	return insert, nil
}
