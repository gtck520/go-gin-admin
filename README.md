# go-gin-admin
用于实践go gin框架搭建后台，以及项目实战
# 目录结构
├── cmd          程序入口<br>
├── common   通用模块代码<br>
├───├─codes  常量状态码<br>
├───├─datasource  数据库<br>
├───├─helper  助手函数<br>
├───├─logger  日志插件<br>
├───├─middleware  中间件<br>
├───├─setting  全局设置<br>
├───├─util  第三方插件<br>
├───├─validator  验证插件<br>
├── config       配置文件<br>
├── controller API控制器<br>
├── docs         数据库文件以及文档<br>
├── logs     查询日志<br>
├── models     数据表实体<br>
├── page        页面数据返回实体<br>
├── repository 数据访问层<br>
├── resource      静态资源<br>
├── router       路由<br>
├── runtime     应用运行数据<br>
├── service      业务逻辑层<br>
├── vue-admin Vue前端页面代码<br>
├── websocket socket通讯服务<br>
### 基于go module使用
go env -w GO111MODULE=on   
go env -w GOPROXY=https://goproxy.cn,direct
### 基于 air 调试启动
根目录 输入 air
### swagger 文档查看
项目根目录 输入 swag init 生成最新文档
访问 http://127.0.0.1:8000/swagger/index.html
