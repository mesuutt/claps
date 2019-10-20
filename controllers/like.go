package controllers

import (
	"fmt"
	"github.com/mesuutt/claps/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mesuutt/claps/models"
)

type LikeController struct{}

func (o LikeController) Increase(c *gin.Context) {
	var like models.Like
	err := c.BindJSON(&like)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error_message": "Invalid JSON."})
		return
	}

	if !strings.HasPrefix(like.PageURL, like.Domain) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": fmt.Sprintf("Only allowed from same origin. page_url must contains %s", like.Domain),
		})
		return
	}

	if !utils.IsRequestURL(like.PageURL) {
		c.JSON(http.StatusBadRequest, gin.H{"error_message": "page_url is not a valid URL"})
		return
	}

	like.Domain = utils.GetHost(c.Request.Header.Get("Origin"))
	err = like.Increase()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error_message": "Like failure"})
		return
	}

	msg := make(map[string]uint)
	msg["count"] = like.Count

	c.JSON(http.StatusOK, msg)

	c.Writer.WriteHeader(http.StatusOK)
}


func (o LikeController) Decrease(c *gin.Context) {
	var like models.Like

	err := c.BindJSON(&like)
	if err != nil {
		if like.PageURL == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error_message": "page_url field is required"})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error_message": "Invalid JSON."})
		return
	}

	if !strings.HasPrefix(like.PageURL, c.Request.Header.Get("Origin")) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": fmt.Sprintf("Only allowed from same origin. page_url must contains %s", c.Request.Header.Get("Origin")),
		})
		return
	}

	if !utils.IsRequestURL(like.PageURL) {
		c.JSON(http.StatusBadRequest, gin.H{"error_message": "page_url is not a valid URL"})
		return
	}

	like.Domain = utils.GetHost(c.Request.Header.Get("Origin"))
	err = like.Decrease()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error_message": "Like count decrease failure"})
		return
	}

	msg := make(map[string]uint)
	msg["count"] = like.Count

	c.JSON(http.StatusOK, msg)

	c.Writer.WriteHeader(http.StatusOK)
}


func (o LikeController) Count(c *gin.Context) {
	var like models.Like

	err := c.BindJSON(&like)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error_message": "Invalid JSON"})
		return
	}

	like.Domain = utils.GetHost(c.Request.Header.Get("Origin"))
	like.Get()
	msg := make(map[string]uint)
	msg["count"] = like.Count

	c.JSON(http.StatusOK, msg)

}
