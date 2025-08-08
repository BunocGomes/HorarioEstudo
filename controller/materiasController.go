package controllers

import (
	"net/http"
	"github.com/BunocGomes/HorarioEstudo/dto"
	"github.com/BunocGomes/HorarioEstudo/services"

	"github.com/gin-gonic/gin"
)

// GetMateriasHandler lida com a requisição para buscar todas as matérias
func GetMateriasHandler(c *gin.Context) {
	materias, err := services.GetAllMaterias()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar matérias"})
		return
	}
	c.JSON(http.StatusOK, materias)
}

// CreateMateriaHandler lida com a requisição para criar uma nova matéria
func CreateMateriaHandler(c *gin.Context) {
	var materiaDTO dto.CreateMateriaDTO

	if err := c.ShouldBindJSON(&materiaDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	novaMateria, err := services.CreateMateria(materiaDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar matéria"})
		return
	}

	c.JSON(http.StatusCreated, novaMateria)
}