package cmd

import "github.com/eztalk/pkg/invoker"

type Input interface {
	UserId() string
	UserInput() string
}

// 交互指令
type Cmd interface {
	// 命令ID
	Id() []string
	// 交互性，true 处理会话，fasle 不处理会话
	Interactive() bool
	// 处理消息
	Handle(Input) (string, error)
	// 回复消息
	Reply(Input) string
}

// 初始化交互命令
func Init(i *invoker.Invoker) []Cmd {
	return []Cmd{
		&jokeCmd{i},
		&topicCmd{invoker: i},
	}
}
