package main

import (
	handler "BwiBOW123/backend-miniproject/internal/app/handler"
	"BwiBOW123/backend-miniproject/internal/app/repository"
	service "BwiBOW123/backend-miniproject/internal/app/services"
	"BwiBOW123/backend-miniproject/internal/domain"
	"BwiBOW123/backend-miniproject/logs"
	"BwiBOW123/backend-miniproject/pkg/config"
	"log"

	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/checkout/session"

	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func createCheckoutSession(w http.ResponseWriter, r *http.Request) {
	domain := "http://localhost:8000"
	params := &stripe.CheckoutSessionParams{
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			&stripe.CheckoutSessionLineItemParams{
				// Provide the exact Price ID (for example, pr_1234) of the product you want to sell
				Price:    stripe.String("100"),
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(domain + "/"),
		CancelURL:  stripe.String(domain + "/"),
	}

	s, err := session.New(params)

	if err != nil {
		log.Printf("session.New: %v", err)
	}

	http.Redirect(w, r, s.URL, http.StatusSeeOther)
}

func main() {

	stripe.Key = "sk_test_51OD87HGiZKMaUrw3UBWM8U2iI4ZqEVTmZKB5UtPz9mNLkCg7SHl6dplKkDBby4UsxFICCxKVcvSaQfOL6JFyqVz500yZrXfCD0"

	//-------------------------------------------------------------

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
	db.AutoMigrate(&domain.Category{}, &domain.Image{}, &domain.File{}, &domain.Member{}, &domain.Cart{}, &domain.Payment{}, &domain.Product{}, &domain.CartProducts{})

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	CartRepo := repository.NewCartRepository(db)
	CartService := service.NewCartService(CartRepo)
	CartHandler := handler.NewCartHandler(CartService)

	r := mux.NewRouter()
	r.Handle("/", http.FileServer(http.Dir("public")))
	r.HandleFunc("/create-checkout-session", createCheckoutSession).Methods("POST")
	r.HandleFunc("/Uploadfile", productHandler.UploadFile).Methods("POST")
	r.HandleFunc("/Uploadimage", productHandler.UploadImage).Methods("POST")
	r.HandleFunc("/products", productHandler.GetProducts).Methods("GET")
	r.HandleFunc("/product", productHandler.GetProduct).Methods("GET")
	r.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
	r.HandleFunc("/users", userHandler.Register).Methods("POST")
	r.HandleFunc("/login", userHandler.Login).Methods("POST")
	//r.HandleFunc("/users", userHandler.GetUser).Methods("GET")
	r.HandleFunc("/users", userHandler.GetUserbyEmail).Methods("GET")
	r.HandleFunc("/productimages", productHandler.GetProductwithImage).Methods("GET")
	r.HandleFunc("/productimagesByCat", productHandler.GetProductwithImageByCat).Methods("GET")
	r.HandleFunc("/productimagesById", productHandler.GetProductwithImageByid).Methods("GET")
	r.HandleFunc("/Carts", CartHandler.CreateCart).Methods("POST")
	r.HandleFunc("/Carts", CartHandler.GetCarts).Methods("GET")
	r.HandleFunc("/Cart", CartHandler.GetCart).Methods("GET")
	r.HandleFunc("/CartProduct", CartHandler.GetCartProducts).Methods("GET")
	r.HandleFunc("/CartProductEmail", CartHandler.GetCartProductsEmail).Methods("GET")
	r.HandleFunc("/CartProduct", CartHandler.CreateCartProducts).Methods("POST")
	r.HandleFunc("/DeleteProduct", CartHandler.DeleteCartProducts).Methods("DELETE")
	logs.Info("Service on port: 8000")
	hand := c.Handler(r)
	http.ListenAndServe(":8000", hand)
}
