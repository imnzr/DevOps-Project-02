package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/imnzr/DevOps-Project-02/controller"
	"github.com/imnzr/DevOps-Project-02/database"
	"github.com/imnzr/DevOps-Project-02/repository"
	"github.com/imnzr/DevOps-Project-02/service"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db, err := database.GetConnection()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	fmt.Println("Connected to database successfully")

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db)
	userController := controller.NewUserController(userService)

	router := httprouter.New()

	// Endpoint handlers
	router.GET("/users", userController.FindAll)
	router.GET("/user/:id", userController.FindById)
	router.POST("/user/login", userController.Login)
	router.POST("/user/create", userController.Create)
	router.PUT("/user/update/:id", userController.Update)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}

}
