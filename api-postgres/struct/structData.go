package _struct

type Filme struct {
	FilmeID   string `json:filmeid`
	FilmeNome string `json:filmenome`
}

type JsonResponse struct {
	Type    string  `json:type`
	Message string  `json:message`
	Data    []Filme `json:data`
}
