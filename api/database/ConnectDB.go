package database

import (
	"fmt"
	"log"
	"os"
	"chat_app/api/models"

	"gorm.io/driver/postgres"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct{
	DB *gorm.DB
	MYSQLDB *gorm.DB
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
	
	// d.mysqlConnection()

	// d.MYSQLDB.AutoMigrate(&models.Chat{}) 
	// d.MYSQLDB.AutoMigrate(&models.User{})
	// d.MYSQLDB.AutoMigrate(&models.Message{})
	// d.MYSQLDB.AutoMigrate(&models.UsersChat{})
	//d.MYSQLDB.AutoMigrate(&models.Person{})

	// d.DB.AutoMigrate(&models.User{})
	// d.DB.AutoMigrate(&models.Chat{})
	// d.DB.AutoMigrate(&models.Message{})
	// d.DB.AutoMigrate(&models.UsersChat{})

	if(d.autoMigrateDatabase){
		d.updateDatabase()
	}
}

func (d *DB) mysqlConnection(){
	password := "nintendowiiu000"
	user     := "root"
	dbname   := "chat_app"
	dsn      := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil{
		log.Fatal(err)
	}
	
	d.MYSQLDB = db
}

func (d *DB) updateDatabase(){
	d.DB.Migrator().DropTable(&models.Message{}, &models.UsersChat{})
	d.DB.Migrator().CreateTable( &models.Message{}, &models.UsersChat{})
}