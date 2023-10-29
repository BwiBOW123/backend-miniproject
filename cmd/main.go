package main

import (
	handler "BwiBOW123/backend-miniproject/internal/app/handler"
	"BwiBOW123/backend-miniproject/internal/app/repository"
	service "BwiBOW123/backend-miniproject/internal/app/services"
	"BwiBOW123/backend-miniproject/internal/domain"
	"BwiBOW123/backend-miniproject/pkg/config"
	"log"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"net/http"
)

func main() {
	config.LoadConfig()

	dsn := "bwibow:3009@tcp(127.0.0.1:3306)/miniproject?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	log.Println("Load DB success")
	db.AutoMigrate(&domain.User{})

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	log.Println("Add Services")
	r := mux.NewRouter()
	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/users", userHandler.GetUser).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
