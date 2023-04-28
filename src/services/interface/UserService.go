package _interface

import "Job/src/beans"

type UserService interface {
	Register(user *beans.User) (int64, error)
}
