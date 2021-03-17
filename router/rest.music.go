package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/renosyah/private-music-player/api"
	"github.com/renosyah/private-music-player/model"
	uuid "github.com/satori/go.uuid"
)

func HandlerAddMusic(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()

	var param model.Music

	err := ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Music/create/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	return musicModule.Add(ctx, param)
}

func HandlerAllMusic(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	var param model.ListQuery

	err := ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Music/all/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	return musicModule.All(ctx, param)
}

func HandlerOneMusic(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	id, err := uuid.FromString(vars["id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Music/detail"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	return musicModule.One(ctx, model.Music{ID: id})
}

func HandlerOneMusicRandom(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	id, err := uuid.FromString(vars["user_id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Music/detail"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	return musicModule.OneRandom(ctx, model.Music{UserID: id})
}

func HandlerUpdateMusic(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	var param model.Music

	id, err := uuid.FromString(vars["id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Music/update"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	err = ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Music/update/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	param.ID = id

	return musicModule.Update(ctx, param)
}

func HandlerDeleteMusic(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	id, err := uuid.FromString(vars["id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Music/delete"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	return musicModule.Delete(ctx, model.Music{ID: id})
}
