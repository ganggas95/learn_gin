package controllers

import (
	"github.com/gin-gonic/gin"
	"learn_gin/app/models"
	"learn_gin/app/db"
	"net/http"
)

type UserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

func LoginHandler(c *gin.Context){
	var user_req UserReq
	c.BindJSON(&user_req)
	var user models.User
	db.DbEngine.Where("username=?", user_req.Username).Get(&user)
	err := models.CheckPassword([]byte(user.Password), user_req.Password)
	if err != nil{
		c.JSON(http.StatusUnauthorized, gin.H{
			"data": nil,
			"error": true,
			"message": "Username and password invalid",
		})
	}
	token, err := models.GenerateToken(&user)
	if err!=nil{
		c.JSON(http.StatusOK, gin.H{
			"data": &user,
			"error": false,
			"message": string(err.Error()),
		})
	}
	token_resp := TokenResponse{Token:token}
	c.JSON(http.StatusOK, gin.H{
		"data": token_resp,
		"error": false,
		"message": "Username and password valid",
	})
}
