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

	controller "github.com/konger/ckgo/controller/v1/api"
	"github.com/konger/ckgo/repository"
	service "github.com/konger/ckgo/service/v1/api"

	//"github.com/konger/ckgo/controller/v1/admin"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

//InitRouter 初始化Router
func InitRouter() *gin.Engine {
	log_name := filepath.Join(setting.RunPath, "runtime", "debug.log")
	f, _ := os.Create(log_name)
	//gin.DefaultWriter=io.MultiWriter(f)
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
	//启动websocket
	// go websocket.WebsocketManager.Start()
	// go websocket.WebsocketManager.SendService()
	// go websocket.WebsocketManager.SendService()
	// go websocket.WebsocketManager.SendGroupService()
	// go websocket.WebsocketManager.SendGroupService()
	// go websocket.WebsocketManager.SendAllService()
	// go websocket.WebsocketManager.SendAllService()
	// go websocket.TestSendGroup()
	// go websocket.TestSendAll()
	//controller declare
	var user controller.User
	//inject declare

	db := datasource.Db{}
	zap := logger.Logger{}
	//Injection
	var injector inject.Graph
	if err := injector.Provide(
		&inject.Object{Value: &db},
		&inject.Object{Value: &zap},
		&inject.Object{Value: &user},
		&inject.Object{Value: &repository.UserRepository{}},
		&inject.Object{Value: &service.UserService{}},
		&inject.Object{Value: &repository.BaseRepository{}},
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
	//首页
	r.GET("/", func(c *gin.Context) { c.HTML(http.StatusOK, "index.html", nil) })
	//加载静态资源
	r.StaticFS("/resource", http.Dir("./resource"))
	//swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	adminApiPrefix := "/v1/adapi"
	g := r.Group(adminApiPrefix)

	wsGroup := r.Group("/ws")
	{
		wsGroup.GET("/:channel", websocket.WebsocketManager.WsClient)
	}
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
	//sys
	RegisterRouterSys(g)

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
	RegisterRouterApi(ag)

}
