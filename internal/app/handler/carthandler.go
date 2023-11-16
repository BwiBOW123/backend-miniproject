package handler

import (
	service "BwiBOW123/backend-miniproject/internal/app/services"
	"BwiBOW123/backend-miniproject/internal/domain"
	"BwiBOW123/backend-miniproject/logs"
	"encoding/json"
	"net/http"
)

type CartHandler struct {
	service *service.CartService
}

func NewCartHandler(service *service.CartService) *CartHandler {
	return &CartHandler{service: service}
}

func (h *CartHandler) CreateCart(w http.ResponseWriter, r *http.Request) {
	var cart domain.Cart
	if err := json.NewDecoder(r.Body).Decode(&cart); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logs.Error(err)
		return
	}

	if err := h.service.CreateCart(&cart); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logs.Error(err)
		return
	}
	logs.Info("201 Create Suscess")
	w.WriteHeader(http.StatusCreated)
}
func (h *CartHandler) CreateCartProducts(w http.ResponseWriter, r *http.Request) {
	var cartproduct domain.CartProducts
	if err := json.NewDecoder(r.Body).Decode(&cartproduct); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logs.Error(err)
		return
	}

	if err := h.service.CreateCartProduct(&cartproduct); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logs.Error(err)
		return
	}
	h.service.UpdateTotal()
	logs.Info("201 Create Suscess")
	w.WriteHeader(http.StatusCreated)
}

func (p *CartHandler) GetCarts(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	carts, err := p.service.GetCarts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logs.Error(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(carts)
	logs.Info("Get Success")
}

func (p *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id") // Assuming ID is passed as a query parameter
	cart, err := p.service.GetCart(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		logs.Error(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cart)
	logs.Info("Get Success")
}
func (p *CartHandler) GetCartProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id") // Assuming ID is passed as a query parameter
	cartproduct, err := p.service.GetCartProducts(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		logs.Error(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cartproduct)
	logs.Info("Get Success")
}
func (p *CartHandler) GetCartProductsEmail(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id") // Assuming ID is passed as a query parameter
	cartproduct, err := p.service.GetCartProductsEmail(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		logs.Error(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cartproduct)
	logs.Info("Get Success")
}
func (p *CartHandler) DeleteCartProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id") // Assuming ID is passed as a query parameter
	err := p.service.DeleteCartProducts(id)
	p.service.UpdateTotal()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		logs.Error(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("delete Success")
	logs.Info("delete Success")
}
