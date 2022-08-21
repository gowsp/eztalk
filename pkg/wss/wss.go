package wss

import (
	"context"
	"log"

	"github.com/eztalk/pkg/eztalk"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

func New(service eztalk.Service) *Wss {
	return &Wss{
		mux: NewMux(service),
		ctx: context.Background(),
	}
}

type Wss struct {
	url string
	mux *mux
	ctx context.Context
}

func (w *Wss) Connect(url string) error {
	log.Println("start connect to", url)
	conn, _, err := websocket.Dial(context.Background(), url, nil)
	if err != nil {
		return err
	}
	for {
		msg := new(event)
		if err := wsjson.Read(w.ctx, conn, msg); err != nil {
			return err
		}
		w.mux.Serve(conn, msg)
	}
}
