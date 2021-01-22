# go-gin-admin
用于实践go gin框架搭建后台，以及项目实战
# 目录结构
├── cmd          程序入口<br>
├── common   通用模块代码<br>
├── config       配置文件<br>
├── controller API控制器<br>
├── docs         数据库文件以及文档<br>
├── middleware    中间件<br>
├── models     数据表实体<br>
├── page        页面数据返回实体<br>
├── pkg          第三方包<br>
├── repository 数据访问层<br>
├── router       路由<br>
├── runtime     应用运行数据<br>
├── service      业务逻辑层<br>
├── vue-admin Vue前端页面代码<br>
### 基于go module使用
go env -w GO111MODULE=on   
go env -w GOPROXY=https://goproxy.cn,direct
