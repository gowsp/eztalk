package pkg

import (
	"testing"

	"github.com/eztalk/pkg/bot"
	"github.com/eztalk/pkg/invoker"
)

func TestBot(t *testing.T) {
	config, err := invoker.ReadConfig("../configs/config.json")
	if err != nil {
		return
	}
	b := bot.New(config)
	b.Start()
}
