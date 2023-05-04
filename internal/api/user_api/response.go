package user_api

import (
	"go-gin-example/internal/model/sys_user_model"
	"go-gin-example/pkg/util/copy_struct"
	"time"
)

type LoginRes struct {
	Token  string `json:"token"`
	Expire int    `json:"expire"`
}

type UserRes struct {
	ID           int64     `json:"id"`
	Username     string    `json:"username"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	Nickname     string    `json:"nickname"`
	Salt         string    `json:"salt"`
	Phone        string    `json:"phone"`
	Gender       int16     `json:"gender"` // 性别，0：女，1：男，默认1
	Head         string    `json:"head"`
	Remark       string    `json:"remark"`
	State        int16     `json:"state"`        // 状态，0：禁用，1：启用，2：锁定
	DepartmentID int64     `json:"departmentId"` // 部门id
	RoleID       int64     `json:"roleId"`
	Email        string    `json:"email"`
}

func ofUserRes(user sys_user_model.SysUser) *UserRes {
	var res UserRes
	err := copy_struct.CopyStruct(user, res)
	if err != nil {
		return nil
	}
	return &res
}
