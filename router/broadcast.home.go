package router

import (
	"context"

	"github.com/gorilla/websocket"
	"github.com/renosyah/private-music-player/model"
)

func (h *RouterHub) subscribe(id string) (stream chan model.EventData) {
	h.ConnectionMx.Lock()
	defer h.ConnectionMx.Unlock()

	stream = make(chan model.EventData)
	h.Subscriber[id] = stream

	return
}

func (h *RouterHub) unSubscribe(id string) {
	h.ConnectionMx.Lock()
	defer h.ConnectionMx.Unlock()
	if _, ok := h.Subscriber[id]; ok {
		close(h.Subscriber[id])
		delete(h.Subscriber, id)
	}
}

func (h *RouterHub) receiveBroadcastsEvent(ctx context.Context, wsconn *websocket.Conn, id string) {
	subReceiver := h.subscribe(id)
	defer h.unSubscribe(id)

	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-subReceiver:
			if err := wsconn.WriteMessage(websocket.TextMessage, model.ToJson(msg)); err != nil {
				return
			}
		}
	}

}
