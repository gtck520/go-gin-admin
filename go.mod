module github.com/konger/ckgo

go 1.15

require (
	github.com/ahmetb/go-linq v3.0.0+incompatible
	github.com/casbin/casbin v1.9.1
	github.com/coocood/freecache v1.1.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/facebookgo/ensure v0.0.0-20200202191622-63f1cf65ac4c // indirect
	github.com/facebookgo/inject v0.0.0-20180706035515-f23751cae28b
	github.com/facebookgo/stack v0.0.0-20160209184415-751773369052 // indirect
	github.com/facebookgo/structtag v0.0.0-20150214074306-217e25fb9691 // indirect
	github.com/facebookgo/subset v0.0.0-20200203212716-c811ad88dec4 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.62.0
	github.com/go-openapi/spec v0.20.3 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/go-playground/validator/v10 v10.4.1
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/gopherjs/gopherjs v0.0.0-20181103185306-d547d1d9531e // indirect
	github.com/gorilla/websocket v1.4.2
	github.com/jinzhu/gorm v1.9.16
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/lib/pq v1.2.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-sqlite3 v2.0.3+incompatible // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/satori/go.uuid v1.2.0
	github.com/sirupsen/logrus v1.7.0
	github.com/smartystreets/assertions v0.0.0-20190116191733-b6c0e53d7304 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/stretchr/testify v1.6.1
	github.com/swaggo/gin-swagger v1.3.0
	github.com/swaggo/swag v1.7.0 // indirect
	github.com/ugorji/go v1.2.4 // indirect
	go.uber.org/zap v1.16.0
	golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83 // indirect
	golang.org/x/net v0.0.0-20210510120150-4163338589ed // indirect
	golang.org/x/sys v0.0.0-20210514084401-e8d321eab015 // indirect
	golang.org/x/tools v0.1.1 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/ini.v1 v1.51.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace (
	github.com/konger/ckgo/cmd => ./cmd
	github.com/konger/ckgo/common => ./common
	github.com/konger/ckgo/common/datasource => ./common/datasource
	github.com/konger/ckgo/common/helper => ./common/helper
	github.com/konger/ckgo/common/setting => ./common/setting
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
