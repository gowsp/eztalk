package bot

import (
	"errors"
	"log"
	"time"

	"github.com/eztalk/pkg/eztalk"
	"github.com/eztalk/pkg/invoker"
	"github.com/eztalk/pkg/wss"
)

func New(config *invoker.Config) *Bot {
	i := invoker.New(config)
	return &Bot{invoker: i}
}

type Bot struct {
	invoker *invoker.Invoker
}

func (b *Bot) Start() error {
	service := eztalk.New(b.invoker)
	w := wss.New(service)
	for i := 0; i < 3; i++ {
		url, err := service.WsUrl()
		if err != nil {
			log.Println(err)
			time.Sleep(3 * time.Second)
			continue
		}
		i = 0
		if err := w.Connect(url); err != nil {
			log.Println(err)
			time.Sleep(3 * time.Second)
			continue
		}
	}
	return errors.New("connect error")
}
