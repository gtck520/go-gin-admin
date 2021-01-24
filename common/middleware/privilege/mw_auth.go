package privilege

import (
	"strconv"
	"time"

	"github.com/konger/ckgo/common/codes"
	"github.com/konger/ckgo/common/middleware/jwt"
	"github.com/konger/ckgo/controller/common"
	"github.com/konger/ckgo/common/util/cache"
	"github.com/konger/ckgo/common/util/convert"

	"github.com/gin-gonic/gin"
)

// UserAuthMiddleware 用户授权中间件
func UserAuthMiddleware(skipper ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(skipper) > 0 && skipper[0](c) {
			c.Next()
			return
		}
		var uuid string
		if t := c.GetHeader(codes.TOKEN_KEY); t != "" {
			userInfo,ok:=jwt.ParseToken(t)
			if !ok {
					common.ResFailCode(c,"token 无效",50008)
			    return
			}
			exptimestamp, _ := strconv.ParseInt(userInfo["exp"], 10, 64)
      exp := time.Unix(exptimestamp, 0)
			ok=exp.After(time.Now())
			if !ok {
				common.ResFailCode(c,"token 过期",50014)
				return
			}
			uuid=userInfo["uuid"]
		}

		if uuid != "" {
			//查询用户ID
			val,err:=cache.Get([]byte(uuid))
			if err!=nil {
				common.ResFailCode(c,"token 无效",50008)
				return
			}
			userID:=convert.ToUint(string(val))
			c.Set(codes.USER_UUID_Key, uuid)
			c.Set(codes.USER_ID_Key, userID)
		}
		if uuid == "" {
			common.ResFailCode(c,"用户未登录",50008)
			return
		}
	}
}
