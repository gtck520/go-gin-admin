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

//ExistUserByName 判断用户名是否已存在
func (u *UserService) ExistUserByPhone(phone string) bool {
	where := models.User{Phone: phone}
	return u.Repository.ExistUserByName(&where)
}

//AddUser 添加用户
func (u *UserService) AddUser(user *models.User) bool {
	isOK := u.Repository.AddUser(user)
	if !isOK {
		return false
	}
	return true

}
