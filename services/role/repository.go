package role

import (
	"api-foodmarket/helper"
	"gorm.io/gorm"
)

type Repository interface {
	Save(role Role) Role
	FindAll() (roles []Role)
	FindById(Id int) (role Role)
	Delete(role Role)
	FindByName(RoleName string) (role Role)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Save(role Role) Role {
	err := r.db.Create(&role).Error
	helper.PanicIfError(err)

	return role
}

func (r *repository) FindAll() (roles []Role) {
	err := r.db.Find(&roles).Error
	helper.PanicIfError(err)

	return
}

func (r *repository) FindById(Id int) (role Role) {
	err := r.db.Find(&role).Error
	helper.PanicIfError(err)

	return
}

func (r *repository) Delete(role Role) {
	err := r.db.Delete(&role).Error
	helper.PanicIfError(err)

	return
}

func (r *repository) FindByName(RoleName string) (role Role) {
	err := r.db.Where("name = ?", RoleName).Find(&role).Error
	helper.PanicIfError(err)

	return role
}
