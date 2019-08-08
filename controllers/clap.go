package controllers

import (
	"fmt"
	"github.com/mesuutt/claps/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mesuutt/claps/models"
)

type ClapsController struct{}

func (o ClapsController) Add(c *gin.Context) {
	var clap models.Clap

	err := c.BindJSON(&clap)
	if err != nil {
		if clap.PageURL == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "page_url field is required"})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON."})
		return
	}

	if !strings.HasPrefix(clap.PageURL, c.Request.Header.Get("Origin")) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Only allowed from same origin. page_url must contains %s",  c.Request.Header.Get("Origin")),
		})
		return
	}

	if !utils.IsRequestURL(clap.PageURL) {
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

	err := c.BindJSON(&clap)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON"})
		return
	}

	clap.Get()
	msg := make(map[string]uint)
	msg["count"] = clap.Count

	c.JSON(http.StatusOK, msg)

}
