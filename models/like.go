package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/mesuutt/claps/db"
)

// Clap is behaviour objects of clap service
type Like struct {
	gorm.Model
	Domain     string `json:"-" gorm:"type:varchar(50);"`
	Identifier string `json:"identifier" gorm:"type:varchar(250);"`
	PageURL    string `json:"page_url" binding:"required" gorm:"-"`
	Count      uint
}

// Create using for save new clap
func (like *Like) Increase() error {
	conn := db.GetDB()
	tx := conn.Begin()

	if err := tx.Error; err != nil {
		return err
	}

	tx = tx.Set("gorm:query_option", "FOR UPDATE")

	if err := tx.Where("domain = ? and identifier = ?", like.Domain, like.Identifier).First(like).Error; err != nil {
		like.Count = 1
		if err := tx.Create(like).Error; err != nil {
			tx.Rollback()
			return err
		}
	} else {
		like.Count = like.Count + 1
		tx.Save(like)
	}

	return tx.Commit().Error
}

// Get clap from db
func (like *Like) Get() {
	conn := db.GetDB()
	conn.Where("domain = ? and identifier = ?", like.Domain, like.Identifier).First(like)
}


// Create using for save new clap
func (like *Like) Decrease() error {
	conn := db.GetDB()
	tx := conn.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	tx = tx.Set("gorm:query_option", "FOR UPDATE")
	if err := tx.Where("domain = ? and identifier = ?", like.Domain, like.Identifier).First(like).Error; err != nil {
		tx.Rollback()
		return err
	} else {
		if like.Count == 0 {
			return errors.New("like cannot decrease anymore")
		}

		like.Count = like.Count - 1
		tx.Save(like)
	}

	return tx.Commit().Error
}
