package event

import (
	"encoding/json"
	"log"

	"github.com/eztalk/pkg/service"
	"nhooyr.io/websocket"
)

type handler func(*websocket.Conn, *event)

// 鉴权结构体
type auth struct {
	Token   string `json:"token,omitempty"`
	Intents int    `json:"intents,omitempty"`
}

// 鉴权事件
func (m *mux) auth(conn *websocket.Conn, e *event) {
	h := new(heartbeat)
	if err := json.Unmarshal(e.Data, h); err != nil {
		log.Println(err)
		return
	}
	m.heartbeat = h
	write(conn, Identify, auth{
		Token:   m.service.BotToken(),
		Intents: 0 | 1<<30 | 1<<1,
	})
}

// 开启心跳事件
func (m *mux) ping(c *websocket.Conn, e *event) {
	go func() {
		m.heartbeat.start(c, func() uint64 {
			return m.seq
		})
	}()
}

// 消息创建事件
func (m *mux) message(c *websocket.Conn, e *event) {
	msg := new(service.Message)
	if err := json.Unmarshal(e.Data, msg); err != nil {
		log.Println(err)
		return
	}
	m.service.Serve(msg)
}
