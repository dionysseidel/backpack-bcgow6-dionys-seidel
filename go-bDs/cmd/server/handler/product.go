package handler

import (
	"net/http"
	"strconv"

	"github.com/bootcamp-go/go-bDs/internal/products"
	"github.com/bootcamp-go/go-bDs/pkg/web"
	"github.com/gin-gonic/gin"
)

type Product struct {
	service products.Service
}

func NewProduct(service products.Service) *Product {
	return &Product{
		service: service,
	}
}

func (p *Product) GetByName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		productNameToSearch := ctx.Param("name")
		if productNameToSearch == "" {
			ctx.JSON(http.StatusUnprocessableEntity, web.NewResponse(nil, "product should be queried by appropriate name format", 422))
			return
		}

		product, err := p.service.GetByName(ctx, productNameToSearch)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(nil, err.Error(), 404))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(product, "", 200))
	}
}

func (p *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := p.service.GetAll(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(nil, err.Error(), 200))
		}

		ctx.JSON(http.StatusOK, web.NewResponse(products, "", 200))
	}
}

func (p *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt((ctx.Param("id")), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, web.NewResponse(nil, "product should be queried by appropriate ID format", 422))
			return
		}

		err = p.service.Delete(ctx, int(id))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(nil, err.Error(), 500))
			return
		}

		ctx.JSON(http.StatusNoContent, web.NewResponse(gin.H{"delete": id}, "", 204))
	}
}
