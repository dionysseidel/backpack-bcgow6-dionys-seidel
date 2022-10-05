package users

type User struct {
	ID     int    `json:id`
	Name   string `json:"nombre" binding:"required"`
	Active bool   `json:"active"`
	Age    int    `json:"edad" binding:"required"`
}

type Users []User

var usersSlice Users

type Repository interface {
	GetAll() ([]User, error)
	Store(name string, isActive bool, age int) (User, error)
	LastID() (int, error)
	Update(id int, name string, active bool, age int) (User, error)
}

type repository struct {
	// filepath string
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]User, error) {
	return usersSlice, nil
}

func (r *repository) LastID() (int, error) {}

func (r *repository) Store(name string, isActive bool, age int) (User, error) {
	userToCreate := User{
		Name:   name,
		Active: isActive,
		Age:    age,
	}
	userToCreate.ID = len(usersSlice) + 1
	usersSlice = append(usersSlice, userToCreate)
	// fmt.Println("userSlice in Store method in repository", usersSlice)
	return userToCreate, nil
}

func (r *repository) Update(id int, name string, active bool, age int) (User, error) {}
