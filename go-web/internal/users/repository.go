package users

import (
	"fmt"

	"github.com/dionysseidel/backpack-bcgow6-dionys-seidel/internal/domains"
	"github.com/dionysseidel/backpack-bcgow6-dionys-seidel/pkg/store"
)

type Users []domains.User

var usersSlice Users

type Repository interface {
	Delete(id int) error
	GetAll() ([]domains.User, error)
	LastID() (int, error)
	Store(id int, name string, isActive bool, age int) (domains.User, error)
	Update(id int, name string, active bool, age int) (domains.User, error)
	UpdateNameAndAge(id int, name string, age int) (domains.User, error)
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
	var usersInFile []domains.User
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

func (r *repository) GetAll() ([]domains.User, error) {
	var usersFromFile []domains.User
	if err := r.db.Read(&usersFromFile); err != nil {
		return usersFromFile, err
	}
	// usersSlice = append(usersSlice, usersFromFile...)
	return usersFromFile, nil
}

func (r *repository) LastID() (int, error) {
	var usersFromFile []domains.User
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

func (r *repository) Store(id int, name string, isActive bool, age int) (domains.User, error) {
	var usersInFile []domains.User
	if err := r.db.Read(&usersInFile); err != nil {
		return domains.User{}, err
	}
	// usersSlice = append(usersSlice, usersInFile...)
	userToCreate := domains.User{
		ID:       id,
		Name:     name,
		IsActive: isActive,
		Age:      age,
	}
	usersInFile = append(usersInFile, userToCreate)
	// fmt.Println("userSlice in Store method in repository", usersSlice)
	if err := r.db.Write(usersInFile); err != nil {
		return domains.User{}, err
	}
	return userToCreate, nil
}

func (r *repository) Update(id int, name string, isActive bool, age int) (domains.User, error) {
	userToUpdate := domains.User{
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
		return domains.User{}, fmt.Errorf("User %d no encontrade", id)
	}
	return userToUpdate, nil
}

func (r *repository) UpdateNameAndAge(id int, name string, age int) (domains.User, error) {
	var userToReturn domains.User
	updated := false
	for _, userToUpdate := range usersSlice {
		if userToUpdate.ID == id {
			userToUpdate.Name = name
			userToUpdate.Age = age
			userToReturn = userToUpdate
		}
	}
	if !updated {
		return domains.User{}, fmt.Errorf("Usuarie %d no encontrade", id)
	}
	return userToReturn, nil
}
