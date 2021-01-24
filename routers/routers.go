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
	//"github.com/konger/ckgo/controller/v1/admin"
	"github.com/konger/ckgo/repository"
	"github.com/konger/ckgo/service"
	"net/http"
)

//InitRouter 初始化Router
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(cors.CorsHandler())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	Configure(r)
	return r
}

//Configure 配置router
func Configure(r *gin.Engine) {
	//controller declare
	//var user admin.User
	//inject declare
	//var article admin.Article
	db := datasource.Db{}
	zap := logger.Logger{}
	//Injection
	var injector inject.Graph
	if err := injector.Provide(
		//&inject.Object{Value: &article},
		&inject.Object{Value: &db},
		&inject.Object{Value: &zap},
		&inject.Object{Value: &repository.ArticleRepository{}},
		&inject.Object{Value: &service.ArticleService{}},
		//&inject.Object{Value: &user},
		&inject.Object{Value: &repository.UserRepository{}},
		&inject.Object{Value: &service.UserService{}},
		&inject.Object{Value: &repository.RoleRepository{}},
		&inject.Object{Value: &service.RoleService{}},
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
	//首页
	r.GET("/", func(c *gin.Context) { c.HTML(http.StatusOK, "index.html", nil) })
	apiPrefix:="/v1/adapi"
	g := r.Group(apiPrefix)
	// 登录验证 jwt token 验证 及信息提取
	var notCheckLoginUrlArr []string
	notCheckLoginUrlArr = append(notCheckLoginUrlArr, apiPrefix+"/user/login")
	notCheckLoginUrlArr = append(notCheckLoginUrlArr, apiPrefix+"/user/logout")
	g.Use(privilege.UserAuthMiddleware(
		privilege.AllowPathPrefixSkipper(notCheckLoginUrlArr...),
	))
	// 权限验证
	var notCheckPermissionUrlArr []string
	notCheckPermissionUrlArr = append(notCheckPermissionUrlArr, notCheckLoginUrlArr...)
	notCheckPermissionUrlArr = append(notCheckPermissionUrlArr, apiPrefix+"/menu/menubuttonlist")
	notCheckPermissionUrlArr = append(notCheckPermissionUrlArr, apiPrefix+"/menu/allmenu")
	notCheckPermissionUrlArr = append(notCheckPermissionUrlArr, apiPrefix+"/admins/adminsroleidlist")
	notCheckPermissionUrlArr = append(notCheckPermissionUrlArr, apiPrefix+"/user/info")
	notCheckPermissionUrlArr = append(notCheckPermissionUrlArr, apiPrefix+"/user/editpwd")
	notCheckPermissionUrlArr = append(notCheckPermissionUrlArr, apiPrefix+"/role/rolemenuidlist")
	notCheckPermissionUrlArr = append(notCheckPermissionUrlArr, apiPrefix+"/role/allrole")
	g.Use(privilege.CasbinMiddleware(
		privilege.AllowPathPrefixSkipper(notCheckPermissionUrlArr...),
	))
	//sys
	RegisterRouterSys(g)

}
