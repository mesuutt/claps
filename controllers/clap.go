package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mesuutt/claps/models"
)

type ClapsController struct{}

func (o ClapsController) Add(c *gin.Context) {
	var clap models.Clap

	err := c.Bind(&clap)
	if err != nil {
		if clap.PageURL == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "page_url field is required"})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"message": "Claps cannot added."})
		return
	}

	err = clap.Create()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Clap failure"})
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
}
