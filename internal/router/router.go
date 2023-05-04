package router

import (
	"github.com/gin-gonic/gin"
	"go-gin-example/internal/api/user_api"
)

func InitRouter(r *gin.Engine) {

	rApi := r.Group("/api")

	rApi.POST("/login", user_api.Login)
	rApi.POST("/logout", user_api.Logout)
	rApi.GET("/profile", user_api.Profile)
}
