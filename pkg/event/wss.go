package event

import (
	"context"
	"log"

	"github.com/eztalk/pkg/service"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

func New(service service.Service) *Listener {
	return &Listener{
		mux: newMux(service),
		ctx: context.Background(),
	}
}

type Listener struct {
	url string
	mux *mux
	ctx context.Context
}

func (w *Listener) Listen(url string) error {
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
