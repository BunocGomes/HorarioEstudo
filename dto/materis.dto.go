package dto

// CreateMateriaDTO é usado para receber dados para criar uma nova matéria
type CreateMateriaDTO struct {
	Nome      string `json:"nome" binding:"required"`
	Professor string `json:"professor"`
	Descricao string `json:"descricao"`
}

// MateriaResponseDTO é usado para enviar os dados de uma matéria como resposta
type MateriaResponseDTO struct {
	ID        uint   `json:"id"`
	Nome      string `json:"nome"`
	Professor string `json:"professor"`
	Descricao string `json:"descricao"`
}