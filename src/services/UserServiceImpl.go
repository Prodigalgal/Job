package services

import (
	"Job/src/beans"
	"Job/src/dao"
)

type UserServiceImpl struct{}

func (u UserServiceImpl) Register(user *beans.User) (int64, error) {
	insert, err := dao.UserDaoImpl{}.Register(user)
	if err != nil {
		return 0, err
	}
	return insert, nil
}
