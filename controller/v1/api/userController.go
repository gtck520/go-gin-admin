package api

import (
	//"net/http"
	//"strconv"

	//jwt "github.com/appleboy/gin-jwt/v2"
	//"github.com/konger/ckgo/common/codes"
	"github.com/konger/ckgo/common/logger"
	"github.com/konger/ckgo/controller/common"
	//models "github.com/konger/ckgo/models/common"
	//"github.com/konger/ckgo/page"
	service "github.com/konger/ckgo/service/v1/api"
	"github.com/gin-gonic/gin"
	//"net/http"
)

//### 如果是使用Go Module,gin-jwt模块应使用v2
//下载安装，开启Go Module "go env -w GO111MODULE=on",然后执行"go get github.com/appleboy/gin-jwt/v2"
//导入应写成 import "github.com/appleboy/gin-jwt/v2"
//### 如果不是使用Go Module
//下载安装gin-jwt，"go get github.com/appleboy/gin-jwt"
//导入import "github.com/appleboy/gin-jwt"

//User 注入IUserService
type User struct {
	Log     logger.ILogger       `inject:""`
	Service service.IUserService `inject:""`
}

//Logout 退出登录
func (a *User) Logout(c *gin.Context) {
	//common.RespOk(c, http.StatusOK, codes.SUCCESS)
}

// @Summary 注册用户
// @Description 注册一个用户
// @Tags 用户接口
// @Produce  json
// @Param phone query string true "18612345678"
// @Param password query string true "123456"
// @Param code query string true "1234"
// @Success 200 {string} json "{"code":200,"data":{},"message":"ok"}"
// @Router /v1/api/user/register [post]
func (a *User) Register(c *gin.Context) {

	common.ResSuccess(c, "成功")
}



