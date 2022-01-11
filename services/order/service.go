package order

import (
	"api-foodmarket/helper"
	"api-foodmarket/services/product"
	"api-foodmarket/services/user"
	"gorm.io/gorm"
	"time"
)

type Service interface {
	Create(user user.User, productData []ProductQuantityData) Order
	FindById(Id int) (order Order)
	FindAll() (orders []Order)
	update(order Order, user user.User) Order
}

type service struct {
	repository     Repository
	productService product.Service
}

func NewService(repository Repository, productService product.Service) *service {
	return &service{repository: repository, productService: productService}
}

func (s *service) Create(user user.User, productData []ProductQuantityData) Order {
	var items []OrderProduct

	userJson := helper.ConvertDataToJson(user)
	order := Order{
		UserInfo: userJson,
		UserId:   user.Id,
		Code:     "TRX-123",
		Total:    0,
		Status:   STATUS_PENDING,
	}

	order = s.repository.Create(order)
	Total := 0
	for _, data := range productData {
		product := s.productService.FindById(data.ProductId)
		subTotal := product.Price * data.Quantity
		Total += subTotal
		productJson := helper.ConvertDataToJson(product)

		item := OrderProduct{
			OrderId:     order.Id,
			ProductInfo: productJson,
			ProductId:   data.ProductId,
			Quantity:    data.Quantity,
			SubTotal:    subTotal,
		}

		items = append(items, item)
	}

	//update Total
	order.Total = Total
	order = s.repository.update(order)

	_ = s.repository.SaveProductBatch(items)

	history := OrderHistory{
		OrderId:   order.Id,
		Status:    order.Status,
		CreatedBy: user.Name,
	}

	_ = s.repository.SaveHistory(history)

	return order
}

func (s *service) FindById(Id int) (order Order) {
	order = s.repository.FindById(Id)

	return
}

func (s *service) FindAll() (orders []Order) {
	orders = s.repository.FindAll()

	return orders
}

func (s *service) update(order Order, user user.User) Order {
	order = s.repository.update(order)

	history := OrderHistory{
		OrderId:   order.Id,
		Status:    order.Status,
		CreatedBy: user.Name,
	}
	_ = s.repository.SaveHistory(history)

	return order
}

func mapProductToOrderProduct(product product.Product) OrderProduct {
	return OrderProduct{
		Id:          0,
		OrderId:     0,
		ProductInfo: "",
		ProductId:   product.Id,
		Quantity:    0,
		SubTotal:    0,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   gorm.DeletedAt{},
	}
}
