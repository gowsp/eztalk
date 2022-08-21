package wss

import (
	"encoding/json"
	"log"
	"time"

	"github.com/eztalk/pkg/eztalk"
	"nhooyr.io/websocket"
)

type heartbeat struct {
	Interval uint64 `json:"heartbeat_interval,omitempty"`
}

func (h *heartbeat) start(c *websocket.Conn, seq func() uint64) {
	t := time.NewTicker(time.Millisecond * time.Duration(h.Interval))
	for {
		select {
		case <-t.C:
			err := write(c, Heartbeat, seq())
			if err != nil {
				log.Println(err)
				t.Stop()
				return
			}
		}
	}
}

type auth struct {
	Token   string `json:"token,omitempty"`
	Intents int    `json:"intents,omitempty"`
}

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

func (m *mux) ping(c *websocket.Conn, e *event) {
	go func() {
		m.heartbeat.start(c, func() uint64 {
			return m.seq
		})
	}()
}

func (m *mux) message(c *websocket.Conn, e *event) {
	msg := new(eztalk.Message)
	if err := json.Unmarshal(e.Data, msg); err != nil {
		log.Println(err)
	}
	m.service.Serve(msg)
}
