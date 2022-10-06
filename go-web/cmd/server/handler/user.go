package handler

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/dionysseidel/backpack-bcgow6-dionys-seidel/go-web/internal/users"
	"github.com/gin-gonic/gin"
)

type userToCreate struct {
	Name     string `json:"nombre" binding:"required"`
	IsActive bool   `json:"estaActive"`
	Age      int    `json:"edad" binding:"required"`
}

type UserController struct {
	service users.Service
}

var stringifiedFalseBool = strconv.FormatBool(false)
var stringifiedTrueBool = strconv.FormatBool(true)

func NewUserController(u users.Service) *UserController {
	return &UserController{
		service: u,
	}
}

func (controller *UserController) CreateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if isAuthenticated := validateToken(ctx); !isAuthenticated {
			return
		}
		var userToCreate userToCreate
		if err := ctx.ShouldBindJSON(&userToCreate); err != nil {
			if userToCreate.Name == "" {
				returnErrorEmptyField("Nombre")
				return
			} else if userToCreate.Age == 0 {
				returnErrorEmptyField("Edad")
				return
			}
			return
		}
		userToReturn, err := controller.service.Store(userToCreate.Name, userToCreate.IsActive, userToCreate.Age)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, userToReturn)
	}
}

func (controller *UserController) Delete(ctx *gin.Context) {
	if isAuthenticated := validateToken(ctx); !isAuthenticated {
		return
	}
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid ID"})
		return
	}
	err = controller.service.Delete(int(id))
	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": fmt.Sprintf("Le usuarie %d ha sido eliminade", id)})
}

func (controller *UserController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if isAuthenticated := validateToken(ctx); !isAuthenticated {
			return
		}
		allUsers, err := controller.service.GetAll()
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
		}
		_, isPresent := ctx.GetQuery("active")
		if isPresent {
			var filteredUsers []users.User
			for _, user := range allUsers {
				switch ctx.Query("active") == strconv.FormatBool(true) {
				case user.IsActive:
					filteredUsers = append(filteredUsers, user)
				default:
					switch ctx.Query("active") == strconv.FormatBool(false) {
					case !user.IsActive:
						filteredUsers = append(filteredUsers, user)
					}
				}
			}
			ctx.JSON(http.StatusOK, filteredUsers)
		} else {
			ctx.JSON(http.StatusOK, allUsers)
		}
	}
}

func (controller *UserController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if isAuthenticated := validateToken(ctx); !isAuthenticated {
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}
		var userToCreate userToCreate
		if err := ctx.ShouldBindJSON(&userToCreate); err != nil {
			if userToCreate.Name == "" {
				returnErrorEmptyField("Nombre")
				return
			} else if userToCreate.Age == 0 {
				returnErrorEmptyField("Edad")
				return
			}
			return
		}
		userToUpdate, err := controller.service.Update(int(id), userToCreate.Name, userToCreate.IsActive, userToCreate.Age)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, userToUpdate)
	}
}

func (controller *UserController) UpdateNameAndAge(ctx *gin.Context) {
	if isAuthenticated := validateToken(ctx); !isAuthenticated {
		return
	}
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid ID"})
		return
	}
	var userToCreate userToCreate
	if err := ctx.ShouldBindJSON(&userToCreate); err != nil {
		if userToCreate.Name == "" {
			returnErrorEmptyField("Nombre")
			return
		} else if userToCreate.Age == 0 {
			returnErrorEmptyField("Edad")
			return
		}
		return
	}
	userToUpdate, err := controller.service.Update(int(id), userToCreate.Name, userToCreate.IsActive, userToCreate.Age)
	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, userToUpdate)
}

func returnErrorEmptyField(fieldName string) gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		ginContext.JSON(400, gin.H{
			"error": fmt.Sprintf("el campo %s es requerido", fieldName),
		})
	}
}

func validateToken(ctx *gin.Context) (authenticated bool) {
	token := ctx.GetHeader("token") /*ctx.Request.Header.Get("token")*/
	// ... Me imagino que acá querría persistir le token en el archivo .env
	if token != os.Getenv("TOKEN") || token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "no tiene permisos para realizar la petición solicitada"})
		authenticated = false
		return
	}
	return true
}
