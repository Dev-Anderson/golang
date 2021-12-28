# API sem banco de dados
**Pacotes utilizados:**

> - net/http
> - github.com/gin-gonic/gin

1- Criando a estrutura da API, dentro da pasta 'struct' foi criado o arquivo 'struct.go' com a seguinte estrutura: 

    type Veiculo struct {
	    id int `json:id`
	    modelo string `json:modelo`
	    marca string `json:marca`
	    ano int `json:ano`
	    valor float64 `json:valor`
    }
Logo de criar a estrutura foi criado uma variável com o nome 'carro' que vai receber os dados da estrutura criada. 

2- Dentro da pasta 'api' foi adicionado o arquivo 'Api.go' que vai conter a consulta que será feita, dentro desse arquivo foi adicionado uma referência para o arquivo 'struct.go', logo depois criado a rota da seguinte forma: 

    func getCarro(c *gin.Context) {
	    c.IndentedJSON(http.SttusOK, _struct.Carro)
    }
Veja que neste código foi repassado a variável 'carro' que está dentro da estrutura criada anteriormente . 
3- Dentro do arquivo 'main.go' na função 'main', foi colocado o seguinte código: 

    router := gin.Default()
    router.GET("/carro", api.GetCarro)
    router.Run("localhost:3000")
Criado uma rota do tipo 'GET' logo depois repassado para criar um servidor em 'localhost:3000' e a rota criada foi '/carro'. 
