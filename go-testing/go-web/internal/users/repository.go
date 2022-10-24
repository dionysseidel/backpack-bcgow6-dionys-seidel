package users

import (
	"fmt"

	"github.com/dionysseidel/backpack-bcgow6-dionys-seidel/go-testing/go-web/internal/domain"
	"github.com/dionysseidel/backpack-bcgow6-dionys-seidel/go-testing/go-web/pkg/store"
)

type Users []domain.User

var usersSlice Users

type Repository interface {
	Delete(id int) error
	GetAll() ([]domain.User, error)
	LastID() (int, error)
	Store(id int, name string, isActive bool, age int) (domain.User, error)
	Update(id int, name string, active bool, age int) (domain.User, error)
	UpdateNameAndAge(id int, name string, age int) (domain.User, error)
}

type repository struct {
	// filepath string
	db store.Storage
}

func NewRepository(db store.Storage) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Delete(id int) error {
	deleted := false

	var usersInFile []domain.User
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

func (r *repository) GetAll() ([]domain.User, error) {
	var usersFromFile []domain.User
	if err := r.db.Read(&usersFromFile); err != nil {
		return usersFromFile, err
	}
	// usersSlice = append(usersSlice, usersFromFile...)
	return usersFromFile, nil
}

func (r *repository) LastID() (int, error) {
	var usersFromFile []domain.User
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

func (r *repository) Store(id int, name string, isActive bool, age int) (domain.User, error) {
	var usersInFile []domain.User
	if err := r.db.Read(&usersInFile); err != nil {
		return domain.User{}, err
	}

	userToCreate := domain.User{
		ID:       id,
		Name:     name,
		IsActive: isActive,
		Age:      age,
	}

	usersInFile = append(usersInFile, userToCreate)

	if err := r.db.Write(usersInFile); err != nil {
		return domain.User{}, err
	}

	return userToCreate, nil
}

func (r *repository) Update(id int, name string, isActive bool, age int) (domain.User, error) {
	userToUpdate := domain.User{
		Name:     name,
		IsActive: isActive,
		Age:      age,
	}
	isUpdated := false

	var usersInFile []domain.User
	if err := r.db.Read(&usersInFile); err != nil {
		return domain.User{}, err
	}

	for i := range usersInFile {
		if usersInFile[i].ID == id {
			userToUpdate.ID = id
			usersInFile[i] = userToUpdate
			isUpdated = true
		}
	}
	if !isUpdated {
		return domain.User{}, fmt.Errorf("User %d no encontrade", id)
	}

	if err := r.db.Write(usersInFile); err != nil {
		return domain.User{}, err
	}

	return userToUpdate, nil
}

func (r *repository) UpdateNameAndAge(id int, name string, age int) (domain.User, error) {
	var userToReturn domain.User
	updated := false
	for _, userToUpdate := range usersSlice {
		if userToUpdate.ID == id {
			userToUpdate.Name = name
			userToUpdate.Age = age
			userToReturn = userToUpdate
		}
	}
	if !updated {
		return domain.User{}, fmt.Errorf("Usuarie %d no encontrade", id)
	}
	return userToReturn, nil
}
