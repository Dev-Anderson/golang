package api

import (
	"api-gin-teste/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	fmt.Println("Consultando Album")
	c.IndentedJSON(http.StatusOK, models.Albums)
}

func GetAlbumId(c *gin.Context) {
	id := c.Param("id")

	for _, a := range models.Albums {
		if a.Id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}

func PostAlbum(c *gin.Context) {
	var newAlbum = models.Albums

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	models.Albums = append(models.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
