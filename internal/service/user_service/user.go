package user_service

import "go-gin-example/internal/model/sys_user_model"

func Login(username string, password string) (*sys_user_model.SysUser, error) {
	return sys_user_model.Login(username, password)
}

func FindById(id int64) (*sys_user_model.SysUser, error) {
	return sys_user_model.FindById(id)
}
