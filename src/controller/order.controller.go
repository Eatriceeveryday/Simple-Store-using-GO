package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"synapsis/src/model"
	"synapsis/src/utils"

	"gorm.io/gorm"
)

type AddOrderRequest struct {
	CartId []string `validate:"required"`
}

func (c *Controller) AddOrder(w http.ResponseWriter, r *http.Request) {
	customerId := r.Context().Value("ID").(string)
	if customerId == "" {
		utils.JSONError(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var req AddOrderRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println(err)
		utils.JSONError(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	err = utils.Validate.Struct(req)
	if err != nil {
		fmt.Println(err)
		utils.JSONError(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	err = c.DB.Transaction(func(tx *gorm.DB) error {
		// Get Cart to create order
		var carts []model.Carts
		err = tx.Preload("Product").Where("id IN ?", req.CartId).Where("customer_id = ?", customerId).Find(&carts).Error
		if err != nil {
			return err
		}

		// update the cart to delete it
		err = tx.Where("id IN ?", req.CartId).Delete(&model.Carts{}).Error
		if err != nil {
			return err
		}

		// Add cart to order
		order := model.Order{
			CustomerID: customerId,
		}
		err = tx.Create(&order).Error
		if err != nil {
			return err
		}

		var productOrder []model.ProductOrder

		for _, item := range carts {
			productOrder = append(productOrder, model.ProductOrder{
				OrderID:   order.ID,
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
			})
		}
		err = tx.Create(&productOrder).Error
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
		utils.JSONError(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	resp := utils.SuccessResponse{Message: "Order Create"}
	utils.JSONResponse(w, resp, http.StatusOK)
}

type ProductWithQuantity struct {
	Product  model.Product `json:"product"`
	Quantity int           `json:"quantity"`
}

type GetOrderResponse struct {
	OrderId    string                `json:"orderId"`
	TotalPrice int                   `json:"totalPrice"`
	Product    []ProductWithQuantity `json:"products"`
}

func (c *Controller) GetOrder(w http.ResponseWriter, r *http.Request) {
	customerId := r.Context().Value("ID").(string)
	if customerId == "" {
		utils.JSONError(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var orders []model.Order
	err := c.DB.Where("customer_id = ?", customerId).Find(&orders).Error
	if err != nil {
		fmt.Println(err.Error())
		utils.JSONError(w, "Invalid Request 1", http.StatusBadRequest)
		return
	}
	orderId := []string{} // slice to contai all customer order id
	for _, order := range orders {
		orderId = append(orderId, order.ID)
	}

	var productOrder []model.ProductOrder
	err = c.DB.Preload("Product").Where("order_id IN ?", orderId).Find(&productOrder).Error
	if err != nil {
		fmt.Println(err)
		utils.JSONError(w, "Invalid Request ", http.StatusBadRequest)
		return
	}

	// Group customer order by orderId and group the product into an array
	resp := []GetOrderResponse{}
	for i := 0; i < len(productOrder); i++ {
		id := productOrder[i].OrderID
		var products []ProductWithQuantity
		totalPrice := 0
		for i < len(productOrder) {
			//inser the first index into array
			if i == 0 {
				products = append(products, ProductWithQuantity{
					Product:  productOrder[i].Product,
					Quantity: productOrder[i].Quantity,
				})
				totalPrice += productOrder[i].Quantity * productOrder[i].Product.Price
				i++
				continue
			}
			if id == productOrder[i].OrderID {
				products = append(products, ProductWithQuantity{
					Product:  productOrder[i].Product,
					Quantity: productOrder[i].Quantity,
				})
				totalPrice += productOrder[i].Quantity * productOrder[i].Product.Price
				i++
				continue
			} else {
				i-- // Decrease i by 1 so when for loop at the end i wil increase i++
				break
			}
		}
		resp = append(resp, GetOrderResponse{
			OrderId:    id,
			TotalPrice: totalPrice,
			Product:    products,
		})
	}

	utils.JSONResponse(w, resp, http.StatusOK)
}
