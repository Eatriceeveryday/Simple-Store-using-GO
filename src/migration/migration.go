package main

import (
	"fmt"
	"log"
	"synapsis/src/config"
	"synapsis/src/database"
	"synapsis/src/model"

	"github.com/jaswdr/faker"
)

func Migrate() {
	fmt.Println("Starting Migration......")
	fmt.Println("Migration Done!")
	database.DB.AutoMigrate(&model.Customer{}, &model.Product{}, &model.Category{}, &model.Carts{}, &model.ProductCategories{}, &model.Order{}, &model.ProductOrder{})
}

func Seed() {
	fmt.Println("Seeding Database wait! ＼(ﾟｰﾟ＼) ( ﾉ ﾟｰﾟ)ﾉ")
	fake := faker.New()
	// Generate User for Test
	for i := 0; i < 3; i++ {
		customer := model.Customer{
			Name:     fake.Person().Name(),
			Email:    fake.Person().Contact().Email,
			Password: "12345",
		}
		database.DB.Create(&customer)
	}

	// Generate Product for Test
	for i := 0; i < 10; i++ {
		product := model.Product{
			Name:               fake.Company().Name(),
			Price:              fake.IntBetween(10_000, 100_000),
			ProductDescription: "Lorem Ipsum",
		}
		database.DB.Create(&product)
	}

	// Generate Fake Categories
	categories := []string{"Electronic", "Home", "Appliance", "Kitchen", "Food"}
	for i := 0; i < 5; i++ {
		category := model.Category{
			Name: categories[i],
		}
		database.DB.Create(&category)
	}

	// Generate user cart
	var customers []model.Customer
	var products []model.Product

	database.DB.Find(&customers)
	database.DB.Find(&products)

	for i := 0; i < 20; i++ {
		cart := model.Carts{
			CustomerID: customers[fake.IntBetween(0, 2)].ID,
			ProductID:  products[fake.IntBetween(0, 9)].ID,
			Quantity:   fake.IntBetween(0, 9),
		}
		database.DB.Create(&cart)
	}

	// Generate Product Category
	var category []model.Category
	database.DB.Find(&category)
	for i := 0; i < 10; i++ {
		productCategory := model.ProductCategories{
			ProductID:  products[i].ID,
			CategoryID: category[fake.UIntBetween(0, 4)].ID,
		}
		database.DB.Create(&productCategory)
	}

	fmt.Println("Seed Done ╰(*°▽°*)╯")
}

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}
	database.OpenDatabaseConnection(config)
	Migrate()
	Seed()
}
