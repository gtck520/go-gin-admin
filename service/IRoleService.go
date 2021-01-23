package service

import "github.com/konger/ckgo/models"

//IRoleService RoleService接口定义
type IRoleService interface {
	//GetUserRoles 分页返回Articles获取用户身份信息
	GetUserRoles(userName string) []*models.Role
}
