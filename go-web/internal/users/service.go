package users

import "github.com/dionysseidel/backpack-bcgow6-dionys-seidel/internal/domains"

type Service interface {
	Delete(id int) error
	GetAll() ([]domains.User, error)
	Store(name string, active bool, age int) (domains.User, error)
	Update(id int, name string, active bool, age int) (domains.User, error)
	UpdateNameAndAge(id int, name string, age int) (domains.User, error)
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

func (s *service) GetAll() ([]domains.User, error) {
	users, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *service) Store(name string, isActive bool, age int) (domains.User, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return domains.User{}, err
	}

	user, err := s.repository.Store(lastID+1, name, isActive, age)
	if err != nil {
		return domains.User{}, err
	}

	return user, nil
}

func (s *service) Update(id int, name string, isActive bool, age int) (domains.User, error) {
	return s.repository.Update(id, name, isActive, age)
}

func (s *service) UpdateNameAndAge(id int, name string, age int) (domains.User, error) {
	return s.repository.UpdateNameAndAge(id, name, age)
}
