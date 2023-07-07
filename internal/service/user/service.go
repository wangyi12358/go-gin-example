package user

import (
	"go-gin-example/internal/model/sys_user_model"
	"gorm.io/gorm"
)

type Service interface {
	Create() (*sys_user_model.SysUser, error)
	Login(username string, password string) (*sys_user_model.SysUser, error)
	FindById(userId int64) (*sys_user_model.SysUser, error)
}

type service struct {
	db *gorm.DB
}

func New(db *gorm.DB) Service {
	return &service{
		db: db,
	}
}
