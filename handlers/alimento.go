package handlers

import (
	"Spaghetti-Code/Services"
	"Spaghetti-Code/dto"

	"net/http"

	"github.com/gin-gonic/gin"
	//"log"
)

// FUNCIONA
func (handler *AlimentoHandler) CrearAliment(c *gin.Context) {
	var aliment dto.AlimentoDto
	if err := c.ShouldBindJSON(&aliment); err != nil {
		// Si hay un error en el JSON, devuelve un error
		c.JSON(400, gin.H{"error": err.Error()}) //Preguntar si se usa http o no
		return
	} //cambiar si falla "http.StatusBadRequest" por "400"

	if err := handler.service.CrearAliment(aliment); err != nil {
		c.JSON(500, gin.H{"error": "No se pudo intertar el alimento"})
		return
	}

	c.JSON(200, gin.H{"message": "Alimento insertado exitosamente!"})
}

type AlimentoHandler struct {
	service Services.AlimentoInterface
}

func NewAlimentoHandler(service Services.AlimentoInterface) *AlimentoHandler {
	return &AlimentoHandler{service: service}
}

// LISTO (SIN PROBAR)
func (handler AlimentoHandler) ModificarAlimento(c *gin.Context) {
	var aliment dto.AlimentoDto

	if err := c.ShouldBindJSON(&aliment); err != nil {
		// Si hay un error en el JSON, devuelve un error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar el JSON: " + err.Error()})
		return
	}

	resultado := handler.service.ModificarAlimentos(aliment)

	c.JSON(http.StatusCreated, resultado)
}

// FUNCIONA
/*
func (handler *AlimentoHandler) ObtenerTodosAlimentos(c *gin.Context) {
	alimentos, err := handler.service.ObtenerTodosAlimentos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener alimentos"})
		return
	}
	c.JSON(http.StatusOK, alimentos)
}*/

func (handler *AlimentoHandler) ObtenerTodosAlimentosPorID(c *gin.Context) {
	usuarioID := c.Param("usuario_id") // Obtener el ID del usuario desde la URL

	// Llamar al servicio para obtener los alimentos por ID de usuario
	alimentos, err := handler.service.ObtenerTodosAlimentosPorID(usuarioID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener alimentos"})
		return
	}
	c.JSON(http.StatusOK, alimentos)
}

//FALTA EDITAR

/*
func (handler *AlimentoHandler) EliminarAlimento(c *gin.Context) {
	id := c.Param("id")

	aliment := handler.service.EliminarAlimento(id)
	aliment := handler.service.EliminarAlimento(id)

	c.JSON(http.StatusOK, aliment)
}
*/
