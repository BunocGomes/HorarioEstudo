module github.com/BunocGomes/HorarioEstudo

go 1.22

require (
	github.com/gin-gonic/gin v1.10.0
	github.com/joho/godotenv v1.5.1
	gorm.io/driver/postgres v1.5.9
	gorm.io/gorm v1.25.10
)

// ... outras dependências ...

// Adicione esta linha no final
replace github.com/BunocGomes/HorarioEstudo => ./