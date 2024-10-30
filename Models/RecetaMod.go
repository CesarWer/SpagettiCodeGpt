package Models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AlimentoReceta representa la relación entre un alimento y una receta
type AlimentoReceta struct {
	AlimentoID string `bson:"alimentoID"`
	Cantidad   int    `bson:"cantidad"`
}

// Receta model que representa una receta
type Receta struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	Nombre             string             `bson:"nombre"`
	Momento            string             `bson:"momento"`
	Alimentos          []AlimentoReceta   `bson:"alimentos"`
	FechaCreacion      time.Time          `bson:"fecha_creacion"`
	FechaActualizacion time.Time          `bson:"fecha_actualizacion"`
	UsuarioID          string             `bson:"usuario_id"`
	Eliminada          bool               `bson:"eliminada"`
}

/*Métodos adicionales para lógica de negocio (ejemplo)
func (r *Receta) EsDisponible(alimentos []Alimento) bool {
	for _, alimentoReceta := range r.Alimentos {
		disponible := false
		for _, alimento := range alimentos {
			if alimento.ID == alimentoReceta.AlimentoID && alimento.CantidadActual >= int(alimentoReceta.Cantidad) {
				disponible = true
				break
			}
		}
		if !disponible {
			return false
		}
	}
	return true
}*/
