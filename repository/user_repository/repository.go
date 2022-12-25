package user_repository

import (
	"MALIKI-KARIM/entity"
	"MALIKI-KARIM/pkg/errs"
)

type UserRepository interface {
	Register(user *entity.User) (*entity.User, errs.MessageErr)
	Login(user *entity.User) (*entity.User, errs.MessageErr)
	Deposit(username string, user *entity.User) (*entity.User, errs.MessageErr)
	GetUsername(userId int) (*entity.User, errs.MessageErr)
}
