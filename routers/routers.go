package routers

import (
	"log"

	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"
	"github.com/konger/ckgo/common/datasource"
	"github.com/konger/ckgo/common/logger"
	"github.com/konger/ckgo/common/middleware/cors"
	"github.com/konger/ckgo/common/middleware/privilege"
	"github.com/konger/ckgo/common/setting"
	"github.com/konger/ckgo/controller/common"

	"github.com/konger/ckgo/websocket"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/konger/ckgo/controller/sys"
	controller "github.com/konger/ckgo/controller/v1/api"

	//service "github.com/konger/ckgo/service/v1/api"

	"io"
	"net/http"
	"os"
	"path/filepath"
)

//InitRouter 初始化Router
func InitRouter() *gin.Engine {
	log_name := filepath.Join(setting.RunPath, "runtime", "shell.log")
	f, _ := os.Create(log_name)
	//gin.DefaultWriter = io.MultiWriter(f)
	// 如果你需要同时写入日志文件和控制台上显示，使用下面代码
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r := gin.New()
	//if err := logger.InitLogger(); err != nil {
	//	log.Fatal("init logger failed:", err)
	//}
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.CorsHandler())
	gin.SetMode(setting.RunMode)
	Configure(r)
	return r
}

//Configure 配置router
func Configure(r *gin.Engine) {

	//controller declare
	var user controller.User
	var friend controller.Friend
	var ws_controller controller.Ws

	//inject declare
	ServerManager := websocket.SocketServer
	db := datasource.Db{}
	zap := logger.Logger{}
	//Injection
	injector := inject.Graph{}
	if err := injector.Provide(
		&inject.Object{Value: &db},
		&inject.Object{Value: &zap},
		&inject.Object{Value: &user},
		&inject.Object{Value: &friend},
		&inject.Object{Value: &ServerManager},
		&inject.Object{Value: &ws_controller},
	); err != nil {
		log.Fatal("inject fatal: ", err)
	}
	if err := injector.Populate(); err != nil {
		log.Fatal("injector fatal: ", err)
	}

	//zap log init
	zap.Init()
	//database connect
	if err := db.Connect(); err != nil {
		log.Fatal("db fatal:", err)
	}
	//初始化数据库
	datasource.Migration()
	//初始化casbin
	common.InitCsbinEnforcer()

	//websocket广播测试
	go ServerManager.TestChatclient()
	// 开启服务端
	go ServerManager.OnMessage()

	//首页
	r.GET("/", func(c *gin.Context) { c.HTML(http.StatusOK, "index.html", nil) })
	//加载静态资源
	r.StaticFS("/resource", http.Dir("./resource"))
	//swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	adminApiPrefix := "/v1/adapi"
	g := r.Group(adminApiPrefix)
	// 登录验证 jwt token 验证 及信息提取
	var notCheckLoginURLArr []string
	notCheckLoginURLArr = append(notCheckLoginURLArr, adminApiPrefix+"/user/login")
	notCheckLoginURLArr = append(notCheckLoginURLArr, adminApiPrefix+"/user/logout")
	g.Use(privilege.UserAuthMiddleware(
		privilege.AllowPathPrefixSkipper(notCheckLoginURLArr...),
	))
	// 权限验证
	var notCheckPermissionURLArr []string
	notCheckPermissionURLArr = append(notCheckPermissionURLArr, notCheckLoginURLArr...)
	notCheckPermissionURLArr = append(notCheckPermissionURLArr, adminApiPrefix+"/menu/menubuttonlist")
	notCheckPermissionURLArr = append(notCheckPermissionURLArr, adminApiPrefix+"/menu/allmenu")
	notCheckPermissionURLArr = append(notCheckPermissionURLArr, adminApiPrefix+"/admins/adminsroleidlist")
	notCheckPermissionURLArr = append(notCheckPermissionURLArr, adminApiPrefix+"/user/info")
	notCheckPermissionURLArr = append(notCheckPermissionURLArr, adminApiPrefix+"/user/editpwd")
	notCheckPermissionURLArr = append(notCheckPermissionURLArr, adminApiPrefix+"/role/rolemenuidlist")
	notCheckPermissionURLArr = append(notCheckPermissionURLArr, adminApiPrefix+"/role/allrole")
	g.Use(privilege.CasbinMiddleware(
		privilege.AllowPathPrefixSkipper(notCheckPermissionURLArr...),
	))
	//系统后台接口
	menu := sys.Menu{}
	g.GET("/menu/list", menu.List)
	g.GET("/menu/detail", menu.Detail)
	g.GET("/menu/allmenu", menu.AllMenu)
	g.GET("/menu/menubuttonlist", menu.MenuButtonList)
	g.POST("/menu/delete", menu.Delete)
	g.POST("/menu/update", menu.Update)
	g.POST("/menu/create", menu.Create)
	auser := sys.User{}
	g.GET("/user/info", auser.Info)
	g.POST("/user/login", auser.Login)
	g.POST("/user/logout", auser.Logout)
	g.POST("/user/editpwd", auser.EditPwd)
	admins := sys.Admins{}
	g.GET("/admins/list", admins.List)
	g.GET("/admins/detail", admins.Detail)
	g.GET("/admins/adminsroleidlist", admins.AdminsRoleIDList)
	g.POST("/admins/delete", admins.Delete)
	g.POST("/admins/update", admins.Update)
	g.POST("/admins/create", admins.Create)
	g.POST("/admins/setrole", admins.SetRole)
	role := sys.Role{}
	g.GET("/role/list", role.List)
	g.GET("/role/detail", role.Detail)
	g.GET("/role/rolemenuidlist", role.RoleMenuIDList)
	g.GET("/role/allrole", role.AllRole)
	g.POST("/role/delete", role.Delete)
	g.POST("/role/update", role.Update)
	g.POST("/role/create", role.Create)
	g.POST("/role/setrole", role.SetRole)

	apiPrefix := "/v1/api"
	ag := r.Group(apiPrefix)
	// 登录验证 jwt token 验证 及信息提取
	var notCheckLoginURLArrApi []string
	notCheckLoginURLArrApi = append(notCheckLoginURLArrApi, apiPrefix+"/user/login")
	notCheckLoginURLArrApi = append(notCheckLoginURLArrApi, apiPrefix+"/user/logout")
	notCheckLoginURLArrApi = append(notCheckLoginURLArrApi, apiPrefix+"/user/register")
	ag.Use(privilege.UserAuthMiddleware(
		privilege.AllowPathPrefixSkipper(notCheckLoginURLArrApi...),
	))

	//api接口
	ag.POST("/user/register", user.Register)
	ag.POST("/user/login", user.Login)
	ag.POST("/user/logout", user.Logout)
	ag.POST("/user/info", user.Info)
	ag.GET("/friend/friend_list", friend.FriendList)
	ag.POST("/ws/send", ws_controller.SendMessage)

	//websocket
	ws := r.Group("/ws")
	ws.Use(privilege.UserAuthMiddleware())
	ws.GET("/chat", ServerManager.NewChatServer)

}
