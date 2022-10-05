package users

import "fmt"

type User struct {
	ID       int    `json:id`
	Name     string `json:"nombre" binding:"required"`
	IsActive bool   `json:"estaActive"`
	Age      int    `json:"edad" binding:"required"`
}

type Users []User

var usersSlice Users

type Repository interface {
	Delete(id int) error
	GetAll() ([]User, error)
	LastID() (int, error)
	Store(id int, name string, isActive bool, age int) (User, error)
	Update(id int, name string, active bool, age int) (User, error)
	UpdateNameAndAge(id int, name string, age int) (User, error)
}

type repository struct {
	// filepath string
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Delete(id int) error {
	deleted := false
	var index int
	for i := range usersSlice {
		if usersSlice[i].ID == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("usuarie %d no encontrade", id)
	}
	usersSlice = append(usersSlice[:index], usersSlice[index+1:]...)
	return nil
}

func (r *repository) GetAll() ([]User, error) {
	return usersSlice, nil
}

func (r *repository) LastID() (int, error) {
	return len(usersSlice) + 1, nil
}

func (r *repository) Store(id int, name string, isActive bool, age int) (User, error) {
	userToCreate := User{
		ID:       id,
		Name:     name,
		IsActive: isActive,
		Age:      age,
	}
	usersSlice = append(usersSlice, userToCreate)
	// fmt.Println("userSlice in Store method in repository", usersSlice)
	return userToCreate, nil
}

func (r *repository) Update(id int, name string, isActive bool, age int) (User, error) {
	userToUpdate := User{
		Name:     name,
		IsActive: isActive,
		Age:      age,
	}
	isUpdated := false
	for i := range usersSlice {
		if usersSlice[i].ID == id {
			userToUpdate.ID = id
			usersSlice[i] = userToUpdate
			isUpdated = true
		}
	}
	if !isUpdated {
		return User{}, fmt.Errorf("User %d no encontrade", id)
	}
	return userToUpdate, nil
}

func (r *repository) UpdateNameAndAge(id int, name string, age int) (User, error) {
	var userToReturn User
	updated := false
	for _, userToUpdate := range usersSlice {
		if userToUpdate.ID == id {
			userToUpdate.Name = name
			userToUpdate.Age = age
			userToReturn = userToUpdate
		}
	}
	if !updated {
		return User{}, fmt.Errorf("Usuarie %d no encontrade", id)
	}
	return userToReturn, nil
}
