package domain

type User struct {
	ID       int    `json:id`
	Name     string `json:"nombre" binding:"required"`
	IsActive bool   `json:"estaActive"`
	Age      int    `json:"edad" binding:"required"`
}

// Prof añade omitempty a los atributos
