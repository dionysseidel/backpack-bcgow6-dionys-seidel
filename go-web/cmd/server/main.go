package main

import (
	"log"
	"os"

	"github.com/dionysseidel/backpack-bcgow6-dionys-seidel/cmd/server/handler"
	"github.com/dionysseidel/backpack-bcgow6-dionys-seidel/internal/users"
	"github.com/dionysseidel/backpack-bcgow6-dionys-seidel/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
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

	router := gin.Default()

	api := router.Group("/api/v1")

	// Documentación Swagger
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// router.GET("/hello-name", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"message": "Hola " + usersSlice[0].Name})
	// })
	usersEndoint := router.Group("/users")
	{
		usersEndoint.GET("", handler.MiddlewareList(controller.GetAll())...)
		usersEndoint.POST("", handler.MiddlewareList(controller.CreateUser())...)
		usersEndoint.DELETE("/:id", handler.MiddlewareList(controller.Delete)...)
		// usersEndoint.GET("/:id", handler.MiddlewareList(GetOne)...)
		usersEndoint.PATCH("/:id", handler.MiddlewareList(controller.UpdateNameAndAge)...)
		usersEndoint.PUT("/:id", handler.MiddlewareList(controller.Update())...)
	}

	if err := router.Run(); err != nil {
		panic(err)
	}
}

func loadEnd() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("No se pudo cargar las variables de entorno - error: ", err)
	}
}
