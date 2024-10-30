package repositories

import (
	"Spaghetti-Code/Models"
	//"Spaghetti-Code/utils"
	"context"
	//"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// AlimentoRepositoryInterface define las operaciones disponibles para AlimentoRepository
type AlimentoRepositoryInterface interface {
	CrearAliment(alimento Models.Alimento) (*mongo.InsertOneResult, error)
	ModificarAlimento(alimento Models.Alimento) (*mongo.UpdateResult, error)
	ObtenerTodosAlimentosPorID(usuarioID string) ([]Models.Alimento, error)
}

// AlimentoRepository implementa AlimentoRepositoryInterface
type AlimentoRepository struct {
	db DB
}

func NewAlimentoRepository(db DB) *AlimentoRepository {
	return &AlimentoRepository{
		db: db,
	}
}

// FUNCIONA
func (repository *AlimentoRepository) CrearAliment(alimento Models.Alimento) (*mongo.InsertOneResult, error) {
	// Obtener la colección de la base de datos
	collection := repository.db.GetClient().Database("spaghetti_codeDB").Collection("alimentos")
	// Insertar el alimento en la colección
	resultado, err := collection.InsertOne(context.Background(), alimento)
	if err != nil {
		// Podrías loguear el error aquí si es necesario
		return nil, err
	}
	// Devolver el resultado de la inserción y el error (si lo hubo)
	return resultado, nil
}

// SIN PROBAR
func (repository AlimentoRepository) ModificarAlimento(alimento Models.Alimento) (*mongo.UpdateResult, error) {
	collection := repository.db.GetClient().Database("spaghetti_codeDB").Collection("alimentos")

	filtro := bson.M{"_id": alimento.ID}
	entidad := bson.M{
		"$set": bson.M{
			"nombre":             alimento.Nombre,
			"tipo":               alimento.Tipo,
			"precioUnitario":     alimento.PrecioUnitario,
			"cantidadActual":     alimento.CantidadActual,
			"cantidadMinima":     alimento.CantidadMinima,
			"momentos":           alimento.Momentos,
			"fechaActualizacion": time.Now().Truncate(24 * time.Hour),
		},
	}

	resultado, err := collection.UpdateOne(context.TODO(), filtro, entidad)

	//ni idea si este if hace falta o no
	if err != nil {
		return nil, err
	}
	return resultado, err
}

/*
	func (repository AlimentoRepository) ObtenerTodosAlimentosPorID(usuarioID string) ([]Models.Alimento, error) {
		var alimento []Models.Alimento
		collection := repository.db.GetClient().Database("spaghetti_codeDB").Collection("alimentos")

		objectID := utils.GetObjectIDFromStringID(usuarioID)
		filtro := bson.M{"usuario_id": objectID}

		// Consultar documentos que coincidan con el UsuarioID
		cursor, err := collection.Find(context.TODO(), filtro)
		defer cursor.Close(context.Background())

		// Iterar y almacenar los documentos en el slice
		for cursor.Next(context.Background()) {
			err := cursor.Decode(&alimento)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		}
		return alimento, err
	}
*/
func (r *AlimentoRepository) ObtenerTodosAlimentosPorID(usuarioID string) ([]Models.Alimento, error) {
	var alimentos []Models.Alimento
	collection := r.db.GetCollection("alimentos")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"usuario_id": usuarioID}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Println("Error al buscar alimentos:", err)
		return nil, err
	}

	if err = cursor.All(ctx, &alimentos); err != nil {
		log.Println("Error al decodificar resultados:", err)
		return nil, err
	}

	return alimentos, nil
}

// FALTA EDITAR
func (repository AlimentoRepository) EliminarAula(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	collection := repository.db.GetClient().Database("spaghetti_codeDB").Collection("alimentos")

	filtro := bson.M{"_id": id}

	resultado, err := collection.DeleteOne(context.TODO(), filtro)

	return resultado, err
}
