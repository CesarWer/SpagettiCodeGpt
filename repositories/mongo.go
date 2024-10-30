package repositories

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client *mongo.Client
}

func NewMongoDB() *MongoDB {
	instancia := &MongoDB{}
	err := instancia.Connect()
	if err != nil {
		log.Fatalf("Error al conectar a MongoDB: %v", err)
	}

	return instancia
}

func (mongoDB *MongoDB) GetClient() *mongo.Client {
	if mongoDB.Client == nil {
		log.Fatal("La conexión a MongoDB no está inicializada")
	}
	return mongoDB.Client
}

// La dejamos privada, se ejecuta cuando se crea el objeto
func (mongoDB *MongoDB) Connect() error {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/spaghetti_codeDB")
	//clientOptions := options.Client().ApplyURI("mongodb://nuevoUser1:1234@localhost:27017/spaghetti_codeDB")
	//clientOptions := options.Client().ApplyURI("mongodb://usuario:password@localhost:27017/db?authMechanism=SCRAM-SHA-1")
	//clientOptions := options.Client().ApplyURI("mongodb://nuevoUser1:1234@localhost:27017/spaghetti_codeDB?authMechanism=SCRAM-SHA-256")

	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Println("Error al conectar a MongoDB", err)
		return err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}

	mongoDB.Client = client

	log.Println("Conexion Exitosa!")
	return nil
}

func (mongoDB *MongoDB) Disconnect() error {
	if mongoDB.Client == nil {
		return nil
	}
	return mongoDB.Client.Disconnect(context.Background())
}

/*
var client *mongo.Client

func GetClient() *mongo.Client {
	return client
}

// Conectar a la base de datos
func ConnectMongoDB() error {
	clientOptions := options.Client().ApplyURI("mongodb://nuevoUser1:1234@localhost:27017/spaghetti_codeDB")
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("Error al conectar a MongoDB:", err)
		return err
	}

	// Verificar la conexión
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("Error al hacer ping a MongoDB:", err)
		return err
	}

	log.Println("Conexión exitosa a MongoDB")
	return nil
}

// Obtener una colección
func GetCollection(collectionName string) *mongo.Collection {
	if client == nil {
		log.Fatal("La conexión a MongoDB no se ha inicializado")
	}
	return client.Database("spaghetti_codeDB").Collection(collectionName)
}
*/

var client *mongo.Client

// Obtener una colección
// Obtener una colección usando el cliente inicializado
func (mongoDB *MongoDB) GetCollection(collectionName string) *mongo.Collection {
	return mongoDB.GetClient().Database("spaghetti_codeDB").Collection(collectionName)
}
