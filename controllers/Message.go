package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMessage(c *gin.Context){
	println("Hello")
	message, _ := c.GetQuery("m")
	c.String(http.StatusOK, "Message : "+ message)
}
