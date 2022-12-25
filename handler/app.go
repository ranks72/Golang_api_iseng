package handler

import (
	"MALIKI-KARIM/database"
	"MALIKI-KARIM/repository/transaksi_repository/transaksi_pg"
	"MALIKI-KARIM/repository/user_repository/user_pg"
	"MALIKI-KARIM/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

const PORT = ":8080"

func RunServer() *gin.Engine {
	router := gin.Default()
	db := database.GetDb()

	userRepo := user_pg.NewUserPG(db)
	transaksiRepo := transaksi_pg.NewTransaksiPG(db)
	userService := service.NewUserService(userRepo, transaksiRepo)

	userRestHandler := newUserHandler(userService)

	userRoute := router.Group("/users")
	{
		userRoute.POST("/register", userRestHandler.Register)
		userRoute.POST("/login", userRestHandler.Login)
	}

	depositRoute := router.Group("/deposit")
	{
		depositRoute.PUT("/:userId", userRestHandler.Deposit)
	}

	transferRoute := router.Group("/transfer")
	{
		transferRoute.PUT("/:userId", userRestHandler.Transfer)
	}
	fmt.Println("Server running on PORT =>", PORT)
	router.Run(PORT)
	return router
}
