package role

type Service interface {
	Create(role Role) Role
	FindAll() (roles []Role)
	FindById(Id int) (role Role)
	Delete(role Role)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) Create(role Role) Role {
	role = s.repository.Save(role)
	return role
}

func (s *service) FindAll() (roles []Role) {
	roles = s.repository.FindAll()
	return
}

func (s *service) FindById(Id int) (role Role) {
	role = s.repository.FindById(Id)
	return
}

func (s *service) Delete(role Role) {
	s.repository.Delete(role)

	return
}
