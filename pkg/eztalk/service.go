package eztalk

import (
	"fmt"
	"sync"

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
	return &service{
		invoker: i,
	}
}

type service struct {
	invoker *invoker.Invoker
	topic   sync.Map
}

var cmds = map[string]TopicType{
	"/歇后语":   Fable,
	"/猜字谜":   Riddles,
	"/脑经急转弯": BrainTwists,
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
	question := msg.IsRequest()
	if question {
		content = o.GetTopic(msg)
	} else {
		content = o.GetReply(msg)
	}
	resp := new(Message)
	path := fmt.Sprintf("/channels/%s/messages", msg.ChannelID)
	err := o.invoker.Post(path, Reply{Content: content, MsgID: msg.ID}, &resp)
	if err != nil || !question {
		return
	}
	if d, ok := o.topic.LoadAndDelete(msg.ID); ok {
		o.topic.Store(resp.ID, d)
	}
}
func (o *service) GetTopic(msg *Message) string {
	usrMsg := msg.UserInput()
	if question, ok := cmds[usrMsg]; ok {
		d, err := NewTopic(question)
		if err == nil {
			o.topic.Store(msg.ID, d)
			return "客官请听题：" + d.Riddle
		}
		return "服务开小差"
	}
	return "未知"
}
func (o *service) GetReply(msg *Message) string {
	id := msg.ReplyId()
	fmt.Println(id)
	data, ok := o.topic.Load(id)
	if !ok {
		return "服务开小差"
	}
	q := data.(*Topic)
	a, err := q.GetAnswer()
	if err != nil {
		return "服务开小差"
	}
	usrMsg := msg.UserInput()
	if usrMsg == "答案" {
		return a.Answer
	}
	if a.isCorrect(usrMsg) {
		return "客官厉害啊"
	}
	return "客官差点儿就答对了，再加把劲儿"
}
