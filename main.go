package main

import (
	"log"
	"net/http"
	"synapsis/src/config"
	"synapsis/src/controller"
	"synapsis/src/database"
	"synapsis/src/middleware"
	"synapsis/src/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

var (
	c controller.Controller
)

func init() {
	utils.Validate = validator.New(validator.WithRequiredStructEnabled())
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}
	database.OpenDatabaseConnection(config)
	c = controller.NewController(database.DB)
}

func main() {
	router := chi.NewRouter()

	router.Route("/api/v1", func(r chi.Router) {
		r.Post("/login", c.Login)
		r.Post("/register", c.Register)

		r.Group(func(r chi.Router) {
			r.Use(middleware.Authentication)
			r.Get("/cart", c.GetCart)
			r.Post("/cart", c.AddCart)
			r.Delete("/cart", c.DeleteCart)

			r.Get("/product", c.GetAllProduct)
			r.Get("/product/{category}", c.GetProductbyCategory)

			r.Post("/order", c.AddOrder)
			r.Get("/order", c.GetOrder)
		})
	})

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Println(err)
	}
}
