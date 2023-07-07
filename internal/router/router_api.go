package router

import (
	"github.com/gin-gonic/gin"
	user_api "go-gin-example/internal/api/user"
	"go-gin-example/pkg/model"
)

func InitRouterApi(r *gin.Engine) {

	rApi := r.Group("/api")

	user := user_api.New(model.DB)
	rApi.POST("/login", user.Login)
	rApi.POST("/user", user.Create)
}
