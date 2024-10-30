package dto

import "time"

type CompraDTO struct {
	Fecha      time.Time           `json:"fecha"`
	Productos  []ProductoCompraDTO `json:"productos"`
	CostoTotal float64             `json:"costo_total"`
}

type ProductoCompraDTO struct {
	AlimentoID       string  `json:"alimento_id"`
	Nombre           string  `json:"nombre"`
	CantidadComprada int     `json:"cantidad_comprada"`
	PrecioUnitario   float64 `json:"precio_unitario"`
	CostoTotal       float64 `json:"costo_total"`
}
