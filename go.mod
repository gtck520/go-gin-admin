module github.com/konger/ckgo

go 1.15

require (
	github.com/appleboy/gin-jwt/v2 v2.6.4
	github.com/astaxie/beego v1.12.3
	github.com/facebookgo/inject v0.0.0-20180706035515-f23751cae28b
	github.com/facebookgo/structtag v0.0.0-20150214074306-217e25fb9691 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.62.0
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/lib/pq v1.2.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	github.com/smartystreets/goconvey v0.0.0-20190731233626-505e41936337 // indirect
	github.com/ugorji/go v1.2.3 // indirect
	github.com/unknwon/com v1.0.1
	go.uber.org/zap v1.16.0
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad // indirect
	golang.org/x/sys v0.0.0-20210119212857-b64e53b001e4 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/ini.v1 v1.47.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace (
	github.com/konger/ckgo/cmd => ./cmd
	github.com/konger/ckgo/common => ./common
	github.com/konger/ckgo/common/setting => ./common/setting
	github.com/konger/ckgo/common/helper => ./common/helper
	github.com/konger/ckgo/common/datasource => ./common/datasource
	github.com/konger/ckgo/config => ./config
	github.com/konger/ckgo/controller => ./controller
	github.com/konger/ckgo/docs => ./docs
	github.com/konger/ckgo/models => ./models
	github.com/konger/ckgo/page => ./page
	github.com/konger/ckgo/repository => ./repository
	github.com/konger/ckgo/routers => ./routers
	github.com/konger/ckgo/runtime => ./runtime
	github.com/konger/ckgo/service => ./service
)
