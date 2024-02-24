package controller

import (
	"fmt"
	"net/http"
	"synapsis/src/model"
	"synapsis/src/utils"

	"github.com/go-chi/chi/v5"
)

func (c *Controller) GetAllProduct(w http.ResponseWriter, r *http.Request) {
	var products []model.Product

	err := c.DB.Find(&products).Error
	if err != nil {
		fmt.Println(err)
		utils.JSONError(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	utils.JSONResponse(w, products, http.StatusOK)
}

func (c *Controller) GetProductbyCategory(w http.ResponseWriter, r *http.Request) {
	categoryName := chi.URLParam(r, "category")

	var category model.Category
	err := c.DB.Where("name = ?", categoryName).Find(&category).Error
	if err != nil {
		fmt.Println(err)
		utils.JSONError(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	var products []model.ProductCategories
	err = c.DB.Preload("Product").Where("category_id = ? ", category.ID).Find(&products).Error
	if err != nil {
		fmt.Println(err)
		utils.JSONError(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	utils.JSONResponse(w, products, http.StatusOK)
}
