package app

import (
	"api-foodmarket/helper"
	"api-foodmarket/services/product"
	"api-foodmarket/services/role"
	"api-foodmarket/services/user"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"time"
)

func SeederRun() {

	fmt.Println("seeder running..")
	configuration := New()
	db := NewDB(configuration)

	//roleSeeder(db)
	//userSeeder(db)
	productSeeder(db)
}

func roleSeeder(db *gorm.DB) {
	fmt.Println("running seeder roles")

	roleRepository := role.NewRepository(db)
	roleService := role.NewService(roleRepository)

	roles := role.List()

	for name, display := range roles {
		roleData := role.Role{
			Name:    name,
			Display: display,
		}
		_ = roleService.Create(roleData)

	}

	fmt.Println("ended seeder roles")
}

func userSeeder(db *gorm.DB) {
	fmt.Println("running seeder user")

	userRepository := user.NewRepository(db)
	roleRepository := role.NewRepository(db)
	userService := user.NewService(userRepository, roleRepository)

	userData := user.User{
		Name:      "Administrator",
		Email:     "admin@mailinator.com",
		Password:  "password",
		Phone:     "081918827384",
		Address:   "Jalan example no.31",
		City:      "Bandung",
		Photo:     "",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	userData = userService.Create(userData, role.Admin)
	fmt.Println("ended seeder user")

}

func productSeeder(db *gorm.DB) {
	fmt.Println("running seeder product")

	productRepository := product.NewRepository(db)
	productService := product.NewService(productRepository)

	ingredients := []string{"terigu", "cacing", "ketombe"}
	ingredientJson, err := json.Marshal(ingredients)
	helper.PanicIfError(err)

	productData := product.Product{
		Name:        "Product 1",
		Description: "Ini adalah deskripsi product 1",
		Ingredient:  string(ingredientJson),
		Price:       20000,
		Rate:        0,
		Type:        product.NEW,
		Image:       "",
	}

	_ = productService.Create(productData)

	fmt.Println("ended seeder product")
}
