package dto

type RecetaDto struct {
	Nombre    string              `json:"nombre"`
	Momento   string              `json:"momento"`
	Alimentos []AlimentoRecetaDto `json:"alimentos"`
	UsuarioID string              `json:"usuarioId"`
}

type AlimentoRecetaDto struct {
	AlimentoID string `json:"alimentoId"`
	Cantidad   int    `json:"cantidad"`
}
type FiltrosRecetaDTO struct {
	TipoProducto   string `json:"tipo_producto,omitempty"`
	NombreProducto string `json:"nombre_producto,omitempty"`
}
