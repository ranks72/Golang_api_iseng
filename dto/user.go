package dto

import "MALIKI-KARIM/entity"

type RegisterRequest struct {
	Username string `json:"username" valid:"required~username cannot be empty" binding:"required"`
}

type LoginRequest struct {
	Username string `json:"username" valid:"required~username cannot be empty" binding:"required"`
}

type RegisterResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

func DataRegisterResponse(data entity.User) RegisterResponse {
	return RegisterResponse{
		ID:       data.ID,
		Username: data.Username,
		Message:  "Akun Telah dibuat",
	}
}

type LoginResponse struct {
	Message string `json:"message"`
}

type DepositRequest struct {
	Deposit int `json:"deposit"`
}

type DepositResponse struct {
	Message string `json:"message"`
}

type TransferRequest struct {
	Amount int    `json:"amount"`
	To     string `json:"to"`
}

type TransferResponse struct {
	Message string `json:"message"`
}
