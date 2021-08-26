package page

//User 用户注册信息结构体
type SendData struct {
	FromId uint        `json:"fromid" binding:"required"` // 发送方
	ToId   uint        `json:"toid" binding:"required"`   // 接收方
	Type   string      `json:"type" binding:"required"`   // 发送类型  group:群发  user:私聊
	Msg    interface{} `json:"msg" binding:"required"`    // 内容
}
