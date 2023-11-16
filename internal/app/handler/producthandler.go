package handler

import (
	service "BwiBOW123/backend-miniproject/internal/app/services"
	"BwiBOW123/backend-miniproject/internal/domain"
	"BwiBOW123/backend-miniproject/logs"
	"encoding/json"
	"net/http"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (p *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id") // Assuming ID is passed as a query parameter
	product, err := p.service.GetProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		logs.Error(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
	logs.Info("Get Success")
}

func (p *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	products, err := p.service.GetProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logs.Error(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
	logs.Info("Get Success")
}

func (p *ProductHandler) GetProductwithImage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	products, err := p.service.GetProductwithImage()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logs.Error(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
	logs.Info("Get Success")
}
func (p *ProductHandler) GetProductwithImageByCat(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}
	id := r.URL.Query().Get("id")
	products, err := p.service.GetProductwithImageByCat(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logs.Error(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
	logs.Info("Get Success")
}
func (p *ProductHandler) GetProductwithImageByid(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}
	id := r.URL.Query().Get("id")
	products, err := p.service.GetProductwithImageById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logs.Error(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
	logs.Info("Get Success")
}

func (p *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var product domain.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logs.Error(err)
		return
	}

	// Call service to create product
	err = p.service.CreateProduct(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logs.Error(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
	logs.Info("Create Success")
}

func (p *ProductHandler) UploadImage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var image domain.Image
	err := json.NewDecoder(r.Body).Decode(&image)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logs.Error(err)
		return
	}

	// Call service to create product
	err = p.service.UploadImage(&image)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logs.Error(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(image)
	logs.Info("Upload Success")
}
func (p *ProductHandler) UploadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var file domain.File
	err := json.NewDecoder(r.Body).Decode(&file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logs.Error(err)
		return
	}

	// Call service to create product
	err = p.service.UploadFile(&file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logs.Error(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(file)
	logs.Info("Upload Success")
}
