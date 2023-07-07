package user

import "go-gin-example/internal/model/sys_user_model"

func (s *service) FindById(userId int64) (*sys_user_model.SysUser, error) {
	return sys_user_model.FindById(userId)
}
