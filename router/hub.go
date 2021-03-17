package router

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/renosyah/private-music-player/model"
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type (
	RouterHub struct {
		ConnectionMx sync.RWMutex

		// event in home
		Subscriber     map[string]chan model.EventData
		EventBroadcast chan model.EventData
	}
)

func NewRouterHub() *RouterHub {
	h := &RouterHub{
		ConnectionMx:   sync.RWMutex{},
		Subscriber:     make(map[string]chan model.EventData),
		EventBroadcast: make(chan model.EventData),
	}
	go func() {
		for {
			msg := <-h.EventBroadcast
			h.ConnectionMx.RLock()
			for _, subReceiver := range h.Subscriber {
				select {
				case subReceiver <- msg:
				default:
				}

			}
			h.ConnectionMx.RUnlock()
		}

	}()

	return h
}
