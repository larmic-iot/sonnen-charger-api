package routers

import (
	"github.com/gin-gonic/gin"
	"larmic/sonnen-charger-api/internal/routers/handlers"
)

func InitRouter(chargerIp string) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/", handlers.GetOpenApi)

	api := r.Group("/sonnen-charger-api")
	api.GET("/settings", handlers.GetSettings(chargerIp))
	api.GET("/", handlers.GetOpenApi)

	return r
}
