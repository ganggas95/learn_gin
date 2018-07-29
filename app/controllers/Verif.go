package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"learn_gin/app/models"
)

func VerifHandler(c *gin.Context){
	if c.Request.Method == "OPTIONS" {
		c.Header("Allow", "POST, GET, OPTIONS")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "origin, content-type, accept")
		c.Header("Content-Type", "application/json")
		c.Status(http.StatusOK)
	}else if c.Request.Method == "POST" {
		var user models.User
		c.BindJSON(&user)
		c.JSON(http.StatusOK, gin.H{
			"username": user.Username,
			"password": user.Password,
		})
	}
}
