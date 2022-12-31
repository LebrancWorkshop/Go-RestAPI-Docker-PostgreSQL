package database 

import (
	"fmt"
	"os"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"github.com/LebrancWorkshop/Go-RestAPI-Docker-PostgreSQL/models" 
)

type Dbinstance struct {
	Db *gorm.DB 
}

var DB Dbinstance; 

func ConnectDb() { 
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Bangkok",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}); 

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err);
		os.Exit(2);
	}

	log.Println("Connected.");
	db.Logger = logger.Default.LogMode(logger.Info);

	log.Println("Running Migration.");
	db.AutoMigrate(&models.Quiz{});

	DB = Dbinstance{
		Db: db, 
	}
}