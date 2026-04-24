package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"smallcase/models"
	"smallcase/service"
)

type handler struct {
	service service.Service
}

func NewHandler(s service.Service) *handler {
	return &handler{
		service: s,
	}
}

func (h *handler) Buy(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("user_id")
	var m models.Stock

	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"Error":   err,
		})
		return
	}
	if m.Quantity < 0 {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"Error":   "Quantity should not be negative",
		})
		return
	}
	err = h.service.Buy(userId, &m)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"Error":   fmt.Errorf("Not able to save the holdings%+v", err),
		})
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"Message": "SuccessFull Buy",
	})
}

func (h *handler) FetchingHoldings(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("user_id")

	stock := h.service.FetchingHoldings(userId)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"stocks":  stock,
	})

}

func (h *handler) FetchingReturns(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("user_id")
	// var currentPrice map[string]int
	currentPrice := make(map[string]int)
	err := json.NewDecoder(r.Body).Decode(&currentPrice)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"Error":   err,
		})
		return
	}
	profit := h.service.FetchingReturns(userId, currentPrice)
	fmt.Printf("Current Price is %+v\n", currentPrice)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"stocks":  profit,
	})
}

func (h *handler) Sell(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("user_id")
	var m models.Stock

	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"Error":   err,
		})
		return
	}
	if m.Quantity < 0 {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"Error":   "Quantity should not be negative",
		})
		return
	}
	err = h.service.Sell(userId, &m)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"Error":   fmt.Errorf("Not able to save the holdings%+v", err),
		})
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"Message": "SuccessFull Sell",
	})
}