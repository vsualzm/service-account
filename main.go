package main

import (
	"log"
	"os"
	"service-account/config"
	"service-account/handler"
	"service-account/middleware"
	"service-account/repository"
	"service-account/service"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET is not set in the environment variables")
	}

	// cheking database
	db := config.ConnectDB()
	defer db.Close()

	// Initialize dependencies
	accountRepo := repository.NewAccountRepository(db)
	accountService := service.NewAccountService(accountRepo, secret)
	accountHandler := handler.NewAccountHandler(accountService)

	// Router
	router := echo.New()

	// test api health
	router.GET("/health", accountHandler.TestHealtAPI)

	// daftar dan login
	router.POST("/daftar", accountHandler.DaftarAccount)
	router.POST("/login", accountHandler.LoginAccount)

	// tabung dan tarik
	router.POST("transaction/tarik-saldo", accountHandler.Tarik, middleware.AuthMiddleware)
	router.POST("transaction/tabung-saldo", accountHandler.Tabung, middleware.AuthMiddleware)
	router.GET("transaction/get-saldo/:norekening", accountHandler.CheckSaldo, middleware.AuthMiddleware)

	// start server
	router.Logger.Fatal(router.Start(":8080"))

}
