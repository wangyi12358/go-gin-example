package sys_user_model

import (
	"doc-server/models"
	"fmt"
)

func Login(username string, password string) (*SysUser, error) {
	var sysUser SysUser
	err := models.DB.Where(&SysUser{
		Username: username,
		Password: password,
	}).First(sysUser).Error
	if err != nil {
		fmt.Printf("sql error: %s\n", err.Error())
		return nil, err
	}
	return &sysUser, nil
}

func FindById(id int64) (*SysUser, error) {
	var user SysUser
	if err := models.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
