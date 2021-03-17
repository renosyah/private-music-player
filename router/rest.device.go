package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/renosyah/private-music-player/api"
	"github.com/renosyah/private-music-player/model"
	uuid "github.com/satori/go.uuid"
)

func HandlerAddDevice(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()

	var param model.Device

	err := ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Device/create/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	return deviceModule.Add(ctx, param)
}

func HandlerAllDevice(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	var param model.ListQuery

	err := ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Device/all/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	return deviceModule.All(ctx, param)
}

func HandlerOneDevice(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	id, err := uuid.FromString(vars["id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Device/detail"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	return deviceModule.One(ctx, model.Device{ID: id})
}

func HandlerUpdateDevice(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	var param model.Device

	id, err := uuid.FromString(vars["id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Device/update"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	err = ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Device/update/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	param.ID = id

	return deviceModule.Update(ctx, param)
}

func HandlerDeleteDevice(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	id, err := uuid.FromString(vars["id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Device/delete"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	return deviceModule.Delete(ctx, model.Device{ID: id})
}
