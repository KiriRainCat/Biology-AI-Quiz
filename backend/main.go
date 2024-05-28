package main

import (
	"log"
	"net/http"
	"os"
	"stp/internal"
	"stp/internal/middleware"
	"stp/internal/pkg/config"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialization
	config.Init()
	// internal.InitDB()
	internal.InitLLM()

	// Create gin-engine and base router-group
	server := gin.New()

	server.Use(func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "*")
	})

	r := server.Group("/api")
	r.Use(gin.LoggerWithWriter(os.Stdout, "/api/ping")).
		Use(gin.Recovery()).
		Use(middleware.Auth.Authenticate)

	//* --------------------------- API Registration --------------------------- *//
	// PING API
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "Pong",
			"data": nil,
		})
	})

	// Register API Routes
	c := internal.Controller
	r.GET("/generate-quiz", c.GenerateQuiz)
	r.GET("/records", c.GetRecords)
	r.POST("/record", c.PutRecord)

	log.Println("启动成功")
	server.Run(":" + config.Config.Server.Port)
}
