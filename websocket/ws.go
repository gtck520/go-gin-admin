package websocket

import (
	"encoding/json"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/konger/ckgo/common/codes"
	"github.com/konger/ckgo/common/logger"
	"github.com/konger/ckgo/common/util/convert"
	service "github.com/konger/ckgo/service/v1/api"

	"log"
	"net/http"
)

//服务管理
type ServerManager struct {
	Connlist      map[string]*User //所有连接通道
	register      chan *User       //注册客户端消息通道
	ClientMessage chan *Message    //发送客户端消息通道
	ServerMessage chan *Message    //发送服务端消息通道
	SysBroadcast  chan *Message    //系统广播通道
	Len           int              //长度
	Lock          sync.Mutex
	UserService   *service.UserService `inject:""`
	Log           logger.ILogger       `inject:""`
}
type User struct {
	Conn     *websocket.Conn
	Name     string
	Id       string
	Avator   string
	To_id    string
	group_id string
}

// 消息载体
type Message struct {
	conn        *websocket.Conn
	context     *gin.Context
	content     []byte
	messageType int
}

//发送内容content
type TypeMessage struct {
	Type interface{} `json:"type"`
	Data interface{} `json:"data"`
}

//TypeMessage Data
type ClientMessage struct {
	Name      string `json:"name"`
	Avator    string `json:"avator"`
	Id        string `json:"id"`
	VisitorId string `json:"visitor_id"`
	Group     string `json:"group"`
	Time      string `json:"time"`
	ToId      string `json:"to_id"`
	Content   string `json:"content"`
	City      string `json:"city"`
	ClientIp  string `json:"client_ip"`
	Refer     string `json:"refer"`
}

//客户端连接列表
var upgrader = websocket.Upgrader{}
var SocketServer ServerManager

func init() {
	//初始化socket参数
	SocketServer = ServerManager{
		Connlist:      make(map[string]*User),
		register:      make(chan *User, 128),
		ClientMessage: make(chan *Message, 128),
		ServerMessage: make(chan *Message, 128),
		SysBroadcast:  make(chan *Message, 128),
		Len:           0,
	}
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		// 解决跨域问题
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
}
func (s *ServerManager) NewChatServer(c *gin.Context) {
	UserId, ok := c.Get(codes.USER_ID_Key)
	log.Println("userid:", UserId)
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	//客户端连接成功
	log.Print("客户端连接成功：" + conn.RemoteAddr().String())

	if !ok {
		content, _ := json.Marshal("用户id不存在，无法绑定")
		s.ClientMessage <- &Message{
			conn:        conn,
			content:     content,
			context:     c,
			messageType: 1,
		}
	}
	//用户数据
	UserInfo := s.UserService.Repository.GetUserByID(UserId.(uint))
	user := &User{
		Conn:     conn,
		Name:     UserInfo.Phone,
		Id:       convert.ToString(UserId),
		Avator:   "string",
		To_id:    "string",
		group_id: "string",
	}
	s.register <- user

	for {
		//接受消息
		var receive []byte
		var recevString string
		messageType, receive, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		recevString = string(receive)
		log.Println("客户端:", recevString)

		var typeMsg TypeMessage
		json.Unmarshal(receive, &typeMsg)
		if typeMsg.Type == nil || typeMsg.Data == nil {
			log.Println("消息格式错误：", recevString)
			content, _ := json.Marshal("消息格式错误")
			s.ClientMessage <- &Message{
				conn:        conn,
				content:     content,
				context:     c,
				messageType: messageType,
			}
			continue
		}

		s.ServerMessage <- &Message{
			conn:        conn,
			content:     receive,
			context:     c,
			messageType: messageType,
		}
	}
}
func (s *ServerManager) OnMessage() {
	for {
		select {
		case message := <-s.ClientMessage:
			log.Println("服务端:", string(message.content))
			s.Lock.Lock()
			message.conn.WriteMessage(websocket.TextMessage, message.content)
			s.Lock.Unlock()
		case message := <-s.ServerMessage:
			var typeMsg TypeMessage
			json.Unmarshal(message.content, &typeMsg)
			conn := message.conn
			msgType := typeMsg.Type.(string)
			log.Println("客户端:", string(message.content))

			switch msgType {
			//心跳
			case "ping":
				msg := TypeMessage{
					Type: "pong",
				}
				str, _ := json.Marshal(msg)
				s.Lock.Lock()
				conn.WriteMessage(websocket.TextMessage, str)
				s.Lock.Unlock()
			}
		case user := <-s.register:
			log.Println("注册用户:", user)
			s.Len++
			s.Connlist[strconv.Itoa(s.Len)] = user
		}

	}

}
func (s *ServerManager) TestChatclient() {
	for {
		time.Sleep(time.Second * 25)
		for key, otherKefu := range s.Connlist {
			str, _ := json.Marshal("发送给：" + key + "；" + time.Now().Format("2006-01-02 15:04:05"))

			err := otherKefu.Conn.WriteMessage(websocket.TextMessage, str)
			if err == nil {

			}
		}
		log.Println("发送")
	}
}
