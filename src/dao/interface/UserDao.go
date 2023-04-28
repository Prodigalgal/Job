package _interface

import "Job/src/beans"

type UserDao interface {
	Register(user *beans.User) (int64, error)
}
