package main

import (
	"fmt"
	"log"
	"net/http"
	"todo-list-api/config"
	"todo-list-api/controller"
	"todo-list-api/helper"
	"todo-list-api/repository"
	"todo-list-api/service"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db, err := config.GetConnection()
	if err != nil {
		log.Fatal("database initialization error")
	}

	defer db.Close()
	fmt.Println("database connection successfully")

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db)
	userController := controller.NewUserController(userService)

	router := httprouter.New()

	router.GET("/api/users", userController.FindByAll)
	router.GET("/api/user/:userId", userController.FindById)
	router.POST("/api/user/signup", userController.Create)
	router.PUT("/api/user/:userId", userController.Update)
	router.DELETE("/api/user/:userId", userController.Delete)

	server := http.Server{
		Addr:    "localhost:8089",
		Handler: router,
	}

	err = server.ListenAndServe()
	helper.PanicIfError(err)
}
