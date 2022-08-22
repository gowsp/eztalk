package cmd

import "github.com/eztalk/pkg/invoker"

// 用户输入信息
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
	// 回复消息, true 解除会话关联
	Reply(Input) (string, bool)
}

// 命令构建器
type CmdBuilder func(i *invoker.Invoker) Cmd

// 默认命令
func Default(i *invoker.Invoker) []CmdBuilder {
	return []CmdBuilder{
		newJokeCmd,
		newTopicCmd,
	}
}
