package users

type Service interface {
	Delete(id int) error
	GetAll() ([]User, error)
	Store(name string, active bool, age int) (User, error)
	Update(id int, name string, active bool, age int) (User, error)
	UpdateNameAndAge(id int, name string, age int) (User, error)
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

func (s *service) GetAll() ([]User, error) {
	users, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *service) Store(name string, isActive bool, age int) (User, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return User{}, err
	}

	user, err := s.repository.Store(lastID+1, name, isActive, age)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (s *service) Update(id int, name string, isActive bool, age int) (User, error) {
	return s.repository.Update(id, name, isActive, age)
}

func (s *service) UpdateNameAndAge(id int, name string, age int) (User, error) {
	return s.repository.UpdateNameAndAge(id, name, age)
}
