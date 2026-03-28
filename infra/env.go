package infra

// import (
// 	"log"
// 	"os"

// 	"github.com/joho/godotenv"
// )

// type AppConfig struct {
// 	Port         string
// 	DBPassword   string
// 	DBPort       string
// 	DBName       string
// 	DBHost       string
// 	DBUser       string

// }

// var Configurations AppConfig

// func InitEnv() {
// 	// Load .env file
// 	err := godotenv  .Load()
// 	godotenv.Load()
// 	// godotenv.Load()
// 	godotenv.Load()
// 	godotenv.Load()


// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// 	// Variable ka dhig qaar aplicarionkeenu isticmaali karo
// 	// Assign values from environment variables
// 	Configurations.Port = os.Getenv("PORT")
// 	Configurations.DBHost = os.Getenv("DB_HOST")
// 	Configurations.DBUser = os.Getenv("DB_USER")
// 	Configurations.DBPassword = os.Getenv("DB_PASSWORD")
// 	Configurations.DBPort = os.Getenv("DB_PORT")
// 	Configurations.DBName = os.Getenv("DB_NAME")
// } 