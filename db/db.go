package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/mesuutt/claps/config"
)

var db *gorm.DB

// Init Initialize DB
func Init() {
	c := config.GetConfig()

	dbHost := c.GetString("db.host")
	dbName := c.GetString("db.name")
	dbUser := c.GetString("db.user")
	dbPort := c.GetString("db.port")
	dbPassword := c.GetString("db.password")
	dbSSLMode := c.GetString("db.sslmode")

	dbString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", dbHost, dbPort, dbUser, dbName, dbPassword, dbSSLMode)

	var err error
	db, err = gorm.Open("postgres", dbString)
	if c.GetBool("debug") == true {
		db.LogMode(true)
	}

	if err != nil {
		panic("failed to connect database")
	}
}

// GetDB gets current using db connection
func GetDB() *gorm.DB {
	return db
}
