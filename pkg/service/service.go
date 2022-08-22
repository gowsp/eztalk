package service

import (
	"bytes"
	"fmt"
	"log"
	"sync"

	"github.com/eztalk/pkg/cmd"
	"github.com/eztalk/pkg/invoker"
)

// 对外服务
type Service interface {
	// 机器人toekn
	BotToken() string
	// 获取ws地址
	WsUrl() (string, error)
	// 监听消息
	Serve(msg *Message)
}

// 创建服务
func New(i *invoker.Invoker) Service {
	service := &service{
		invoker: i,
		command: make(map[string]cmd.Cmd),
	}
	cmds := cmd.Init(i)
	for _, cmd := range cmds {
		service.addCmd(cmd)
	}
	return service
}

type service struct {
	desc    bytes.Buffer
	session sync.Map
	command map[string]cmd.Cmd
	invoker *invoker.Invoker
}

func (s *service) addCmd(cmd cmd.Cmd) {
	for _, id := range cmd.Id() {
		s.command[id] = cmd
		s.desc.WriteString(id + "\n")
	}
}
func (s *service) BotToken() string {
	return s.invoker.BotToken()
}
func (s *service) WsUrl() (string, error) {
	wsUrl := new(wsUrl)
	if err := s.invoker.Get("/gateway", wsUrl); err != nil {
		return "", err
	}
	return wsUrl.Url, nil
}
func (o *service) Serve(msg *Message) {
	var content string
	if val, ok := o.session.Load(msg.UserId()); ok {
		content = val.(cmd.Cmd).Reply(msg)
	} else {
		content = o.Greet(msg)
	}
	resp := new(Message)
	content = msg.Author.Mention() + content
	path := fmt.Sprintf("/channels/%s/messages", msg.ChannelID)
	err := o.invoker.Post(path, Reply{MsgID: msg.ID, Content: content}, &resp)
	if err != nil {
		log.Println(err)
	}
}
func (s *service) Greet(msg *Message) string {
	request := msg.UserInput()
	if cmd, ok := s.command[request]; ok {
		d, err := cmd.Handle(msg)
		if err != nil {
			log.Println(err)
			return "服务开小差了"
		}
		// 交互命令 访问后需存储会话
		if cmd.Interactive() {
			s.session.Store(msg.UserId(), cmd)
		}
		return d
	}
	return "客官来点儿什么，本店有\n\n" + s.desc.String()
}
