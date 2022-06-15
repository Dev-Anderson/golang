package services

import (
	"dfe-admin/repository"
)

func Pedido() {
	pedidos := repository.GetPedido()

	for _, pedido := range pedidos {
		repository.CreatePedido(pedido)
	}
}
