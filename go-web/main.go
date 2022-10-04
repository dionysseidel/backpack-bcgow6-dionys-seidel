package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Users struct {
	ID     int    `json:id`
	Name   string `json:"nombre"`
	Active bool   `json:"active"`
	Age    int    `json:"edad"`
}

type User []Users

var usersSlice User

var stringifiedFalseBool = strconv.FormatBool(false)
var stringifiedTrueBool = strconv.FormatBool(true)

func GetAll(ctx *gin.Context) {
	if ctx.Query("active") == stringifiedTrueBool || ctx.Query("active") == stringifiedFalseBool {
		queryUsersByStatus(ctx)
	} else {
		ctx.JSON(http.StatusOK, &usersSlice)
	}
}

func queryUsersByStatus(ctxt *gin.Context) []Users {
	users := usersSlice
	// fmt.Print("usersSlice en GetAll", users)
	var filtered []Users
	for _, user := range users {
		// fmt.Println(strconv.FormatBool(user.Active), strconv.FormatBool(true))
		switch ctxt.Query("active") == strconv.FormatBool(true) {
		case user.Active:
			filtered = append(filtered, user)
		default:
			switch ctxt.Query("active") == strconv.FormatBool(false) {
			case !user.Active:
				filtered = append(filtered, user)
			}
		}
	}
	// fmt.Println("filtered in filterUsers", filtered)
	ctxt.JSON(http.StatusOK, filtered)
	return filtered
}

func GetOne(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.String(404, "employee information doesn't exist! \n")
		log.Fatal("cannot parse ID to int value", err)
	}
	user := usersSlice[id]
	ctx.String(200, "Employee information %s, name: %s \n", ctx.Param("id"), user)
	// if err {
	// 	ctx.String(200, "Employee information %s, name: %s \n", ctx.Param("id"), user)
	// } else {
	// 	ctx.String(404, "employee information doesn't exist! \n")
	// }
}

func main() {
	data, err := ioutil.ReadFile("./go-web/users.json")
	if err != nil {
		log.Fatal("file is not being read", err)
	}
	err = json.Unmarshal(data, &usersSlice)
	if err != nil {
		log.Fatal("JSON cannot be decoded", err)
	}
	// fmt.Println("usersSlice en main", usersSlice)
	server := gin.Default()
	server.GET("/hello-name", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hola " + usersSlice[0].Name})
	})
	usersEndoint := server.Group("/users")
	{
		usersEndoint.GET("/", GetAll)
		usersEndoint.GET("/users/:id", GetOne)
	}
	server.Run()
}
