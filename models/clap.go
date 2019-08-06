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
func (clap *Clap) Increase() error {
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

	tx = tx.Set("gorm:query_option", "FOR UPDATE")
	if err := tx.Where("page_url=?", clap.PageURL).First(clap).Error; err != nil {
		clap.Count = 1
		if err := tx.Create(clap).Error; err != nil {
			tx.Rollback()
			return err
		}
	} else {
		clap.Count = clap.Count +1
		tx.Save(clap)
	}

	return tx.Commit().Error
}

// Create using for save new clap
func (clap *Clap) Get() {
	conn := db.GetDB()
	conn.Where("page_url=?", clap.PageURL).First(clap)
}
