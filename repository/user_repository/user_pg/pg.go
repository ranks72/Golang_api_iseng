package user_pg

import (
	"MALIKI-KARIM/entity"
	"MALIKI-KARIM/pkg/errs"
	"MALIKI-KARIM/repository/user_repository"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type userPG struct {
	db *gorm.DB
}

func NewUserPG(db *gorm.DB) user_repository.UserRepository {
	return &userPG{
		db: db,
	}
}

func (u *userPG) Register(userPayload *entity.User) (*entity.User, errs.MessageErr) {
	fmt.Println(userPayload)
	if err := u.db.Create(userPayload).Error; err != nil {
		if strings.Contains(err.Error(), "unique") {
			checkerr := errs.NewInternalServerErrorr("username is already used")
			return nil, checkerr
		}
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}
	return userPayload, nil
}

func (u *userPG) Login(userPayload *entity.User) (*entity.User, errs.MessageErr) {
	user := &entity.User{}

	if err := u.db.Select("id", "username", "amount").
		First(user, "username", userPayload.Username).Error; err != nil {
		return nil, errs.NewInternalServerErrorr("username salah")
	}
	fmt.Println(user)
	return user, nil
}

func (u *userPG) GetUsername(userId int) (*entity.User, errs.MessageErr) {

	user := &entity.User{}

	err := u.db.First(user, "id", userId).Error

	if err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	return user, nil

}

func (u *userPG) Deposit(username string, userPayload *entity.User) (*entity.User, errs.MessageErr) {
	query := u.db.Where("username", username).Updates(userPayload)
	err := query.Error
	if err == nil && query.RowsAffected < 1 {
		return nil, errs.NewNotFoundError("user doesn't exit")
	}
	user := &entity.User{}

	err = u.db.First(user, "username", username).Error

	if err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	return user, nil
}
