package _struct

import (
	"time"
)

type StructData struct {
	Numero          int64     `json:"numero"`
	Serie           int64     `json:"serie"`
	DataEmissao     time.Time `json:"dataemissao"`
	HoraEmissao     time.Time `json:"horaemissao"`
	Id              string    `json:"id"`
	Protocolo       string    `json:"protocolo"`
	ValorNotafiscal float64   `json:"valornotafiscal"`
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

type StructDataNfe struct {
	Gid         int64     `json:"gid"`
	Numero      int64     `json:"numero"`
	Serie       int64     `json:"serie"`
	DataEmissao time.Time `json:"dataemissao"`
	HoraEmissao time.Time `json:"horaemissao"`
	Id          string    `json:"id"`
	Protocolo   string    `json:"protocolo"`
}
