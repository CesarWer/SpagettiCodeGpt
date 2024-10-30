package Services

import (
	"Spaghetti-Code/Models"
	"Spaghetti-Code/dto"
	"Spaghetti-Code/repositories"

	//"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive" // Para manejar el ID de MongoDB
)

type AlimentoInterface interface {
	CrearAliment(alimento dto.AlimentoDto) error
	ModificarAlimentos(alimentoDTO dto.AlimentoDto) error
	ObtenerTodosAlimentosPorID(usuarioID string) ([]Models.Alimento, error)
}

type AlimentoService struct {
	repository repositories.AlimentoRepositoryInterface
}

// Crear una nueva instancia del servicio de Alimentos utilizando la conexión a MongoDB
func NewAlimentoService(repository repositories.AlimentoRepositoryInterface) AlimentoInterface {
	return &AlimentoService{repository: repository}
}

// CrearAliment inserta un nuevo alimento en la base de datos MongoDB
func (service *AlimentoService) CrearAliment(alimento dto.AlimentoDto) error {

	// Crear el modelo a partir del DTO
	alimentoModel := Models.Alimento{
		ID:                 primitive.NewObjectID(), // Generar un nuevo ID para MongoDB
		Tipo:               alimento.Tipo,
		Nombre:             alimento.Nombre,
		Momentos:           alimento.Momentos,
		PrecioUnitario:     alimento.PrecioUnitario,
		CantidadActual:     alimento.CantidadActual,
		CantidadMinima:     alimento.CantidadMinima,
		FechaCreacion:      time.Now().Truncate(24 * time.Hour),
		FechaActualizacion: time.Now().Truncate(24 * time.Hour),
		UsuarioID:          alimento.UsuarioID,
	}

	// Validación: Precios y cantidades deben ser positivas
	if alimentoModel.PrecioUnitario < 0 || alimentoModel.CantidadActual < 0 || alimentoModel.CantidadMinima < 0 {
		return errors.New("cantidad o precio no pueden ser negativos")
	}

	// Insertar el alimento en la base de datos usando el repositorio
	result, err := service.repository.CrearAliment(alimentoModel) // Ahora manejamos el InsertOneResult
	if err != nil {
		return err
	}

	// Registrar el ID del nuevo alimento insertado
	log.Println("Alimento insertado con ID:", result.InsertedID)

	return nil
}

// LISTO (SIN PROBAR)
func (service *AlimentoService) ModificarAlimentos(alimentoDTO dto.AlimentoDto) error {
	// Validar datos: los precios y cantidades no deben ser negativos
	if alimentoDTO.PrecioUnitario < 0 || alimentoDTO.CantidadActual < 0 || alimentoDTO.CantidadMinima < 0 {
		return errors.New("los valores de precio y cantidad no pueden ser negativos")
	}

	// Convertir el ID a ObjectID para MongoDB
	objID, err := primitive.ObjectIDFromHex(alimentoDTO.ID)
	if err != nil {
		return errors.New("ID no válido para MongoDB")
	}

	// Convertir el DTO al modelo de la base de datos
	alimento := Models.Alimento{
		ID:                 objID,
		Tipo:               alimentoDTO.Tipo,
		Nombre:             alimentoDTO.Nombre,
		Momentos:           alimentoDTO.Momentos,
		PrecioUnitario:     alimentoDTO.PrecioUnitario,
		CantidadActual:     alimentoDTO.CantidadActual,
		CantidadMinima:     alimentoDTO.CantidadMinima,
		FechaCreacion:      alimentoDTO.GetModel().FechaCreacion,
		FechaActualizacion: time.Now().Truncate(24 * time.Hour),
		UsuarioID:          alimentoDTO.UsuarioID,
	}

	// Llamar al repositorio para actualizar el alimento
	_, err = service.repository.ModificarAlimento(alimento)
	if err != nil {
		return err
	}

	return nil
}

// FUNCIONA
// func (service *AlimentoService) ObtenerTodosAlimentos() ([]Models.Alimento, error) {
// 	alimentos, err := service.repository.ObtenerTodosAlimentos()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return alimentos, nil
// }

func (service *AlimentoService) ObtenerTodosAlimentosPorID(usuarioID string) ([]Models.Alimento, error) {
	return service.repository.ObtenerTodosAlimentosPorID(usuarioID)
}

//Un desastre
/*
import (
	"Spaghetti-Code/Models"
	//"Spaghetti-Code/connections"
	"Spaghetti-Code/repositories"
	"Spaghetti-Code/dto"
	"context"
	"errors"

	//"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/mongo"
)

type AlimentoInterface interface {
	/*ObtenerAulas(filtroNombre string) []*dto.AlimentoDto
	ObtenerAulaPorID(id string) *dto.AlimentoDto
	EliminarAula(id string) bool
	InsertarAlimento(alimento *dto.AlimentoDto) error
	ModificarAula(aula *dto.AlimentoDto) bool/
	InsertarAlimento(aula *dto.AlimentoDto) bool
	ModificarAula(aula *dto.AlimentoDto) bool
	InsertarAlimento(alimento *dto.AlimentoDto) error
	ModificarAula(aula *dto.AlimentoDto) bool
	CrearAliment(alimento dto.AlimentoDto) error
}

type AlimentoService struct {
	collection repositories.AlimentoRepositoryInterface
}

// Crear una nueva instancia del servicio de Alimentos utilizando la conexión a MongoDB
func NewAlimentoService(collection repositories.AlimentoRepositoryInterface) *AlimentoService {
	return &AlimentoService{collection: collection}
}

func (service *AlimentoService) CrearAliment(alimento dto.AlimentoDto) error {
	service.collection.CrearAliment(alimento.GetModel())

	return nil
}

// Crear una nueva instancia del servicio de Alimentos utilizando la conexión a MongoDB
func NewAlimentoService(collection repositories.AlimentoRepositoryInterface) *AlimentoService {
	return &AlimentoService{collection: collection}
}

func (service *AlimentoService) CrearAliment(alimento dto.AlimentoDto) error {
	service.collection.CrearAliment(alimento.GetModel())

	return nil
}

var alimentosCollection = repositories.GetCollection("alimentos")

// CrearAlimento inserta un nuevo alimento en la base de datos

func (s *AlimentoService) CrearAlimento(alimentoDTO dto.AlimentoDto) error {


func (s *AlimentoService) CrearAliment(alimentoDTO dto.AlimentoDto) error {

	alimento := Models.Alimento{
		Tipo:               alimentoDTO.Tipo,
		Nombre:             alimentoDTO.Nombre,
		Momentos:           alimentoDTO.Momentos,
		PrecioUnitario:     alimentoDTO.PrecioUnitario,
		CantidadActual:     alimentoDTO.CantidadActual,
		CantidadMinima:     alimentoDTO.CantidadMinima,
		FechaCreacion:      time.Now(),
		FechaActualizacion: time.Now(),
		UsuarioID:          alimentoDTO.UsuarioID,
	}

	// Validar cantidades y precios aquí si es necesario
	if alimento.PrecioUnitario < 0 || alimento.CantidadActual < 0 || alimento.CantidadMinima < 0 {
		return errors.New("Cantidad o precio no pueden ser negativos")
	}

	// Insertar en MongoDB

	_, err := alimentosCollection.InsertOne(context.TODO(), alimento)
	if err != nil {
		return err
	}

	return nil
}

// Restar stock del alimento
func RestarStockAlimento(id string, cantidad int) error {

	// Buscar el alimento por ID
	var alimento Models.Alimento
	err := alimentosCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&alimento)
	if err != nil {
		return errors.New("Alimento no encontrado")
	}

	// Restar la cantidad utilizando el método del model
	if err := alimento.RestarCantidad(cantidad); err != nil {
		return err
	}

	// Actualizar la base de datos
	_, err = alimentosCollection.ReplaceOne(context.TODO(), bson.M{"_id": alimento.ID}, alimento)
	if err != nil {
		return err
	}

	return nil
}

// Listar todos los alimentos (sin filtros)
func ListarAlimentos() ([]Models.Alimento, error) {
	var alimentos []Models.Alimento

	cursor, err := alimentosCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var alimento Models.Alimento
		if err = cursor.Decode(&alimento); err != nil {
			return nil, err
		}
		alimentos = append(alimentos, alimento)
	}

	return alimentos, nil
}

func ModificarAlimentos(id string, alimentoDTO dto.AlimentoDto) error {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"tipo":               alimentoDTO.Tipo,
			"nombre":             alimentoDTO.Nombre,
			"momentos":           alimentoDTO.Momentos,
			"precioUnitario":     alimentoDTO.PrecioUnitario,
			"cantidadActual":     alimentoDTO.CantidadActual,
			"cantidadMinima":     alimentoDTO.CantidadMinima,
			"fechaActualizacion": time.Now(),
		},
	}
	_, err := alimentosCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func EliminarAlimentoLogico(id string) error {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"eliminado": true,
		},
	}
	_, err := alimentosCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}
func ListarAlimentosPorUsuari(usuarioID string) ([]Models.Alimento, error) {
	var alimentos []Models.Alimento
	filter := bson.M{
		"usuarioID": usuarioID,
		"eliminado": false, // Solo traemos los que no están eliminados
	}

	cursor, err := alimentosCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &alimentos); err != nil {
		return nil, err
	}

	return alimentos, nil
}
*/
