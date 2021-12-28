package main

import (
	"api/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/carro", api.GetCarro)

	router.Run("localhost:3000")
}
