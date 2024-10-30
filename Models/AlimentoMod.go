package Models

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Alimento model que representa un alimento en la aplicación
type Alimento struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	Tipo               string             `bson:"tipo" json:"tipo"`
	Nombre             string             `bson:"nombre" json:"nombre"`
	Momentos           []string           `bson:"momentos" json:"momentos"`
	PrecioUnitario     float64            `bson:"precio_unitario" json:"precioUnitario"`
	CantidadActual     int                `bson:"cantidad_actual" json:"cantidad"`
	CantidadMinima     int                `bson:"cantidad_minima" json:"cantidadMinima"`
	FechaCreacion      time.Time          `bson:"fecha_creacion"`
	FechaActualizacion time.Time          `bson:"fecha_actualizacion"`
	Eliminado          bool               `bson:"eliminado"`
	UsuarioID          string             `bson:"usuario_id"`
}

// Métodos adicionales para manejar lógica de negocio (ejemplo)
func (a *Alimento) RestarCantidad(cantidad int) error {
	if a.CantidadActual < cantidad {
		return errors.New("cantidad insuficiente")
	}
	a.CantidadActual -= cantidad
	a.FechaActualizacion = time.Now().Truncate(24 * time.Hour)
	return nil
}

func (a *Alimento) SumarCantidad(cantidad int) {
	a.CantidadActual += cantidad
	a.FechaActualizacion = time.Now().Truncate(24 * time.Hour)
}
