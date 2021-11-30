package main 

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//declarando a estrutura
type veiculo struct {
	ID int `json:id`
	Modelo string  `json:modelo`
	Marca string `json:marca`
	Valor float64 `json:valor`
}

var carro =[]veiculo{
	{ID: 1, Modelo: "x", Marca: "coisa", Valor: 12.22},
}

func getCarro(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, carro)
}

func main() {
	router := gin.Default() 
	router.GET("/carro", getCarro)

	router.Run("localhost:3000")
}