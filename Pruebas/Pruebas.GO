package handlers

import (
	"net/http"
	"tu_ruta/services"

	"github.com/gin-gonic/gin"
)

type AlimentoHandler struct {
	service services.AlimentoServiceInterface
}

func NewAlimentoHandler(service services.AlimentoServiceInterface) *AlimentoHandler {
	return &AlimentoHandler{service: service}
}

// GetAlimentosPorUsuarioID obtiene una lista de alimentos por usuario_id
func (h *AlimentoHandler) GetAlimentosPorUsuarioID(c *gin.Context) {
	usuarioID := c.Param("usuario_id")
	alimentos, err := h.service.BuscarAlimentosPorUsuarioID(usuarioID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener alimentos"})
		return
	}
	c.JSON(http.StatusOK, alimentos)
}
