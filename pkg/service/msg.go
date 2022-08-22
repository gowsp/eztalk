package service

import (
	"strings"
	"time"
)

// 用户结构体
type User struct {
	ID   string `json:"id"`
	Name string `json:"username"`
	Bot  bool   `json:"bot"`
}

// @前缀
func (u *User) Mention() string {
	return "<@!" + u.ID + ">"
}

// websocket 连接地址信息
type wsUrl struct {
	Url string `json:"url,omitempty"`
}

// 消息结构体
type Message struct {
	Author    User   `json:"author"`
	ChannelID string `json:"channel_id"`
	Content   string `json:"content"`
	GuildID   string `json:"guild_id"`
	ID        string `json:"id"`
	Member    struct {
		JoinedAt time.Time `json:"joined_at"`
		Nick     string    `json:"nick"`
		Roles    []string  `json:"roles"`
	} `json:"member"`
	Mentions         []User `json:"mentions"`
	MessageReference struct {
		MessageID string `json:"message_id"`
	} `json:"message_reference"`
	Seq          int       `json:"seq"`
	SeqInChannel string    `json:"seq_in_channel"`
	Timestamp    time.Time `json:"timestamp"`
}

// 消息创建者ID
func (m *Message) UserId() string {
	return m.Author.ID
}

// 获取用户输入消息
func (m *Message) UserInput() string {
	content := m.Content
	for _, user := range m.Mentions {
		content = strings.ReplaceAll(content, user.Mention(), "")
	}
	return strings.TrimSpace(content)
}

// 回复消息
type Reply struct {
	Content string `json:"content"`
	MsgID   string `json:"msg_id"`
}
