package routes

import (
	"database/sql"

	"github.com/bootcamp-go/go-bDs/cmd/server/handler"
	"github.com/bootcamp-go/go-bDs/internal/products"
	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	db       *sql.DB
	endpoint *gin.RouterGroup
	engine   *gin.Engine
}

func NewRouter(engine *gin.Engine, db *sql.DB) Router {
	return &router{
		db:     db,
		engine: engine,
	}
}

func (r *router) MapRoutes() {
	r.setGroup()
	r.buildProductRoutes()
}

func (r *router) setGroup() {
	r.endpoint = r.engine.Group("/api/v1")
}

func (r *router) buildProductRoutes() {
	repository := products.NewRepository(r.db)
	service := products.NewService(repository)
	handler := handler.NewProduct(service)

	r.endpoint.GET("/products/:name", handler.GetByName())
	r.endpoint.GET("/products", handler.GetAll())
	r.endpoint.DELETE("/products/:id", handler.Delete())
}
