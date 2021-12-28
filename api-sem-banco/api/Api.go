package api

import (
	_struct "api/struct"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCarro(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, _struct.Carro)
}
