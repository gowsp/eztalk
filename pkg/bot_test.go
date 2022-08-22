package pkg

import (
	"log"
	"testing"

	"github.com/eztalk/pkg/bot"
	"github.com/eztalk/pkg/cmd"
	"github.com/eztalk/pkg/invoker"
)

func TestNewBot(t *testing.T) {
	bot := bot.New("xxx", "xxxxxx")
	bot.Start()
}

func TestCustomBot(t *testing.T) {
	bot, err := bot.NewByFile("../configs/config.json")
	if err != nil {
		log.Println(err)
		return
	}
	bot.AddCmd(func(i *invoker.Invoker) cmd.Cmd { return &customCmd{} })
	bot.Start()
}

type customCmd struct {
	i *invoker.Invoker
}

func (c *customCmd) Id() []string                     { return []string{"hello"} }
func (c *customCmd) Interactive() bool                { return false }
func (c *customCmd) Handle(cmd.Input) (string, error) { return "自定义", nil }
func (c *customCmd) Reply(cmd.Input) string           { return "" }
