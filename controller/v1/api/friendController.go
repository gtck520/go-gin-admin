package api

import (
	//"net/http"
	//"strconv"

	"github.com/konger/ckgo/common/codes"
	"github.com/konger/ckgo/common/logger"
	"github.com/konger/ckgo/controller/common"

	"github.com/gin-gonic/gin"
	service "github.com/konger/ckgo/service/v1/api"
	//"net/http"
)

//### 如果是使用Go Module,gin-jwt模块应使用v2
//下载安装，开启Go Module "go env -w GO111MODULE=on",然后执行"go get github.com/appleboy/gin-jwt/v2"
//导入应写成 import "github.com/appleboy/gin-jwt/v2"
//### 如果不是使用Go Module
//下载安装gin-jwt，"go get github.com/appleboy/gin-jwt"
//导入import "github.com/appleboy/gin-jwt"

//User 注入UserService
type Friend struct {
	Log     logger.ILogger         `inject:""`
	Service *service.FriendService `inject:""`
}

// @Summary 获取好友列表
// @Description 获取好友列表
// @Tags 用户接口
// @Produce  json
// @Success 200 {string} json "{"code":200,"data":{},"message":"ok"}"
// @Router /v1/api/friend/friend_list [get]

func (u *Friend) FriendList(c *gin.Context) {
	user_id, _ := c.Get(codes.USER_ID_Key)
	friendlist := u.Service.GetFriendList(user_id.(uint))
	common.ResSuccess(c, friendlist)
}
