package _struct

type Filme struct {
	ID        int    `json:id`
	FilmeID   string `json:filmeid`
	FilmeNome string `json:filmenome`
}

type JsonResponse struct {
	Type    string  `json:type`
	Data    []Filme `json:data`
	Message string  `json:message`
}
