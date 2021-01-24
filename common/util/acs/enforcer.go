package acs
import (
	"github.com/konger/ckgo/common/datasource"
	"github.com/casbin/casbin"
	"github.com/casbin/gorm-adapter"
)

var Enforcer *casbin.Enforcer

func init() {
	// mysql 适配器
	db := datasource.Db{}
	adapter := gormadapter.NewAdapterByDB(db.DB())
	// 通过mysql适配器新建一个enforcer
	Enforcer = casbin.NewEnforcer("config/keymatch2_model.conf", adapter)
	// 日志记录
	Enforcer.EnableLog(true)
}