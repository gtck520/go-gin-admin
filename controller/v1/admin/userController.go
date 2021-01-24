package admin
//
//import (
//	//"net/http"
//	"strconv"
//
//	jwt "github.com/appleboy/gin-jwt/v2"
//	"github.com/konger/ckgo/common/codes"
//	"github.com/konger/ckgo/common/logger"
//	"github.com/konger/ckgo/controller/common"
//	"github.com/konger/ckgo/models"
//	"github.com/konger/ckgo/page"
//	"github.com/konger/ckgo/service"
//	"github.com/gin-gonic/gin"
//)
//
////### 如果是使用Go Module,gin-jwt模块应使用v2
////下载安装，开启Go Module "go env -w GO111MODULE=on",然后执行"go get github.com/appleboy/gin-jwt/v2"
////导入应写成 import "github.com/appleboy/gin-jwt/v2"
////### 如果不是使用Go Module
////下载安装gin-jwt，"go get github.com/appleboy/gin-jwt"
////导入import "github.com/appleboy/gin-jwt"
//
////User 注入IUserService
//type User struct {
//	Log     logger.ILogger       `inject:""`
//	Service service.IUserService `inject:""`
//}
//
////GetUserInfo 根据token获取用户信息
//func (a *User) GetUserInfo(c *gin.Context) {
//	roles := jwt.ExtractClaims(c)
//	userName := roles["userName"].(string)
//	avatar := a.Service.GetUserAvatar(userName)
//	code := codes.SUCCESS
//	userRoles := a.Service.GetRoles(userName)
//	data := page.User{Roles: userRoles, Introduction: "", Avatar: *avatar, Name: userName}
//	//common.RespData(c, http.StatusOK, code, &data)
//}
//
////Logout 退出登录
//func (a *User) Logout(c *gin.Context) {
//	//common.RespOk(c, http.StatusOK, codes.SUCCESS)
//}
//
////GetUsers 获取用户信息
//func (a *User) GetUsers(c *gin.Context) {
//	var maps string
//	code := codes.SUCCESS
//	name := c.Query("name")
//	if name != "" {
//		maps = "username LIKE '%" + name + "%'"
//	}
//	page, pagesize := common.GetPage(c)
//	data := a.Service.GetUsers(page, pagesize, maps)
//	//common.RespData(c, http.StatusOK, code, data)
//}
//
////AddUser 新建用户
//func (a *User) AddUser(c *gin.Context) {
//	user := models.User{}
//	code := codes.INVALID_PARAMS
//	err := c.Bind(&user)
//	if err != nil {
//		a.Log.Error(err)
//	} else {
//		roles := jwt.ExtractClaims(c)
//		createdBy := roles["userName"].(string)
//		user.CreatedBy = createdBy
//		user.State = 1
//		user.Avatar = "https://zbj-bucket1.oss-cn-shenzhen.aliyuncs.com/avatar.JPG"
//		if !a.Service.ExistUserByName(user.Username) {
//			if a.Service.AddUser(&user) {
//				code = codes.SUCCESS
//			} else {
//				code = codes.ERROR
//			}
//		} else {
//			code = codes.ERROR_EXISTS_USER
//		}
//	}
//
//	//common.RespOk(c, http.StatusOK, code)
//}
//
////UpdateUser 修改用户
//func (a *User) UpdateUser(c *gin.Context) {
//	user := models.User{}
//	code := codes.INVALID_PARAMS
//	err := c.Bind(&user)
//	if err != nil {
//		a.Log.Error(err)
//	} else {
//		roles := jwt.ExtractClaims(c)
//		modifiedBy := roles["userName"].(string)
//		user.ModifiedBy = modifiedBy
//		if a.Service.UpdateUser(&user) {
//			code = codes.SUCCESS
//		} else {
//			code = codes.ERROR
//		}
//	}
//	//common.RespOk(c, http.StatusOK, code)
//}
//
////DeleteUser 删除用户
//func (a *User) DeleteUser(c *gin.Context) {
//	id, _ := strconv.Atoi(c.Param("id"))
//	code := codes.SUCCESS
//	if !a.Service.DeleteUser(id) {
//		code = codes.ERROR
//		//common.RespFail(c, http.StatusOK, code, "不允许删除admin账号!")
//	} else {
//		//common.RespOk(c, http.StatusOK, code)
//	}
//}
