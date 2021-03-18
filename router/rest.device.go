package router

import (
	"net/http"
	"sort"
	"strings"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/renosyah/private-music-player/api"
	"github.com/renosyah/private-music-player/model"
	uuid "github.com/satori/go.uuid"
)

func (h *RouterHub) addDevice(d model.Device) {

	h.EventBroadcast <- model.EventData{UserID: d.UserID.String(), Name: "HOME_EVENT_DEVICE_UPDATE"}

	h.ConnectionMx.Lock()
	defer h.ConnectionMx.Unlock()

	h.Devices[d.ID] = d
}

func (h *RouterHub) removeDevice(d model.Device) {

	h.EventBroadcast <- model.EventData{UserID: d.UserID.String(), Name: "HOME_EVENT_DEVICE_UPDATE"}

	h.ConnectionMx.Lock()
	defer h.ConnectionMx.Unlock()

	if _, exist := h.Devices[d.ID]; exist {
		delete(h.Devices, d.ID)
	}
}

func (h *RouterHub) HandlerUpdateAllDevice(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)
	var param model.Device

	uID, err := uuid.FromString(vars["user_id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Devices/on"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	err = ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Music/update/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	h.ConnectionMx.Lock()
	defer h.ConnectionMx.Unlock()

	for k, v := range h.Devices {
		if v.UserID == uID {
			h.Devices[k] = model.Device{
				ID:     v.ID,
				UserID: v.UserID,
				Name:   v.Name,
				Role:   model.ROLE_MEDIA_CONTROLLER,
			}
		}
	}

	if _, exist := h.Devices[param.ID]; exist {
		h.Devices[param.ID] = param
	}

	return param, nil
}

func (h *RouterHub) HandlerAllDevice(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	vars := mux.Vars(r)

	uID, err := uuid.FromString(vars["user_id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Devices/on"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	h.ConnectionMx.Lock()
	defer h.ConnectionMx.Unlock()

	devices := []model.Device{}
	for _, v := range h.Devices {
		if v.UserID == uID {
			devices = append(devices, v)
		}
	}

	sort.Slice(devices, func(i, j int) bool {
		switch strings.Compare(devices[i].Name, devices[j].Name) {
		case -1:
			return true
		case 1:
			return false
		}
		return devices[i].Name > devices[j].Name
	})

	return devices, nil
}
