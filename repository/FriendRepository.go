package repository

import (
	"github.com/konger/ckgo/common/logger"
	models "github.com/konger/ckgo/models/common"
)

//FriendRepository 注入IDb
type FriendRepository struct {
	Log  logger.ILogger `inject:""`
	Base BaseRepository `inject:"inline"`
}

//AddFriend 新建朋友关系
func (a *FriendRepository) AddFriend(friend *models.Friend) bool {
	if err := a.Base.Create(friend); err != nil {
		a.Log.Errorf("新建好友关系失败", err)
		return false
	}
	return true
}

//ExistUserByName
func (a *FriendRepository) GetFriendList(where interface{}) (bool, interface{}) {
	var friends []models.Friend
	if err := a.Base.Find(where, &friends, ""); err != nil {
		a.Log.Errorf("获取好友列表异常", err)
	}
	return true, &friends
}
