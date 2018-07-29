package main

import (
	"github.com/gin-gonic/gin"
	"learn_gin/routes"
	"learn_gin/db"
		"github.com/go-xorm/xorm"
	"github.com/gin-contrib/cors"
	"time"
)

var router *gin.Engine
var DBEngine *xorm.Engine

func main(){
	router = gin.Default()
	gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"PUT", "POST", "DELETE", "GET"},
		AllowHeaders:     []string{"Origin", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:8080"
		},
		MaxAge: 12 * time.Hour,
	}))

	DBEngine = db.InitDB()

	routes.InitializeRoutes(router)
	router.Run("localhost:8017")
}
