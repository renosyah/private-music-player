package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/renosyah/private-music-player/model"
)

func (h *RouterHub) HandleStreamHome(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	vars := mux.Vars(r)

	uID := vars["id"]
	if uID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	wsconn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer wsconn.Close()

	go h.receiveBroadcastsEvent(ctx, wsconn, uID)

	for {

		mType, msg, err := wsconn.ReadMessage()
		if err != nil {
			break
		}

		if mType != websocket.TextMessage {
			break
		}

		event := (&model.EventData{}).FromJson(msg)
		h.EventBroadcast <- event
	}
}
