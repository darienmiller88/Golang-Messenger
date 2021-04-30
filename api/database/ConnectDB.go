package database

import (
	"fmt"
	"log"
	"os"
	"chat_app/api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct{
	DB *gorm.DB
	autoMigrateDatabase bool
}

//Boolean to decide whether or not to allow automatic dropping and creation of every table in the database.
func (d *DB) InitDB(autoMigrateDatabase bool){
	var err error
	d.autoMigrateDatabase = autoMigrateDatabase
	db_host     := os.Getenv("POSTGRES_HOST")
	db_name     := os.Getenv("POSTGRES_DBNAME")
	db_username := os.Getenv("POSTGRES_USERNAME")
	db_password := os.Getenv("POSTGRES_PASSWORD")
	db_url      := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", db_host, db_username, db_password, db_name)
	d.DB, err   = gorm.Open(postgres.Open(db_url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil{
		log.Fatal(err)
	}
	
	if(d.autoMigrateDatabase){
		d.updateDatabase()
	}
}

func (d *DB) updateDatabase(){
	d.DB.Migrator().DropTable(&models.User{}, &models.Message{})
	d.DB.Migrator().CreateTable(&models.User{}, &models.Message{})
}