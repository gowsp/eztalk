package pkg

import (
	"testing"

	"github.com/eztalk/pkg/bot"
	"github.com/eztalk/pkg/eztalk"
	"github.com/eztalk/pkg/invoker"
)

func TestAuth(t *testing.T) {
	config := invoker.Config{
		Env: "https://sandbox.api.sgroup.qq.com",
		// Env:    "https://api.sgroup.qq.com",
		AppId:  "102021662",
		Secret: "xxxxxxxxxx",
		Token:  "xxxxxxxxxxxxxxxxxxxxx",
	}
	b := bot.New(&config)
	b.Start()
}

func TestQuest(t *testing.T) {
	data, err := eztalk.NewTopic(eztalk.BrainTwists)
	if err != nil {
		return
	}
	data.GetAnswer()
}
func TestAns(t *testing.T) {
}
