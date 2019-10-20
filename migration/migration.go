package migration

import (
	"fmt"
	"github.com/mesuutt/claps/db"
	"github.com/mesuutt/claps/models"
)

// Migrate migrate scheme changes to db
func Migrate() {
	conn := db.GetDB()
	fmt.Println("Migrations running")
	conn.AutoMigrate(models.Clap{})
	conn.AutoMigrate(models.Like{})
}
