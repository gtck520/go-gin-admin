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

//ExistUserByName where or 暂时不用
func (a *FriendRepository) GetFriendList(where interface{}) (bool, interface{}) {
	// var friends []models.Friend
	type result struct {
		FriendId int64  `json:"friend_id"`
		UserName string `json:"user_name"`
		NickName string `json:"nick_name"`
	}
	res := []result{}
	wherenew, ok := where.(*Where)
	if ok {
		where = wherenew
	}
	if err := a.Base.Join("go_com_friend", "left join go_com_user on go_com_friend.friend_id = go_com_user.id", where, &res, "go_com_friend.friend_id,go_com_user.user_name,go_com_user.nick_name"); err != nil {
		a.Log.Errorf("获取好友列表异常", err)
	}
	return true, &res
}
