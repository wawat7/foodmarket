package order

import (
	"api-foodmarket/helper"
	"errors"
	"gorm.io/gorm"
)

type Repository interface {
	Create(order Order) Order
	FindById(Id int) (order Order)
	FindAll() (orders []Order)
	update(order Order) Order
	SaveHistory(history OrderHistory) OrderHistory
	SaveProductBatch(products []OrderProduct) []OrderProduct
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Create(order Order) Order {
	err := r.db.Create(&order).Error
	helper.PanicIfError(err)

	return order
}

func (r *repository) FindById(Id int) (order Order) {
	err := r.db.Preload("OrderProduct").Preload("OrderHistory").Where("id = ?", Id).Find(&order).Error
	helper.PanicIfError(err)

	if order.Id == 0 {
		helper.PanicIfError(errors.New("order not found"))
	}

	return
}

func (r *repository) FindAll() (orders []Order) {
	err := r.db.Preload("OrderProduct").Preload("OrderHistory").Find(&orders).Error
	helper.PanicIfError(err)

	return
}

func (r *repository) update(order Order) Order {
	err := r.db.Save(&order).Error
	helper.PanicIfError(err)

	return order
}

func (r *repository) SaveHistory(history OrderHistory) OrderHistory {
	err := r.db.Create(&history).Error
	helper.PanicIfError(err)

	return history
}

func (r *repository) SaveProductBatch(products []OrderProduct) []OrderProduct {
	err := r.db.Create(&products).Error
	helper.PanicIfError(err)

	return products
}
