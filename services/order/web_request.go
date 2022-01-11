package order

type CreateOrderRequest struct {
	UserId    int `json:"user_id" binding:"required"`
	ProductId int `json:"product_id" binding:"required"`
	Quantity  int `json:"quantity" binding:"required"`
}

type ProductQuantityData struct {
	ProductId int
	Quantity  int
}

type GetOrderDetail struct {
	Id int `uri:"id" binding:"required"`
}

type UpdateStatusOrderRequest struct {
	UserId int    `json:"user_id" binding:"required"`
	Status string `json:"status" binding:"required"`
}
