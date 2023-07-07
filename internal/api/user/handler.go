package user

import (
	"github.com/gin-gonic/gin"
	"go-gin-example/internal/service/user"
	"gorm.io/gorm"
)

type Handler interface {
	Create(c *gin.Context)
	Login(c *gin.Context)
}

type handler struct {
	db      *gorm.DB
	service user.Service
}

func New(db *gorm.DB) Handler {
	return &handler{
		db:      db,
		service: user.New(db),
	}
}
