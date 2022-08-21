package eztalk

import (
	"strings"
	"time"
)

// websocket 连接
type wsUrl struct {
	Url string `json:"url,omitempty"`
}

// 消息
type Message struct {
	Author struct {
		Avatar   string `json:"avatar"`
		Bot      bool   `json:"bot"`
		ID       string `json:"id"`
		Username string `json:"username"`
	} `json:"author"`
	ChannelID string `json:"channel_id"`
	Content   string `json:"content"`
	GuildID   string `json:"guild_id"`
	ID        string `json:"id"`
	Member    struct {
		JoinedAt time.Time `json:"joined_at"`
		Nick     string    `json:"nick"`
		Roles    []string  `json:"roles"`
	} `json:"member"`
	Mentions []struct {
		Avatar   string `json:"avatar"`
		Bot      bool   `json:"bot"`
		ID       string `json:"id"`
		Username string `json:"username"`
	} `json:"mentions"`
	MessageReference struct {
		MessageID string `json:"message_id"`
	} `json:"message_reference"`
	Seq          int       `json:"seq"`
	SeqInChannel string    `json:"seq_in_channel"`
	Timestamp    time.Time `json:"timestamp"`
}

// 回复消息ID
func (m *Message) ReplyId() string {
	return m.MessageReference.MessageID
}

// 是否请求
func (m *Message) IsRequest() bool {
	return m.ReplyId() == ""
}

//用户输入消息
func (m *Message) UserInput() string {
	val := strings.TrimLeft(m.Content, "\u003c@!xxxxxxxxxxxxxx\u003e ")
	return strings.TrimSpace(val)
}

// 消息请求
type Reply struct {
	Content string `json:"content"`
	MsgID   string `json:"msg_id"`
}
