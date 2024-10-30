package main

//PROBAR INSERTAR ALIMENTO

// // Iniciar conexión con MongoDB
// mongoDB := repositories.NewMongoDB()
// defer mongoDB.Disconnect()

// // Crear el repositorio de alimentos
// alimentoRepository := repositories.NewAlimentoRepository(mongoDB)

// // Crear el servicio de alimentos
// alimentoService := Services.NewAlimentoService(alimentoRepository)

// // Crear un router con GIN
// r := gin.Default()

// // Ruta para crear un alimento
// r.POST("/alimentos", func(c *gin.Context) {
// 	// Crear un DTO vacío para recibir el JSON
// 	var alimentoDto dto.AlimentoDto

// 	// Parsear el cuerpo de la solicitud (JSON) al DTO
// 	if err := c.ShouldBindJSON(&alimentoDto); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
// 		return
// 	}

// 	// Llamar al servicio para crear el alimento
// 	err := alimentoService.CrearAliment(alimentoDto)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el alimento: " + err.Error()})
// 		return
// 	}

// 	// Devolver una respuesta exitosa
// 	c.JSON(http.StatusCreated, gin.H{"message": "Alimento creado con éxito!"})
// })

// // Iniciar el servidor en el puerto 8080
// if err := r.Run(":8080"); err != nil {
// 	log.Fatal("Error al iniciar el servidor: ", err)
// }

//PROBAR EL LISTADO DE ALIMENTOS

// r := gin.Default()

// // Inicializar MongoDB y configurar las dependencias
// db := repositories.NewMongoDB()
// alimentoRepo := repositories.NewAlimentoRepository(db)
// alimentoService := Services.NewAlimentoService(alimentoRepo)
// alimentoHandler := handlers.NewAlimentoHandler(alimentoService)

// // Ruta para obtener todos los alimentos (API)
// r.GET("/api/alimentos", alimentoHandler.ObtenerTodosAlimentos)

// // Configuración para servir archivos estáticos desde la carpeta "Front"
// r.Static("/front", filepath.Join(".", "Front"))

// // Ruta para el archivo HTML de listado de alimentos
// r.GET("/alimentos/listado", func(c *gin.Context) {
// 	c.File("./Front/alimentos/index.html")
// })

// // Iniciar el servidor en el puerto 8080
// r.Run(":8080")
