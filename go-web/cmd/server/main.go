package main

import (
	"github.com/dionysseidel/backpack-bcgow6-dionys-seidel/go-web/cmd/server/handler"
	"github.com/dionysseidel/backpack-bcgow6-dionys-seidel/go-web/internal/users"
	"github.com/gin-gonic/gin"
)

// func GetOne(ctx *gin.Context) {
// 	id, err := strconv.Atoi(ctx.Param("id"))
// 	if err != nil {
// 		ctx.String(404, "employee information doesn't exist! \n")
// 		log.Fatal("cannot parse ID to int value", err)
// 	}
// 	user := usersSlice[id]
// 	ctx.String(200, "Employee information %s, name: %s \n", ctx.Param("id"), user)
// }

func main() {
	repo := users.NewRepository()
	service := users.NewService(repo)
	controller := handler.NewUser(service)

	// data, err := ioutil.ReadFile("./go-web/users.json")
	// if err != nil {
	// 	log.Fatal("file is not being read", err)
	// }
	// err = json.Unmarshal(data, &usersSlice)
	// if err != nil {
	// 	log.Fatal("JSON cannot be decoded", err)
	// }
	// fmt.Println("usersSlice en main", usersSlice)
	server := gin.Default()
	// server.GET("/hello-name", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"message": "Hola " + usersSlice[0].Name})
	// })
	usersEndoint := server.Group("/users")
	{
		usersEndoint.GET("", controller.GetAll())
		usersEndoint.POST("", controller.CreateUser())
		// usersEndoint.GET("/users/:id", GetOne)
	}
	server.Run()
}
