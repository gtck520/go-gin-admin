package api

import (
	"github.com/konger/ckgo/common/logger"
	models "github.com/konger/ckgo/models/common"
	//pageModel "github.com/konger/ckgo/page"
	//"github.com/konger/ckgo/page/emun"
	"github.com/konger/ckgo/repository"
)

// UserService 注入IUserRepository
type UserService struct {
	Repository     repository.IUserRepository `inject:""`
	Log            logger.ILogger             `inject:""`
}
//AddUser 新建用户，同时新建用户角色
func (a *UserService) AddUser(user *models.User) bool {
	//用业务逻辑实现事务效果
	isOK := a.Repository.AddUser(user)
	if !isOK {
		return false
	}
	return true

}
