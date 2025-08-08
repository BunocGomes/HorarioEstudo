package services

import (
	"github.com/BunocGomes/HorarioEstudo/database"
	"github.com/BunocGomes/HorarioEstudo/dto"
	"github.com/BunocGomes/HorarioEstudo/models"
)

// GetAllMaterias busca todas as matérias no banco de dados
func GetAllMaterias() ([]dto.MateriaResponseDTO, error) {
	var materias []models.Materia
	if err := database.DB.Find(&materias).Error; err != nil {
		return nil, err
	}

	var materiasDTO []dto.MateriaResponseDTO
	for _, m := range materias {
		materiasDTO = append(materiasDTO, dto.MateriaResponseDTO{
			ID:        m.ID,
			Nome:      m.Nome,
			Professor: m.Professor,
			Descricao: m.Descricao,
		})
	}
	return materiasDTO, nil
}

// CreateMateria cria uma nova matéria no banco de dados
func CreateMateria(materiaDTO dto.CreateMateriaDTO) (dto.MateriaResponseDTO, error) {
	materia := models.Materia{
		Nome:      materiaDTO.Nome,
		Professor: materiaDTO.Professor,
		Descricao: materiaDTO.Descricao,
	}

	if err := database.DB.Create(&materia).Error; err != nil {
		return dto.MateriaResponseDTO{}, err
	}

	responseDTO := dto.MateriaResponseDTO{
		ID:        materia.ID,
		Nome:      materia.Nome,
		Professor: materia.Professor,
		Descricao: materia.Descricao,
	}

	return responseDTO, nil
}