package user

import (
	"api-foodmarket/helper"
	"api-foodmarket/services/role"
	"errors"
	"gorm.io/gorm"
)

type Repository interface {
	Save(user User) User
	FindAll() []User
	FindById(Id int) User
	Update(user User) User
	Delete(user User)
	FindAllPagination(page int, pageSize int) []User
	SaveRole(user User, role2 role.Role) UserRole
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (repository *repository) Save(user User) User {
	err := repository.db.Create(&user).Error
	helper.PanicIfError(err)

	return user
}

func (repository *repository) FindAll() (users []User) {
	err := repository.db.Preload("Roles.Role").Find(&users).Error
	helper.PanicIfError(err)

	return
}

func (repository *repository) FindById(Id int) (user User) {
	err := repository.db.Preload("Roles.Role").Where("id = ?", Id).Find(&user).Error
	helper.PanicIfError(err)

	if user.Id == 0 {
		helper.PanicIfError(errors.New("user not found"))
	}

	return
}

func (repository *repository) Update(user User) User {
	err := repository.db.Save(&user).Error
	helper.PanicIfError(err)

	return user
}

func (repository *repository) Delete(user User) {
	err := repository.db.Delete(&user).Error
	helper.PanicIfError(err)

	return
}

func (repository *repository) FindAllPagination(page int, pageSize int) []User {
	//TODO implement me
	panic("implement me")
}

func (repository *repository) SaveRole(user User, role role.Role) UserRole {
	userRole := UserRole{
		UserId: user.Id,
		RoleId: role.Id,
	}
	err := repository.db.Create(&userRole).Error
	helper.PanicIfError(err)

	return userRole
}
