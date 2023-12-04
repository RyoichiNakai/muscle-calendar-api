package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	r.Run(":8080")
}
