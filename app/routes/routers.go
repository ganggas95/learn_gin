package routes

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/gin-gonic/contrib/jwt"
		controllers2 "learn_gin/app/controllers"

)

func InitializeRoutes(router *gin.Engine) {
	auth_router := router.Group("/api/auth")
	auth_router.POST("/login", controllers2.LoginHandler)

	router.POST("/api", controllers2.VerifHandler)
	router.OPTIONS("/api", controllers2.VerifHandler)
	router.GET("/api", controllers2.GetMessage)

	todo_route := router.Group("/api/todo")
	todo_route.Use(jwt.Auth("b1sm1llahirahmanirrahim1234567890987676765656"))
	todo_route.POST("/add", controllers2.AddTodo)
	todo_route.GET("/all", controllers2.GetTodos)
	todo_route.Any("/detail/:id_todo", controllers2.EditTodo)

}
