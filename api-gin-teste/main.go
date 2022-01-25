package main

import (
	"api-gin-teste/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", api.GetAlbums)
	router.GET("/teste/:id", api.GetAlbumId)

	router.Run(":3000")
}
