package handler

import (
	"log"
	"net/http"
	"os"

	"github.com/dionysseidel/backpack-bcgow6-dionys-seidel/pkg/web"
	"github.com/gin-gonic/gin"
)

func validateToken(ctx *gin.Context) {
	log.Println("Este es le primer middleware")
	token := ctx.GetHeader("token") /*ctx.Request.Header.Get("token")*/
	// ... Me imagino que acá querría persistir le token en el archivo .env
	if token != os.Getenv("TOKEN") || token == "" {
		// ctx.JSON(http.StatusUnauthorized, web.NewResponse(401, nil, "no tiene permisos para realizar la petición solicitada"))
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, web.NewResponse(401, nil, "no tiene permisos para realizar la petición solicitada"))
		// authenticated = false
		return
	}
	// return true
	ctx.Next()
}

func MiddlewareList(f gin.HandlerFunc) []gin.HandlerFunc {
	list := []gin.HandlerFunc{
		validateToken,
	}
	list = append(list, f)
	return list
}
