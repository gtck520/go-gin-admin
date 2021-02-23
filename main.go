package main

import (
	"fmt"
	"net/http"

	"github.com/konger/ckgo/common/setting"
	"github.com/konger/ckgo/routers"
)

// @title Go Gin Admin
// @version 1.0
// @description Gin 聊天项目
// @contact.name konger
// @contact.url https://www.kanglan.vip
// @contact.email 496317580@qq.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
func main() {
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()

}
