package _struct

import (
	"time"
)

type StructData struct {
	Gid             int64     `json:"gid"`
	Numero          int64     `json:"numero"`
	Serie           int64     `json:"serie"`
	Id              string    `json:"id"`
	NumeroProtocolo string    `json:numeroprotocolo`
	DigValue        string    `json:digvalue`
	DataEmissao     time.Time `json:dataemissao`
	HoraEmissao     time.Time `json:horaemissao`
}

type StructDataCanceladas struct {
	Gid         int64     `json:"gid"`
	Numero      int64     `json:"Numero"`
	Serie       int64     `json:"Serie"`
	Id          string    `json:"id"`
	DataEmissao time.Time `json:"dataemissao"`
	HoraEmissao time.Time `json:"horaemissao"`
	CanceladoEm time.Time `json:"canceladoem"`
}
