package Models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Compra struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	UsuarioID          string             `bson:"usuario_id"`
	FechaCompra        time.Time          `bson:"fecha_compra"`
	AlimentosComprados []AlimentoCompra   `bson:"alimentos_comprados"`
	Total              float64            `bson:"total"`
	FechaActualizacion time.Time          `bson:"fecha_actualizacion"`
}

type AlimentoCompra struct {
	AlimentoID     string  `bson:"alimento_id"`
	Cantidad       float64 `bson:"cantidad"`
	PrecioUnitario float64 `bson:"precio_unitario"`
}
