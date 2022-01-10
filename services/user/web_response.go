package user

import (
	"time"
)

type FormatUser struct {
	ID        int          `json:"id"`
	Name      string       `json:"name"`
	Email     string       `json:"email"`
	Phone     string       `json:"phone"`
	Address   string       `json:"address"`
	City      string       `json:"city"`
	Photo     string       `json:"photo"`
	Roles     []FormatRole `json:"role"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

type FormatRole struct {
	Name string `json:"name"`
}

func UserFormat(user User) FormatUser {

	formatRoles := []FormatRole{}
	roles := user.Roles
	for _, role := range roles {
		formatRole := FormatRole{Name: role.Role.Display}
		formatRoles = append(formatRoles, formatRole)
	}
	return FormatUser{
		ID:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		Address:   user.Address,
		City:      user.City,
		Photo:     user.Photo,
		Roles:     formatRoles,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func UsersFormat(users []User) []FormatUser {
	formats := []FormatUser{}

	for _, user := range users {
		format := UserFormat(user)
		formats = append(formats, format)
	}

	return formats
}
