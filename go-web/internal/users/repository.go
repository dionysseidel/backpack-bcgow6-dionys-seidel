package users

import (
	"fmt"

	"github.com/dionysseidel/backpack-bcgow6-dionys-seidel/pkg/store"
)

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
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Delete(id int) error {
	deleted := false
	var usersInFile []User
	if err := r.db.Read(&usersInFile); err != nil {
		return err
	}
	var index int
	for i := range usersInFile {
		if usersInFile[i].ID == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("usuarie %d no encontrade", id)
	}
	usersInFile = append(usersInFile[:index], usersInFile[index+1:]...)
	if err := r.db.Write(usersInFile); err != nil {
		return err
	}
	return nil
}

func (r *repository) GetAll() ([]User, error) {
	var usersFromFile []User
	if err := r.db.Read(&usersFromFile); err != nil {
		return usersFromFile, err
	}
	// usersSlice = append(usersSlice, usersFromFile...)
	return usersFromFile, nil
}

func (r *repository) LastID() (int, error) {
	var usersFromFile []User
	var usersInFile int
	if err := r.db.Read(&usersFromFile); err != nil {
		return len(usersSlice), err
	}
	usersInFile = len(usersFromFile)
	totalUsers := len(usersSlice) + usersInFile
	if totalUsers == 0 {
		return 0, nil
	}
	return totalUsers, nil
}

func (r *repository) Store(id int, name string, isActive bool, age int) (User, error) {
	var usersInFile []User
	if err := r.db.Read(&usersInFile); err != nil {
		return User{}, err
	}
	// usersSlice = append(usersSlice, usersInFile...)
	userToCreate := User{
		ID:       id,
		Name:     name,
		IsActive: isActive,
		Age:      age,
	}
	usersInFile = append(usersInFile, userToCreate)
	// fmt.Println("userSlice in Store method in repository", usersSlice)
	if err := r.db.Write(usersInFile); err != nil {
		return User{}, err
	}
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
