package routers

import (
	"log"

	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"

	"github.com/konger/ckgo/common/datasource"
	"github.com/konger/ckgo/common/logger"
	"github.com/konger/ckgo/common/middleware/cors"
	"github.com/konger/ckgo/common/middleware/jwt"
	"github.com/konger/ckgo/common/setting"
	"github.com/konger/ckgo/controller/v1/admin"
	"github.com/konger/ckgo/repository"
	"github.com/konger/ckgo/service"
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
	var user admin.User
	//var tag cv1.Tag
	var myjwt jwt.JWT
	//inject declare
	var article admin.Article
	db := datasource.Db{}
	zap := logger.Logger{}
	//Injection
	var injector inject.Graph
	if err := injector.Provide(
		&inject.Object{Value: &article},
		&inject.Object{Value: &db},
		&inject.Object{Value: &zap},
		&inject.Object{Value: &repository.ArticleRepository{}},
		&inject.Object{Value: &service.ArticleService{}},
		&inject.Object{Value: &user},
		&inject.Object{Value: &repository.UserRepository{}},
		&inject.Object{Value: &service.UserService{}},
		&inject.Object{Value: &repository.RoleRepository{}},
		&inject.Object{Value: &service.RoleService{}},
		&inject.Object{Value: &myjwt},
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
	var authMiddleware = myjwt.GinJWTMiddlewareInit(&jwt.AllUserAuthorizator{})
	r.NoRoute(authMiddleware.MiddlewareFunc(), jwt.NoRouteHandler)
	r.POST("/login", authMiddleware.LoginHandler)
	userAPI := r.Group("/user")
	{
		userAPI.GET("/refresh_token", authMiddleware.RefreshHandler)
	}
	userAPI.Use(authMiddleware.MiddlewareFunc())
	{
		userAPI.GET("/table/list", article.GetTables)
		userAPI.GET("/info", user.GetUserInfo)
		userAPI.POST("/logout", user.Logout)
	}

	var adminMiddleware = myjwt.GinJWTMiddlewareInit(&jwt.AdminAuthorizator{})
	apiv1 := r.Group("/admin/v1")
	//使用AdminAuthorizator中间件，只有admin权限的用户才能获取到接口
	apiv1.Use(adminMiddleware.MiddlewareFunc())
	{
		//vue获取table信息
		//apiv1.GET("/table/list", article.GetTables)
		apiv1.GET("/user/list", user.GetUsers)
		apiv1.POST("/user", user.AddUser)
		apiv1.PUT("/user", user.UpdateUser)
		apiv1.DELETE("/user/:id", user.DeleteUser)
		apiv1.GET("/article/list", article.GetArticles)
		apiv1.GET("/article/detail/:id", article.GetArticle)
		apiv1.POST("/article", article.AddArticle)
		// apiv1.PUT("/articles/:id", article.EditArticle)
		// apiv1.DELETE("/articles/:id", article.DeleteArticle)
	}
}
