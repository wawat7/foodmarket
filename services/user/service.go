package user

import (
	"api-foodmarket/services/role"
)

type Service interface {
	Create(user User, RoleName string) User
	FindAll() (users []User)
	FindById(Id int) (user User)
	Update(user User) User
	Delete(user User)
}

type service struct {
	repository     Repository
	roleRepository role.Repository
}

func NewService(repository Repository, roleRepository role.Repository) *service {
	return &service{repository: repository, roleRepository: roleRepository}
}

func (s *service) FindAll() (users []User) {
	users = s.repository.FindAll()
	return
}

func (s *service) Create(user User, RoleName string) User {
	if RoleName == "" {
		RoleName = role.Customer
	}

	roleData := s.roleRepository.FindByName(RoleName)
	user = s.repository.Save(user)
	_ = s.repository.SaveRole(user, roleData)

	return user
}

func (s *service) FindById(Id int) (user User) {
	user = s.repository.FindById(Id)
	return
}

func (s *service) Update(user User) User {
	user = s.repository.Update(user)

	return user
}

func (s *service) Delete(user User) {
	s.repository.Delete(user)

	return
}
