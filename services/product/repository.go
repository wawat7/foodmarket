package product

import (
	"api-foodmarket/helper"
	"gorm.io/gorm"
)

type Repository interface {
	Create(product Product) Product
	FindAll() (products []Product)
	FindById(Id int) (product Product)
	Update(product Product) Product
	Delete(product Product)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Create(product Product) Product {
	err := r.db.Create(&product).Error
	helper.PanicIfError(err)

	return product
}

func (r *repository) FindAll() (products []Product) {
	err := r.db.Find(&products).Error
	helper.PanicIfError(err)

	return
}

func (r *repository) FindById(Id int) (product Product) {
	err := r.db.Where("id = ?", Id).Find(&product).Error
	helper.PanicIfError(err)

	return
}

func (r *repository) Update(product Product) Product {
	err := r.db.Save(&product).Error
	helper.PanicIfError(err)

	return product
}

func (r *repository) Delete(product Product) {
	err := r.db.Delete(&product).Error
	helper.PanicIfError(err)

	return
}

