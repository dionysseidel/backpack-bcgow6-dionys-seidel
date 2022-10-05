package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dionysseidel/backpack-bcgow6-dionys-seidel/go-web/internal/users"
	"github.com/gin-gonic/gin"
)

type userToCreate struct {
	Name   string `json:"nombre" binding:"required"`
	Active bool   `json:"active"`
	Age    int    `json:"edad" binding:"required"`
}

type User struct {
	service users.Service
}

var stringifiedFalseBool = strconv.FormatBool(false)
var stringifiedTrueBool = strconv.FormatBool(true)

func NewUser(u users.Service) *User {
	return &User{
		service: u,
	}
}

func (userInMemory *User) CreateUser() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		token := ginContext.Request.Header.Get("token") /*ginContext.GetHeader("token") */
		if token != "123456" || token == "" {
			ginContext.JSON(http.StatusUnauthorized, gin.H{"error": "no tiene permisos para realizar la petición solicitada"})
			return
		}
		var userToCreate userToCreate
		if err := ginContext.ShouldBindJSON(&userToCreate); err != nil {
			if userToCreate.Name == "" {
				returnErrorEmptyField("Nombre")
				return
			} else if userToCreate.Age == 0 {
				returnErrorEmptyField("Edad")
				return
			}
			return
		}
		// userToCreate.ID = len(usersSlice) + 1
		// usersSlice = append(usersSlice, userToCreate)
		// fmt.Println("userSlice in CreateUser", usersSlice)
		userToReturn, err := userInMemory.service.Store(userToCreate.Name, userToCreate.Active, userToCreate.Age)
		if err != nil {
			ginContext.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ginContext.JSON(200, userToReturn)
	}
}

func returnErrorEmptyField(fieldName string) gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		ginContext.JSON(400, gin.H{
			"error": fmt.Sprintf("el campo %s es requerido", fieldName),
		})
	}
}

func (u *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" || token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "no tiene permisos para realizar la petición solicitada"})
			return
		}
		allUsers, err := u.service.GetAll()
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
				case user.Active:
					filteredUsers = append(filteredUsers, user)
				default:
					switch ctx.Query("active") == strconv.FormatBool(false) {
					case !user.Active:
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
