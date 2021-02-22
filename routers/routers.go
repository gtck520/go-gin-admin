package routers

import (
	"log"

	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"
	"github.com/konger/ckgo/controller/common"
	"github.com/konger/ckgo/common/datasource"
	"github.com/konger/ckgo/common/logger"
	"github.com/konger/ckgo/common/middleware/cors"
	"github.com/konger/ckgo/common/middleware/privilege"
	"github.com/konger/ckgo/common/setting"


	//"github.com/konger/ckgo/controller/v1/admin"
	"net/http"
)

//InitRouter 初始化Router
func InitRouter() *gin.Engine {
	r := gin.New()
	if err := logger.InitLogger(); err != nil {
		log.Fatal("init logger failed:", err)
	}
	r.Use(logger.GinLogger())
	r.Use(logger.GinRecovery(true))
	r.Use(cors.CorsHandler())
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
	//Injection
	var injector inject.Graph
	if err := injector.Provide(
		//&inject.Object{Value: &article},
		&inject.Object{Value: &db},
		//&inject.Object{Value: &repository.ArticleRepository{}},
		//&inject.Object{Value: &service.ArticleService{}},
		//&inject.Object{Value: &user},
		// &inject.Object{Value: &repository.UserRepository{}},
		// &inject.Object{Value: &service.UserService{}},
		// &inject.Object{Value: &repository.RoleRepository{}},
		// &inject.Object{Value: &service.RoleService{}},
		// &inject.Object{Value: &repository.BaseRepository{}},
	); err != nil {
		log.Fatal("inject fatal: ", err)
	}
	if err := injector.Populate(); err != nil {
		log.Fatal("injector fatal: ", err)
	}
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
	apiPrefix := "/v1/adapi"
	g := r.Group(apiPrefix)
	// 登录验证 jwt token 验证 及信息提取
	var notCheckLoginURLArr []string
	notCheckLoginURLArr = append(notCheckLoginURLArr, apiPrefix+"/user/login")
	notCheckLoginURLArr = append(notCheckLoginURLArr, apiPrefix+"/user/logout")
	g.Use(privilege.UserAuthMiddleware(
		privilege.AllowPathPrefixSkipper(notCheckLoginURLArr...),
	))
	// 权限验证
	var notCheckPermissionURLArr []string
	notCheckPermissionURLArr = append(notCheckPermissionURLArr, notCheckLoginURLArr...)
	notCheckPermissionURLArr = append(notCheckPermissionURLArr, apiPrefix+"/menu/menubuttonlist")
	notCheckPermissionURLArr = append(notCheckPermissionURLArr, apiPrefix+"/menu/allmenu")
	notCheckPermissionURLArr = append(notCheckPermissionURLArr, apiPrefix+"/admins/adminsroleidlist")
	notCheckPermissionURLArr = append(notCheckPermissionURLArr, apiPrefix+"/user/info")
	notCheckPermissionURLArr = append(notCheckPermissionURLArr, apiPrefix+"/user/editpwd")
	notCheckPermissionURLArr = append(notCheckPermissionURLArr, apiPrefix+"/role/rolemenuidlist")
	notCheckPermissionURLArr = append(notCheckPermissionURLArr, apiPrefix+"/role/allrole")
	g.Use(privilege.CasbinMiddleware(
		privilege.AllowPathPrefixSkipper(notCheckPermissionURLArr...),
	))
	//sys
	RegisterRouterSys(g)

}
