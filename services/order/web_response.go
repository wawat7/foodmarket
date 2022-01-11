package order

import (
	"api-foodmarket/helper"
	"api-foodmarket/services/product"
	"api-foodmarket/services/user"
	"encoding/json"
	"time"
)

type FormatOrder struct {
	Id        int                  `json:"id"`
	Code      string               `json:"code"`
	Status    string               `json:"status"`
	Total     int                  `json:"total"`
	UserId    int                  `json:"user_id"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
	Products  []FormatOrderProduct `json:"products"`
	Histories []FormatOrderHistory `json:"histories"`
	User      user.FormatUser      `json:"user"`
}

type FormatOrderProduct struct {
	Id       int    `json:"id"`
	Quantity int    `json:"quantity"`
	SubTotal int    `json:"subTotal"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Image    string `json:"image"`
}

type FormatOrderHistory struct {
	Status    string    `json:"status"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}

func productFormat(orderProduct OrderProduct) FormatOrderProduct {
	var item product.Product
	err := json.Unmarshal([]byte(orderProduct.ProductInfo), &item)
	helper.PanicIfError(err)
	return FormatOrderProduct{
		Id:       item.Id,
		Quantity: orderProduct.Quantity,
		SubTotal: orderProduct.SubTotal,
		Name:     item.Name,
		Price:    item.Price,
		Image:    item.Image,
	}
}

func historyFormat(history OrderHistory) FormatOrderHistory {
	return FormatOrderHistory{
		Status:    history.Status,
		CreatedBy: history.CreatedBy,
		CreatedAt: history.CreatedAt,
	}
}

func OrderFormat(order Order) FormatOrder {
	var formatProducts []FormatOrderProduct
	var formatHistories []FormatOrderHistory
	var userData user.User

	products := order.OrderProduct
	for _, product := range products {
		formatProduct := productFormat(product)
		formatProducts = append(formatProducts, formatProduct)
	}

	histories := order.OrderHistory
	for _, history := range histories {
		formatHistory := historyFormat(history)
		formatHistories = append(formatHistories, formatHistory)
	}
	err := json.Unmarshal([]byte(order.UserInfo), &userData)
	helper.PanicIfError(err)

	formatUser := user.UserFormat(userData)

	format := FormatOrder{
		Id:        order.Id,
		Code:      order.Code,
		Status:    order.Status,
		Total:     order.Total,
		UserId:    order.UserId,
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
		Products:  formatProducts,
		Histories: formatHistories,
		User:      formatUser,
	}

	return format
}

func OrdersFormat(orders []Order) []FormatOrder {
	formats := []FormatOrder{}

	for _, order := range orders {
		format := OrderFormat(order)
		formats = append(formats, format)
	}

	return formats
}
