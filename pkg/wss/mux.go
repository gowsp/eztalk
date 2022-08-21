package wss

import (
	"log"

	"github.com/eztalk/pkg/eztalk"
	"nhooyr.io/websocket"
)

type handler func(*websocket.Conn, *event)

var emptyHandler = func(*websocket.Conn, *event) {}

// 创建处理器
func NewMux(service eztalk.Service) *mux {
	m := &mux{
		service: service,
		entry:   make(map[string]handler),
	}
	m.init()
	return m
}

type mux struct {
	seq       uint64
	entry     map[string]handler
	service   eztalk.Service
	heartbeat *heartbeat
}

func (m *mux) init() {
	m.entry[Hello.Event("")] = m.auth
	m.entry[Dispatch.Event("READY")] = m.ping
	m.entry[Heartbeat_ACK.Event("")] = emptyHandler
	m.entry[Dispatch.Event("AT_MESSAGE_CREATE")] = m.message
}

func (m *mux) Serve(conn *websocket.Conn, msg *event) {
	if msg.Seq > m.seq {
		m.seq = msg.Seq
	}
	event := msg.key()
	if handler, ok := m.entry[event]; ok {
		handler(conn, msg)
	} else {
		log.Printf("not support %s\n", event)
	}
}
