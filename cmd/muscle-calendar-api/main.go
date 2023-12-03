package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @BasePath /v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} OK
// @Router /health [get]
func Health(g *gin.Context) {
	g.JSON(http.StatusOK, "OK")
}

// todo: いい感じにhandlerに落とし込んでいく
func main() {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/health", Health)

	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")
}
