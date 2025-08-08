package main

import (
	"log"
	"github.com/BunocGomes/HorarioEstudo/database"
	"github.com/BunocGomes/HorarioEstudo/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Carrega as variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	// Conecta ao banco de dados
	database.ConnectDB()

	// Inicializa o roteador do Gin
	router := gin.Default()

	// Configura as rotas da aplicação
	routes.SetupRouter(router)

	// Inicia o servidor na porta definida no .env
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080" // Porta padrão
	}
	router.Run(":" + port)
}