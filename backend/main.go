package main

import (
	"log"
	"os"
	"stp/internal"
	"stp/internal/middleware"
	"stp/internal/pkg/config"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialization
	config.Init()
	internal.InitDB()
	internal.InitLLM()

	// Create gin-engine and base router-group
	server := gin.New()

	server.Use(func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "*")
	})

	r := server.Group("/api")
	r.Use(gin.LoggerWithWriter(os.Stdout)).
		Use(gin.Recovery()).
		Use(middleware.Auth.Authenticate)

	//* --------------------------- API Registration --------------------------- *//

	// Register API Routes
	c := internal.Controller
	r.GET("/generate-quiz", c.GenerateQuiz)
	r.GET("/record/:name", c.GetRecord)
	r.GET("/records", c.GetRecords)
	r.POST("/record", c.PutRecord)

	log.Println("启动成功")
	server.Run(":" + config.Config.Server.Port)
}
