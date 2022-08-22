package event

import (
	"log"

	"github.com/eztalk/pkg/service"
	"nhooyr.io/websocket"
)

var emptyHandler = func(*websocket.Conn, *event) {}

// 创建处理器
func newMux(service service.Service) *mux {
	m := &mux{
		service: service,
		entry:   make(map[string]handler),
	}
	m.init()
	return m
}

// 处理器结构体
type mux struct {
	seq       uint64
	entry     map[string]handler
	service   service.Service
	heartbeat *heartbeat
}

// 初始化事件处理器
func (m *mux) init() {
	m.entry[Hello.Event("")] = m.auth
	m.entry[Dispatch.Event("READY")] = m.ping
	m.entry[Heartbeat_ACK.Event("")] = emptyHandler
	m.entry[Dispatch.Event("AT_MESSAGE_CREATE")] = m.message
}

// 服务收到的消息
func (m *mux) Serve(conn *websocket.Conn, msg *event) {
	if msg.Seq > m.seq {
		m.seq = msg.Seq
	}
	event := msg.key()
	if handler, ok := m.entry[event]; ok {
		go handler(conn, msg)
	} else {
		log.Printf("event %s are not supported\n", event)
	}
}
