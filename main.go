package main

import (
	"Spaghetti-Code/Services"
	"Spaghetti-Code/handlers"
	"Spaghetti-Code/repositories"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	alimentoHandler *handlers.AlimentoHandler
	router          *gin.Engine
)

func main() {
	// Inicializar dependencias
	Dependencies()

	// Configurar Gin
	router = gin.Default()

	// Configurar CORS
	/*router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:5500"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))*/
	router.Use(cors.Default())
	// Inicializar rutas
	//Rutas()
	router.POST("/alimentos", alimentoHandler.CrearAliment)

	// Levantar el servidor
	log.Println("Iniciando el servidor en el puerto 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error al iniciar servidor:", err)
	}
}

func Rutas() {
	// Ruta para servir el archivo HTML de listado de alimentos
	router.GET("/alimentos/listado", func(c *gin.Context) {
		c.File("./index.html")
	})

	// Grupo de rutas del API para alimentos
	group := router.Group("/alimentos")
	group.GET("/:usuario_id", alimentoHandler.ObtenerTodosAlimentosPorID)
	group.POST("", alimentoHandler.CrearAliment)
	group.PUT("/:id", alimentoHandler.ModificarAlimento)
	// group.DELETE("/:id", alimentoHandler.EliminarAlimento)

}

func Dependencies() {
	// Definición de variables de interface
	var db repositories.DB
	var alimentoRepo repositories.AlimentoRepositoryInterface
	var alimentoService Services.AlimentoInterface

	// Inicializar MongoDB y configurar las dependencias
	db = repositories.NewMongoDB()
	// Crear el repositorio de alimentos
	alimentoRepo = repositories.NewAlimentoRepository(db)
	// Crear el servicio de alimentos
	alimentoService = Services.NewAlimentoService(alimentoRepo)
	// Crear el handler de alimentos
	alimentoHandler = handlers.NewAlimentoHandler(alimentoService)
}
