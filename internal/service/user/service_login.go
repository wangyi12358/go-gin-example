package user

import "go-gin-example/internal/model/sys_user_model"

func (s *service) Login(username string, password string) (*sys_user_model.SysUser, error) {
	return sys_user_model.Login(username, password)
}
