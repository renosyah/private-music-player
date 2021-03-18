package router

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/renosyah/private-music-player/model"
	uuid "github.com/satori/go.uuid"
)

func (h *RouterHub) HandleStreamHome(w http.ResponseWriter, r *http.Request) {

	var err error
	ctx := r.Context()
	device := model.Device{
		ID:   r.FormValue("id"),
		Name: r.FormValue("name"),
		Role: model.ROLE_MEDIA_CONTROLLER,
	}

	if device.ID == "" || device.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	device.UserID, err = uuid.FromString(r.FormValue("u_id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	wsconn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.addDevice(device)
	defer h.removeDevice(device)

	defer wsconn.Close()

	go h.receiveBroadcastsEvent(ctx, wsconn, device.ID)

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
