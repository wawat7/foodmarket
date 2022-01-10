package main

import (
	"api-foodmarket/app"
	"api-foodmarket/services/role"
	"api-foodmarket/services/user"
	"fmt"
	"gorm.io/gorm"
	"time"
)

func main() {

	fmt.Println("seeder running..")
	configuration := app.New()
	db := app.NewDB(configuration)

	roleSeeder(db)
	userSeeder(db)
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
