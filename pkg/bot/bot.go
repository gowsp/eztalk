package bot

import (
	"errors"
	"log"
	"time"

	"github.com/eztalk/pkg/event"
	"github.com/eztalk/pkg/invoker"
	"github.com/eztalk/pkg/service"
)

func New(config *invoker.Config) *Bot {
	i := invoker.New(config)
	return &Bot{invoker: i}
}

// 机器人结构体
type Bot struct {
	invoker *invoker.Invoker
}

// 启动机器人
func (b *Bot) Start() error {
	service := service.New(b.invoker)
	w := event.New(service)
	for i := 0; i < 3; i++ {
		url, err := service.WsUrl()
		if err != nil {
			log.Println(err)
			time.Sleep(3 * time.Second)
			continue
		}
		i = 0
		if err := w.Listen(url); err != nil {
			log.Println(err)
			time.Sleep(3 * time.Second)
			continue
		}
	}
	return errors.New("connect error")
}
