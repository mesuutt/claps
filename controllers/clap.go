package controllers

import (
	"github.com/mesuutt/claps/utils"
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

	if utils.IsRequestURL(clap.PageURL) == false {
		c.JSON(http.StatusBadRequest, gin.H{"message": "page_url is not a valid URL"})
		return
	}

	err = clap.Increase()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Clap failure"})
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
}

func (o ClapsController) Count(c *gin.Context) {
	var clap models.Clap

	err := c.Bind(&clap)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON"})
		return
	}

	clap.Get()
	msg := make(map[string]uint)
	msg["count"] = clap.Count

	c.JSON(http.StatusOK, msg)
}
