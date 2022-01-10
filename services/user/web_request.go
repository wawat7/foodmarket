package user

type GetUserDetailParam struct {
	Id int `uri:"id" binding:"required"`
}

type CreateUserRequest struct {
	Name     string `form:"name" binding:"required"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
	Phone    string `form:"phone" binding:"required"`
	Address  string `form:"address" binding:"required"`
	City     string `form:"city" binding:"required"`
	Photo    string `form:"photo" binding:"required"`
}

type UpdateUserRequest struct {
	Name    string `form:"name"`
	Phone   string `form:"phone"`
	Address string `form:"address"`
	City    string `form:"city"`
	Photo   string `form:"photo"`
}
