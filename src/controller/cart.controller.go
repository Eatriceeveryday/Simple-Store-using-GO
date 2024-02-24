package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"synapsis/src/model"
	"synapsis/src/utils"
)

func (c *Controller) GetCart(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("ID").(string)
	if id == "" {
		utils.JSONError(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var cart []model.Carts

	err := c.DB.Preload("Product").Where("customer_id = ? ", id).Find(&cart).Error
	if err != nil {
		fmt.Println(err)
		utils.JSONError(w, "Bad Request", http.StatusBadRequest)
		return
	}

	utils.JSONResponse(w, cart, http.StatusOK)
}

type AddCartRequest struct {
	ProductId string `validate:"required"`
	Quantity  int    `validate:"required"`
}

func (c *Controller) AddCart(w http.ResponseWriter, r *http.Request) {
	customerId := r.Context().Value("ID").(string)
	if customerId == "" {
		utils.JSONError(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var addReq AddCartRequest
	err := json.NewDecoder(r.Body).Decode(&addReq)
	if err != nil {
		fmt.Println(err)
		utils.JSONError(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	err = utils.Validate.Struct(addReq)
	if err != nil {
		utils.JSONError(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	cart := model.Carts{
		CustomerID: customerId,
		ProductID:  addReq.ProductId,
		Quantity:   addReq.Quantity,
	}

	err = c.DB.Create(&cart).Error
	if err != nil {
		utils.JSONError(w, "Failed to add to cart", http.StatusBadRequest)
		return
	}

	resp := utils.SuccessResponse{Message: "Added to cart"}
	utils.JSONResponse(w, resp, http.StatusOK)
}

func (c *Controller) DeleteCart(w http.ResponseWriter, r *http.Request) {
	customerId := r.Context().Value("ID").(string)
	if customerId == "" {
		utils.JSONError(w, "Bad Request", http.StatusBadRequest)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		utils.JSONError(w, "Bad Request", http.StatusBadRequest)
		return
	}

	result := c.DB.Where("id = ? AND customer_id = ? ", id, customerId).Delete(&model.Carts{})
	if result.Error != nil {
		utils.JSONError(w, "Failed to delete item", http.StatusBadRequest)
		return
	} else if result.RowsAffected < 1 {
		utils.JSONError(w, "Invalid cart id", http.StatusBadRequest)
		return
	}

	resp := utils.SuccessResponse{Message: "Item deleted"}
	utils.JSONResponse(w, resp, http.StatusOK)
}
