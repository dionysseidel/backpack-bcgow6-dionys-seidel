package users

import "github.com/dionysseidel/backpack-bcgow6-dionys-seidel/go-testing/go-web/internal/domain"

type Service interface {
	Delete(id int) error
	GetAll() ([]domain.User, error)
	Store(name string, active bool, age int) (domain.User, error)
	Update(id int, name string, active bool, age int) (domain.User, error)
	UpdateNameAndAge(id int, name string, age int) (domain.User, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

func (s *service) GetAll() ([]domain.User, error) {
	users, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *service) Store(name string, isActive bool, age int) (domain.User, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return domain.User{}, err
	}

	user, err := s.repository.Store(lastID+1, name, isActive, age)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (s *service) Update(id int, name string, isActive bool, age int) (domain.User, error) {
	return s.repository.Update(id, name, isActive, age)
}

func (s *service) UpdateNameAndAge(id int, name string, age int) (domain.User, error) {
	return s.repository.UpdateNameAndAge(id, name, age)
}
