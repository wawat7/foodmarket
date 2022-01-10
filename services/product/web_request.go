package product

type GetProductDetail struct {
	Id int `uri:"id" binding:"required"`
}

type CreateProductRequest struct {
	Name        string   `form:"name" binding:"required"`
	Description string   `form:"description" binding:"required"`
	Ingredient  []string `form:"ingredient[]" binding:"required"`
	Price       int      `form:"price" binding:"required"`
	Type        string   `form:"type" binding:"required"`
	Image       string   `form:"image" binding:"required"`
}

type UpdateProductRequest struct {
	Name        string   `form:"name"`
	Description string   `form:"description"`
	Ingredient  []string `form:"ingredient[]"`
	Price       int      `form:"price"`
	Type        string   `form:"type"`
	Image       string   `form:"image"`
}
