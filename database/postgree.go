package database

import (
	"fmt"
	"log"
	"github.com/BunocGomes/HorarioEstudo/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB inicializa a conexão com o banco de dados PostgreSQL
func ConnectDB() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Falha ao conectar ao banco de dados: ", err)
	}

	fmt.Println("Conexão com o banco de dados estabelecida com sucesso.")

	// AutoMigrate vai criar ou atualizar as tabelas do banco de dados
	// com base nos seus models
	err = DB.AutoMigrate(&models.Materia{})
	if err != nil {
		log.Fatal("Falha na migração do banco de dados: ", err)
	}
	fmt.Println("Migração do banco de dados concluída.")
}