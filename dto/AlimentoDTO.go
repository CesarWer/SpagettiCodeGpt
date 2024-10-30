package dto

import (
	"Spaghetti-Code/Models"
	"Spaghetti-Code/utils"
)

// ProductoDTO define la estructura para la información de un producto
type AlimentoDto struct {
	ID             string   `json:"id,omitempty"`
	Tipo           string   `json:"tipo"`
	Nombre         string   `json:"nombre"`
	Momentos       []string `json:"momentos"`
	PrecioUnitario float64  `json:"precio_unitario"`
	CantidadActual int      `json:"cantidad_actual"`
	CantidadMinima int      `json:"cantidad_minima"`
	UsuarioID      string   `json:"usuario_id"`
}

func NewAlimento(alimento Models.Alimento) *AlimentoDto {
	return &AlimentoDto{
		ID:     utils.GetStringIDFromObjectID(alimento.ID),
		Nombre: alimento.Nombre,
	}
}

func (alimento AlimentoDto) GetModel() Models.Alimento {
	return Models.Alimento{
		ID:     utils.GetObjectIDFromStringID(alimento.ID),
		Nombre: alimento.Nombre,
	}
}
