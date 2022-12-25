package service

import (
	"MALIKI-KARIM/dto"
	"MALIKI-KARIM/entity"
	"MALIKI-KARIM/pkg/errs"
	"MALIKI-KARIM/pkg/helpers"
	"MALIKI-KARIM/repository/transaksi_repository"
	"MALIKI-KARIM/repository/user_repository"
	"fmt"
	"time"
)

type UserService interface {
	Register(userPayload *dto.RegisterRequest) (*entity.User, errs.MessageErr)
	Login(userPayload *dto.LoginRequest) (*dto.LoginResponse, errs.MessageErr)
	Deposit(userId int, userPayload *dto.DepositRequest) (*dto.DepositResponse, errs.MessageErr)
	Transfer(userId int, userPayload *dto.TransferRequest) (*dto.TransferResponse, errs.MessageErr)
}

type userService struct {
	userRepo      user_repository.UserRepository
	transaksiRepo transaksi_repository.TransaksiRepository
}

func NewUserService(
	userRepo user_repository.UserRepository,
	transaksiRepo transaksi_repository.TransaksiRepository,
) UserService {
	return &userService{
		userRepo:      userRepo,
		transaksiRepo: transaksiRepo,
	}
}

func (u *userService) Register(userPayload *dto.RegisterRequest) (*entity.User, errs.MessageErr) {
	err := helpers.ValidateStruct(userPayload)

	if err != nil {
		return nil, err
	}
	user := &entity.User{
		Username:  userPayload.Username,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err != nil {
		return nil, err
	}
	//fmt.Println(user)
	data_user, err := u.userRepo.Register(user)

	if err != nil {
		return nil, err
	}

	return data_user, nil
}

func (u *userService) Login(userPayload *dto.LoginRequest) (*dto.LoginResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(userPayload)

	if err != nil {
		return nil, err
	}

	payload := &entity.User{
		Username: userPayload.Username,
	}

	user, err := u.userRepo.Login(payload)

	if err != nil {

		return nil, err
	}

	username := user.Username
	amount := user.Amount
	conv_amount := fmt.Sprintf("%#v", amount)

	msg := "Hello " + username + "! Your balance is $" + conv_amount
	response := &dto.LoginResponse{
		Message: msg,
	}

	return response, nil
}

func (u *userService) Deposit(userId int, userPayload *dto.DepositRequest) (*dto.DepositResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(userPayload)
	if err != nil {
		return nil, err
	}

	user, err := u.userRepo.GetUsername(userId)

	if err != nil {

		return nil, err
	}

	amount := user.Amount + userPayload.Deposit

	newpayload := &entity.User{
		Amount: amount,
	}

	new_result, err := u.userRepo.Deposit(user.Username, newpayload)

	if err != nil {
		return nil, err
	}
	amount = new_result.Amount
	conv_amount := fmt.Sprintf("%#v", amount)
	msg := "Your balance is $" + conv_amount
	response := &dto.DepositResponse{
		Message: msg,
	}

	return response, nil
}

func (u *userService) Transfer(userId int, userPayload *dto.TransferRequest) (*dto.TransferResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(userPayload)
	if err != nil {
		return nil, err
	}

	user1, err := u.userRepo.GetUsername(userId)

	if err != nil {

		return nil, err
	}

	user_amount := 0
	if user1.Amount < userPayload.Amount {
		userPayload.Amount = user1.Amount
	} else {
		user_amount = user1.Amount - userPayload.Amount
	}

	newpayload := &entity.User{
		Amount: user_amount,
	}

	user1, err = u.userRepo.Deposit(user1.Username, newpayload)
	if err != nil {
		return nil, err
	}

	payload_transaksi := &entity.User{
		Username: userPayload.To,
	}
	user2, err := u.userRepo.Login(payload_transaksi)

	amount := user2.Amount + userPayload.Amount
	user2_payload := &entity.User{
		Amount: amount,
	}
	user2, err = u.userRepo.Deposit(user2.Username, user2_payload)
	if err != nil {
		return nil, err
	}

	response := &dto.TransferResponse{
		Message: "berhasil",
	}

	return response, nil
}
