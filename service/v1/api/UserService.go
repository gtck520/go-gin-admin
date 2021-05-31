package api

import (
	"errors"
	"time"

	"github.com/konger/ckgo/common/codes"
	"github.com/konger/ckgo/common/logger"
	models "github.com/konger/ckgo/models/common"

	//pageModel "github.com/konger/ckgo/page"
	"github.com/konger/ckgo/common/helper"
	"github.com/konger/ckgo/common/util/cache"
	"github.com/konger/ckgo/common/util/convert"
	"github.com/konger/ckgo/common/util/hash"
	"github.com/konger/ckgo/repository"

	"github.com/konger/ckgo/common/middleware/jwt"
	"github.com/konger/ckgo/common/util/uuids"
)

// UserService
type UserService struct {
	Repository *repository.UserRepository `inject:""`
	Log        logger.ILogger             `inject:""`
}

//ExistUserByName 判断用户名是否已存在
func (u *UserService) ExistUserByPhone(phone string) bool {
	where := models.User{Phone: phone}
	isExits, _ := u.Repository.ExistUserByName(&where)
	return isExits
}

//AddUser 添加用户
func (u *UserService) AddUser(user *models.User) bool {
	//生成用户私有盐值
	user.Salt = helper.GetCode(4)
	user.UserPass = hash.Md5String(codes.MD5_PREFIX + user.Salt + user.UserPass)

	isOK := u.Repository.AddUser(user)
	if !isOK {
		return false
	}
	return true

}

//Login 用户登录
func (u *UserService) Login(user *models.User) (interface{}, error) {
	where := models.User{Phone: user.Phone}
	isExits, User := u.Repository.ExistUserByName(&where)
	if !isExits {
		return false, errors.New("用户不存在")
	} //生成用户私有盐值
	Salt := User.Salt

	user.UserPass = hash.Md5String(codes.MD5_PREFIX + Salt + user.UserPass)

	isOK, User := u.Repository.CheckUser(user)
	u.Log.Infof("%+v", User)
	if !isOK {
		return false, errors.New("用户密码错误")
	}
	// 缓存或者redis
	uuid := uuids.GetUUID()
	err := cache.Set([]byte(uuid), []byte(convert.ToString(User.ID)), 60*60) // 1H
	if err != nil {
		return false, errors.New("缓存设置失败")
	}
	// token jwt
	userInfo := make(map[string]string)
	userInfo["exp"] = convert.ToString(time.Now().Add(time.Hour * time.Duration(1)).Unix()) // 1H
	userInfo["iat"] = convert.ToString(time.Now().Unix())
	userInfo["uuid"] = uuid
	token := jwt.CreateToken(userInfo)
	// 发至页面
	resData := make(map[string]string)
	resData["token"] = token
	return resData, nil

}
