package users

type Service interface {
	GetAll() ([]User, error)
	Store(name string, active bool, age int) (User, error)
	Update(id int, name string, active bool, age int) (User, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]User, error) {
	users, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *service) Store(name string, active bool, age int) (User, error) {

}

func (s *service) Update(id int, name string, active bool, age int) (User, error) {}
