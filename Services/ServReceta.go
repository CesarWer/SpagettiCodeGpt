package Services

/*
import (

	"Spaghetti-Code/Models"

	"Spaghetti-Code/connections"

	"Spaghetti-Code/Models"
	"Spaghetti-Code/dto"
	//"Spaghetti-Code/repositories"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var recetaCollection *mongo.Collection = connections.GetCollection("recetas")
var alimentoCollection *mongo.Collection = connections.GetCollection("alimentos")

var recetasCollection *mongo.Collection
	//"go.mongodb.org/mongo-driver/mongo"
)

//var recetaCollection *mongo.Collection = repositories.GetCollection("recetas")
//var alimentoCollection *mongo.Collection = repositories.GetCollection("alimentos")

// Crea una receta y actualiza el stock de los alimentos
func CrearReceta(recetaDTO dto.RecetaDto) error {
	var alimentosReceta []Models.AlimentoReceta
	ctx := context.TODO()

	// Validar y descontar cantidades de alimentos
	for _, alimentoReceta := range recetaDTO.Alimentos {
		var alimento Models.Alimento

		// Buscar el alimento en la base de datos
		err := alimentoCollection.FindOne(ctx, bson.M{"_id": alimentoReceta.AlimentoID}).Decode(&alimento)
		if err != nil {
			return errors.New("alimento no encontrado")
		}

		// Validar si el alimento coincide con el uso (momento)
		if alimento.Momentos[0] != recetaDTO.Momento {
			return errors.New("el alimento " + alimento.Nombre + " no coincide con el uso de la receta")
		}

		// Validar si hay suficiente stock disponible
		if alimento.CantidadActual < alimentoReceta.Cantidad {
			return errors.New("no hay suficiente stock del alimento " + alimento.Nombre)
		}

		// Restar la cantidad del alimento
		alimento.CantidadActual -= alimentoReceta.Cantidad
		_, err = alimentoCollection.UpdateOne(
			ctx,
			bson.M{"_id": alimento.ID},
			bson.M{"$set": bson.M{"cantidad_actual": alimento.CantidadActual}},
		)
		if err != nil {
			return err
		}

		// Mapeo a modelo de receta
		alimentosReceta = append(alimentosReceta, Models.AlimentoReceta{
			AlimentoID: alimento.ID.String(),
			Cantidad:   alimentoReceta.Cantidad,
		})
	}

	// Crear la receta
	receta := Models.Receta{
		Nombre:             recetaDTO.Nombre,
		Momento:            recetaDTO.Momento,
		Alimentos:          alimentosReceta,
		FechaCreacion:      time.Now(),
		FechaActualizacion: time.Now(),
		UsuarioID:          recetaDTO.UsuarioID,
	}

	_, err := recetaCollection.InsertOne(ctx, receta)
	if err != nil {
		return err
	}

	return nil
}

// EliminarReceta: Elimina una receta
func EliminarReceta(recetaID string) error {
	var receta Models.Receta
	ctx := context.TODO()

	// Buscar la receta en la base de datos
	err := recetaCollection.FindOne(ctx, bson.M{"_id": recetaID}).Decode(&receta)
	if err != nil {
		return errors.New("receta no encontrada")
	}

	// Devolver la cantidad a los alimentos
	for _, alimentoReceta := range receta.Alimentos {
		_, err := alimentoCollection.UpdateOne(
			ctx,
			bson.M{"_id": alimentoReceta.AlimentoID},
			bson.M{"$inc": bson.M{"cantidad_actual": alimentoReceta.Cantidad}},
		)
		if err != nil {
			return err
		}
	}

	// Eliminar la receta de la base de datos
	_, err = recetaCollection.DeleteOne(ctx, bson.M{"_id": recetaID})
	if err != nil {
		return err
	}

	return nil
}

// BuscarRecetasDisponibles: Busca recetas disponibles
func BuscarRecetasDisponibles(filtros dto.FiltrosRecetaDTO, usuarioID string) ([]Models.Receta, error) {
	var recetas []Models.Receta
	ctx := context.TODO()

	// Crear filtro base con los nuevos criterios
	filtro := bson.M{
		"usuario_id": usuarioID, // Filtrar por el usuario
		"eliminado":  false,     // Solo las recetas que no están eliminadas
	}

	if filtros.TipoProducto != "" {
		filtro["alimentos.tipo"] = bson.M{"$regex": filtros.TipoProducto, "$options": "i"}
	}

	if filtros.NombreProducto != "" {
		filtro["alimentos.nombre"] = bson.M{"$regex": filtros.NombreProducto, "$options": "i"}
	}

	cursor, err := recetaCollection.Find(ctx, filtro)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Validar que los productos tienen stock suficiente
	for cursor.Next(ctx) {
		var receta Models.Receta
		if err := cursor.Decode(&receta); err != nil {
			return nil, err
		}

		disponible := true
		for _, alimentoReceta := range receta.Alimentos {
			var alimento Models.Alimento
			err := alimentoCollection.FindOne(ctx, bson.M{"_id": alimentoReceta.AlimentoID}).Decode(&alimento)
			if err != nil {
				return nil, err
			}

			if alimento.CantidadActual < alimentoReceta.Cantidad {
				disponible = false
				break
			}
		}

		if disponible {
			recetas = append(recetas, receta)
		}
	}

	return recetas, nil
}
*/
