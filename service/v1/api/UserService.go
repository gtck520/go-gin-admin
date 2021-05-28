package api

import (
	"github.com/konger/ckgo/common/logger"
	models "github.com/konger/ckgo/models/common"

	//pageModel "github.com/konger/ckgo/page"
	//"github.com/konger/ckgo/page/emun"
	"github.com/konger/ckgo/repository"
)

// UserService
type UserService struct {
	Repository *repository.UserRepository `inject:""`
	Log        logger.ILogger             `inject:""`
}

//AddUser 添加用户
func (u *UserService) AddUser(user *models.User) bool {
	//用业务逻辑实现事务效果
	isOK := u.Repository.AddUser(user)
	if !isOK {
		return false
	}
	return true

}
