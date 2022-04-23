package database

// import packages
import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/joho/godotenv"
	"github.com/kmkz-69/chatr.fr.backend/src/models"
)

// DB represents the database connection
type DB struct {
	Db *gorm.DB
}

var database *DB
// ConnectToDB connects the server with database
func ConnectToDB() {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file \n", err)
	}

	// connect to database
	db, err := gorm.Open(postgres.Open(
		"host="+os.Getenv("DB_HOST")+
			" port="+os.Getenv("DB_PORT")+
			" user="+os.Getenv("DB_USER")+
			" dbname="+os.Getenv("DB_NAME")+
			" password="+os.Getenv("DB_PASSWORD"),
	), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
        os.Exit(2)
	}

	// migrate the schema
	db.AutoMigrate(&models.PrivacyAttributes{})
	db.AutoMigrate(&models.HostAttributes{})
	db.AutoMigrate(&models.AdminAttributes{})

	// set the database connection
	database = &DB{Db: db}
	
}