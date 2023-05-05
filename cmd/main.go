package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-gin-example/internal/router"
	"go-gin-example/pkg/config"
	"go-gin-example/pkg/model"
)

func init() {
	config.Setup()
	model.Setup()
}

func main() {
	r := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(corsConfig))
	router.InitRouter(r)
	r.Run(fmt.Sprintf(":%d", config.Config.Server.Port))
}
