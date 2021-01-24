package privilege
import (
	"github.com/konger/ckgo/common/util/acs"
	"github.com/konger/ckgo/common/util/response"
	"github.com/konger/ckgo/common/util/cache"
	"github.com/gin-gonic/gin"
	"log"
)

func Privilege() gin.HandlerFunc {
	return func(c *gin.Context) {

		var userName = c.GetHeader("userName")
		if userName == "" {
			response.RespFailMsg(c,"header miss userName")
			c.Abort()
			return
		}
		path := c.Request.URL.Path
		method := c.Request.Method
		cacheName := userName + path + method
		// 从缓存中读取&判断
		entry, err := cache.GlobalCache.Get(cacheName)
		if err == nil && entry != nil {
			if string(entry) == "true" {
				c.Next()
			} else {
				response.RespFailMsg(c,"access denied")
				c.Abort()
				return
			}
		} else {
			// 从数据库中读取&判断
			//记录日志
			acs.Enforcer.EnableLog(true)
			// 加载策略规则
			err := acs.Enforcer.LoadPolicy()
			if err != nil {
				log.Println("loadPolicy error")
				panic(err)
			}
			// 验证策略规则
			result, err := acs.Enforcer.EnforceSafe(userName, path, method)
			if err != nil {
				response.RespFailMsg(c,"No permission found")
				c.Abort()
				return
			}
			if !result {
				// 添加到缓存中
				cache.GlobalCache.Set(cacheName, []byte("false"))
				response.RespFailMsg(c,"access denied")
				c.Abort()
				return
			} else {
				cache.GlobalCache.Set(cacheName, []byte("true"))
			}
			c.Next()
		}
	}
}
