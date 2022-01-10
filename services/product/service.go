package product

type Service interface {
	Create(product Product) Product
	FindAll() (products []Product)
	FindById(Id int) (product Product)
	Update(product Product) Product
	Delete(product Product)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) Create(product Product) Product {
	product = s.repository.Create(product)
	return product
}

func (s *service) FindAll() (products []Product) {
	products = s.repository.FindAll()
	return
}

func (s *service) FindById(Id int) (product Product) {
	product = s.repository.FindById(Id)
	return
}

func (s *service) Update(product Product) Product {
	product = s.repository.Update(product)
	return product
}

func (s *service) Delete(product Product) {
	s.repository.Delete(product)

	return
}
