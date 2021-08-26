package api

import (
	//"net/http"
	//"strconv"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/konger/ckgo/common/codes"
	"github.com/konger/ckgo/common/logger"
	"github.com/konger/ckgo/common/util/convert"
	"github.com/konger/ckgo/controller/common"
	"github.com/konger/ckgo/page"
	service "github.com/konger/ckgo/service/v1/api"
	"github.com/konger/ckgo/websocket"
	//"net/http"
)

//User 注入UserService
type Ws struct {
	Log      logger.ILogger           `inject:""`
	WsServer *websocket.ServerManager `inject:""`
	Uservice *service.UserService     `inject:""`
	Fservice *service.FriendService   `inject:""`
}

// @Summary 绑定客户端
// @Description 绑定客户端
// @Tags ws接口
// @Produce  json
// @Param phone query string true "18612345678"
// @Param password query string true "123456"
// @Param code query string true "1234"
// @Success 200 {string} json "{"code":200,"data":{},"message":"ok"}"
// @Router /v1/api/user/register [post]

func (u *Ws) SendMessage(c *gin.Context) {
	sentdata := page.SendData{}
	err := c.ShouldBind(&sentdata)
	if err != nil {
		common.ResFail(c, err.Error())
		return
	}
	user := u.WsServer.Connlist[convert.ToString(sentdata.ToId)]
	//进行一系列消息存储操作
	if user == nil {
		//对方未上线 存储消息后返回成功
		common.ResSuccess(c, sentdata)
	} else {
		log.Println("shiti:", user)
		u.WsServer.SendMessage(codes.SENDTYPE_CLIENT, user, websocket.TypeMessage{"message", sentdata.Msg}, c, codes.MESSAGETYPE_TEXT)
		common.ResSuccess(c, sentdata)
	}
}
