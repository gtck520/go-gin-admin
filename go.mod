module github.com/konger/ckgo

go 1.15

require (
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
	github.com/smartystreets/goconvey v0.0.0-20190731233626-505e41936337 // indirect
	github.com/ugorji/go v1.2.3 // indirect
	github.com/unknwon/com v1.0.1
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad // indirect
	golang.org/x/sys v0.0.0-20210119212857-b64e53b001e4 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/ini.v1 v1.47.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace (
	github.com/konger/ckgo/cmd => ./cmd
	github.com/konger/ckgo/common => ./common
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
