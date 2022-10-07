package main

import (
	"log"

	"github.com/dionysseidel/backpack-bcgow6-dionys-seidel/cmd/server/handler"
	"github.com/dionysseidel/backpack-bcgow6-dionys-seidel/internal/users"
	"github.com/dionysseidel/backpack-bcgow6-dionys-seidel/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

// @title          Bootcamp Go Wave 6 - API
// @version        1.0
// @description    This API Handle MELI Users. This is a simple API development conducted by Digital House's team.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name  API Support Dionys Seidel
// @contact.url   http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @host     localhost:8080
// @BasePath /api/v1
func main() {
	loadEnd()
	db := store.New(store.FileType, "./users.json")
	repo := users.NewRepository(db)
	service := users.NewService(repo)
	controller := handler.NewUserController(service)

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
		usersEndoint.DELETE("/:id", controller.Delete)
		// usersEndoint.GET("/:id", GetOne)
		usersEndoint.PATCH("/:id", controller.UpdateNameAndAge)
		usersEndoint.PUT("/:id", controller.Update())
	}
	server.Run()
}

func loadEnd() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("No se pudo cargar las variables de entorno - error: ", err)
	}
}
