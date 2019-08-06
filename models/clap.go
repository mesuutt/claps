package models

import (
	"github.com/jinzhu/gorm"
	"github.com/mesuutt/claps/db"
)


// Clap is behaviour objects of clap service
type Clap struct {
	gorm.Model
	PageURL string `json:"page_url" binding:"required", gorm:"type:text;unique_index"`
	Count  uint
}

// Create using for save new clap
func (claps *Clap) Create() error {
	conn := db.GetDB()
	tx := conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(claps).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
