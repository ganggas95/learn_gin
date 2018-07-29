package controllers

import (
	"github.com/gin-gonic/gin"
	"learn_gin/models"
	"net/http"
	"learn_gin/db"
	)

func AddTodo(c *gin.Context){
	if c.Request.Method == "POST" {
		var todo models.Todo
		c.BindJSON(&todo)
		db.DbEngine.Insert(&todo)
		c.JSON(http.StatusCreated, gin.H{
			"data": &todo,
			"error": false,
			"message": "Success add todo",
		})

	}
}

func EditTodo(c *gin.Context){
	id_todo := c.Param("id_todo")
	var todo models.Todo
	if c.Request.Method == "GET" {

		db.DbEngine.Where("id=?", id_todo).Get(&todo)
		c.JSON(http.StatusCreated, gin.H{
			"data": &todo,
			"error": false,
			"message": "Success get todo",
		})
	}
	if c.Request.Method == "POST" {
		c.BindJSON(&todo)
		db.DbEngine.Where("id=?", id_todo).Update(&todo)
		c.JSON(http.StatusCreated, gin.H{
			"data": &todo,
			"error": false,
			"message": "Success edit todo",
		})
	}
	if c.Request.Method == "DELETE" {
		db.DbEngine.ID(id_todo).Delete(&todo)
		c.JSON(http.StatusOK, gin.H{
			"data": id_todo,
			"error": false,
			"message": "success remove todo",
		})
	}
}

func GetTodos(c *gin.Context){
	var todos []models.Todo
	db.DbEngine.Find(&todos)
	c.JSON(http.StatusOK, gin.H{
		"data": &todos,
		"error": false,
		"message": "Success get todo list",
	})
}
