package event

import (
	"log"
	"time"

	"nhooyr.io/websocket"
)

// 心跳结构体
type heartbeat struct {
	Interval uint64 `json:"heartbeat_interval,omitempty"`
}

// 开始周期性发送心跳
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
