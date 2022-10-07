package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dionysseidel/backpack-bcgow6-dionys-seidel/internal/domains"
	"github.com/dionysseidel/backpack-bcgow6-dionys-seidel/internal/users"
	"github.com/dionysseidel/backpack-bcgow6-dionys-seidel/pkg/web"
	"github.com/gin-gonic/gin"
)

type userRequest struct {
	Name     string `json:"nombre" binding:"required"`
	IsActive bool   `json:"estaActive"`
	Age      int    `json:"edad" binding:"required"`
}

type userRequestPatch struct {
	Name string `json:"nombre" binding:"required"`
	Age  int    `json:"edad" binding:"required"`
}

type UserController struct {
	service users.Service
}

func NewUserController(u users.Service) *UserController {
	return &UserController{
		service: u,
	}
}

var stringifiedFalseBool = strconv.FormatBool(false)
var stringifiedTrueBool = strconv.FormatBool(true)

// StoreUser godoc
// @Summary     Store user
// @Tags        Users
// @Description store users
// @Accept      json
// @Produce     json
// @Param       token header   string      true "token requeridx"
// @Param       user  body     userRequest true "User to store"
// @Success     200   {object} web.Response
// @Failure     400   {object} web.Response
// @Failure     401   {object} web.Response
// @Failure     409   {object} web.Response
// @Router      /users [POST]
func (controller *UserController) CreateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userToCreate userRequest
		if err := ctx.ShouldBindJSON(&userToCreate); err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, err.Error()))
			return
		}

		if userToCreate.Name == "" {
			returnErrorEmptyField("Nombre")
			return
		} else if userToCreate.Age <= 0 {
			returnErrorEmptyField("Edad")
			return
		}

		userToReturn, err := controller.service.Store(userToCreate.Name, userToCreate.IsActive, userToCreate.Age)
		if err != nil {
			ctx.JSON(http.StatusConflict, web.NewResponse(409, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(200, userToReturn, ""))
	}
}

// DeleteUser
// @Summary Delete user
// @Tags    Users
// @Param   id    path   int    true "user ID"
// @Param   token header string true "token"
// @Success 204
// @Failure 400 {object} web.Response
// @Failure 401 {object} web.Response
// @Failure 404 {object} web.Response
// @Router  /users/{id} [DELETE]
func (controller *UserController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "ID inválio - "+err.Error()))
		return
	}

	if err = controller.service.Delete(int(id)); err != nil {
		ctx.JSON(http.StatusNotFound, web.NewResponse(404, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusNoContent, web.NewResponse(204, fmt.Sprintf("Le usuarie %d ha sido eliminade", id), ""))
}

// ListUsers godoc
// @Summary     Show list users
// @Tags        Users
// @Description get users
// @Accept      json
// @Produce     json
// @Param       token header   string       true "token"
// @Success     200   {object} web.Response "List users"
// @Failure     401   {object} web.Response "unauthorized"
// @Failure     404   {object} web.Response "not found users"
// @Failure     500   {object} web.Response "internal server error"
// @Router      /users [GET]
func (controller *UserController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		allUsers, err := controller.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(500, nil, err.Error()))
		}
		usersToReturn := allUsers

		_, isPresent := ctx.GetQuery("active")
		if isPresent {
			var filteredUsers []domains.User
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
			usersToReturn = filteredUsers
		}

		if len(usersToReturn) == 0 {
			ctx.JSON(http.StatusNotFound, web.NewResponse(404, nil, "no hay usuaries registrades"))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(200, usersToReturn, ""))
	}
}

// UpdateUser godoc
// @Summary Update user
// @Tags    Users
// @Accept  json
// @Produce json
// @Param   id    path     int         true  "ID user"
// @Param   token header   string      false "token"
// @Param   user  body     userRequest true  "User to update"
// @Success 200   {object} web.Response
// @Failure 400   {object} web.Response
// @Failure 401   {object} web.Response
// @Failure 404   {object} web.Response
// @Router  /users/{id} [PUT]
func (controller *UserController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "ID inválido - "+err.Error()))
			return
		}

		var userToCreate userRequest
		if err := ctx.ShouldBindJSON(&userToCreate); err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, err.Error()))
			return
		}

		if userToCreate.Name == "" {
			returnErrorEmptyField("Nombre")
			return
		} else if userToCreate.Age <= 0 {
			returnErrorEmptyField("Edad")
			return
		}

		userToUpdate, err := controller.service.Update(int(id), userToCreate.Name, userToCreate.IsActive, userToCreate.Age)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(404, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(200, userToUpdate, ""))
	}
}

// UpdateUserNameAndAge
// @Summary     Update user name and age
// @Tags        Users
// @Accept      json
// @Produce     json
// @Description This endpoint updates name and age fields from an user
// @Param       token header string           true "token header"
// @Param       id    path   int              true "user ID"
// @Param       name  body   userRequestPatch true "user name"
// @Param       age   body   userRequestPatch true "user age"
func (controller *UserController) UpdateNameAndAge(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "ID inválido - "+err.Error()))
		return
	}

	var userToCreate userRequestPatch
	if err := ctx.ShouldBindJSON(&userToCreate); err != nil {
		ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}

	if userToCreate.Name == "" {
		returnErrorEmptyField("Nombre")
		return
	} else if userToCreate.Age <= 0 {
		returnErrorEmptyField("Edad")
		return
	}

	userToUpdate, err := controller.service.UpdateNameAndAge(int(id), userToCreate.Name, userToCreate.Age)
	if err != nil {
		ctx.JSON(http.StatusNotFound, web.NewResponse(404, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, web.NewResponse(200, userToUpdate, ""))
}

func returnErrorEmptyField(fieldName string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, fmt.Sprintf("el campo %s es requerido", fieldName)))
	}
}
