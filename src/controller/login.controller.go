package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"synapsis/src/config"
	"synapsis/src/model"
	"synapsis/src/utils"

	"golang.org/x/crypto/bcrypt"
)

type loginRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

type loginResponse struct {
	Token    string         `json:"token"`
	Customer model.Customer `json:"customer"`
}

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
	var request loginRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.JSONError(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	err = utils.Validate.Struct(request)
	if err != nil {
		fmt.Println(err)
		utils.JSONError(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	var customer model.Customer

	err = c.DB.Where("email = ?", request.Email).First(&customer).Error
	if err != nil {
		utils.JSONError(w, "No user found", http.StatusBadRequest)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(request.Password))
	if err != nil {
		utils.JSONError(w, "Password Wrong", http.StatusBadRequest)
		return
	}

	config, err := config.LoadConfig(".")
	if err != nil {
		fmt.Println(err.Error())
		utils.JSONError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	accessToken, err := utils.CreateToken(customer.ID, config.AccessKey)
	if err != nil {
		fmt.Println(err.Error())
		utils.JSONError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	resp := loginResponse{Token: accessToken, Customer: customer}
	utils.JSONResponse(w, resp, http.StatusOK)
}
