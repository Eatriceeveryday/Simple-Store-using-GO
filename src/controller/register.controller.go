package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"synapsis/src/model"
	"synapsis/src/utils"
)

type registerRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
	Name     string `validate:"required"`
}

func (c *Controller) Register(w http.ResponseWriter, r *http.Request) {
	var request registerRequest
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

	customer := model.Customer{
		Email:    request.Email,
		Password: request.Password,
		Name:     request.Name,
	}
	result := c.DB.Create(&customer)

	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
			utils.JSONError(w, "Email already used", http.StatusBadRequest)
			return
		}
		utils.JSONError(w, "Failed to create user", http.StatusBadRequest)
		return

	}

	resp := utils.SuccessResponse{Message: "User Created"}
	utils.JSONResponse(w, resp, http.StatusOK)
}
