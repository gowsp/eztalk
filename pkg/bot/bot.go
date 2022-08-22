package bot

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"time"

	"github.com/eztalk/pkg/cmd"
	"github.com/eztalk/pkg/event"
	"github.com/eztalk/pkg/invoker"
	"github.com/eztalk/pkg/service"
)

//  从配置文件中 创建机器人
func NewByFile(path string) (*Bot, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	config := new(invoker.Config)
	if err := json.NewDecoder(file).Decode(config); err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return newBot(config), nil
}

// 创建机器人
func New(appId, token string) *Bot {
	config := &invoker.Config{
		AppId: appId,
		Token: token,
	}
	return newBot(config)
}

func newBot(config *invoker.Config) *Bot {
	i := invoker.New(config)
	service := service.New(i)
	return &Bot{service: service}
}

// 机器人结构体
type Bot struct {
	service service.Service
}

// 增加命令
func (b *Bot) AddCmd(builder cmd.CmdBuilder) {
	b.service.AddCmd(builder)
}

// 启动机器
func (b *Bot) Start() error {
	event := event.New(b.service)
	for i := 0; i < 3; i++ {
		url, err := b.service.WsUrl()
		if err != nil {
			log.Println(err)
			time.Sleep(3 * time.Second)
			continue
		}
		i = 0
		if err := event.Listen(url); err != nil {
			log.Println(err)
			time.Sleep(3 * time.Second)
			continue
		}
	}
	return errors.New("connect error")
}
