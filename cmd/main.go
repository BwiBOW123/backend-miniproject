package main

import (
	handler "BwiBOW123/backend-miniproject/internal/app/handler"
	"BwiBOW123/backend-miniproject/internal/app/repository"
	service "BwiBOW123/backend-miniproject/internal/app/services"
	"BwiBOW123/backend-miniproject/internal/domain"
	"BwiBOW123/backend-miniproject/logs"
	"BwiBOW123/backend-miniproject/pkg/config"
	"log"

	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allows all origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	config.LoadConfig()
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	db.AutoMigrate(&domain.Member{})
	db.AutoMigrate(&domain.Payment{})
	db.AutoMigrate(&domain.Product{})
	db.AutoMigrate(&domain.Category{})
	db.AutoMigrate(&domain.Cart{})

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	r := mux.NewRouter()
	r.HandleFunc("/products", productHandler.GetProducts).Methods("GET")
	r.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
	r.HandleFunc("/users", userHandler.Register).Methods("POST")
	r.HandleFunc("/login", userHandler.Login).Methods("POST")
	r.HandleFunc("/users", userHandler.GetUser).Methods("GET")
	logs.Info("Service on port: 8000")
	hand := c.Handler(r)
	http.ListenAndServe(":8000", hand)

}
