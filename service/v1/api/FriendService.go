package api

import (
	"github.com/konger/ckgo/common/logger"
	models "github.com/konger/ckgo/models/common"

	//pageModel "github.com/konger/ckgo/page"

	"github.com/konger/ckgo/repository"
)

// FriendService
type FriendService struct {
	Repository *repository.FriendRepository `inject:""`
	Log        logger.ILogger               `inject:""`
}

//ExistUserByName 判断用户名是否已存在
func (u *FriendService) GetFriendList(UserId uint) interface{} {
	userid := uint64(UserId)
	// where := repository.Where{
	// 	"or",
	// 	make(map[string]interface{}),
	// 	make(map[string]interface{}),
	// }
	// where.Op = "or"
	// where.Wh1["user_id"] = userid
	// where.Wh2["friend_id"] = userid
	where := models.Friend{UserId: userid}
	ok, list := u.Repository.GetFriendList(&where)
	if !ok {
		u.Log.Errorf("获取列表失败")
	}
	return list
}

//AddUser 添加用户
func (u *FriendService) AddFriend(friend *models.Friend) bool {

	isOK := u.Repository.AddFriend(friend)
	if !isOK {
		return false
	}
	return true

}
