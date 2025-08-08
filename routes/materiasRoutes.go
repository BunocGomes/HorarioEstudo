package routes

import (
	"github.com/BunocGomes/HorarioEstudo/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter configura todas as rotas da aplicação
func SetupRouter(router *gin.Engine) {
	// Agrupa as rotas por versão da API
	api := router.Group("/api/v1")
	{
		// Rotas para Matérias
		materias := api.Group("/materias")
		{
			materias.GET("/", controllers.GetMateriasHandler)
			materias.POST("/", controllers.CreateMateriaHandler)
			// Aqui você adicionaria outras rotas: GET por ID, PUT, DELETE
		}
	}
}