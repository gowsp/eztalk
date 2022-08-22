package cmd

import (
	"fmt"
	"net/url"

	"github.com/eztalk/pkg/invoker"
)

// 讲笑话命令
type jokeCmd struct {
	invoker *invoker.Invoker
}

// 命令ID
func (c *jokeCmd) Id() []string {
	return []string{"/讲笑话"}
}

// 交互性，true 处理会话，fasle 不处理会话
func (c *jokeCmd) Interactive() bool {
	return false
}

// 笑话消息体
type joke struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// 格式化输出
func (j *joke) String() string {
	return fmt.Sprintf("%s\n%s", j.Title, j.Content)
}

// 处理消息
func (c *jokeCmd) Handle(input Input) (string, error) {
	joke := new(joke)
	if err := c.invoker.PostForm("/xiaohua", make(url.Values), joke); err != nil {
		return "", err
	}
	return joke.String(), nil
}

// 回复消息
func (c *jokeCmd) Reply(input Input) string {
	return ""
}
