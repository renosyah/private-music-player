package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/renosyah/private-music-player/api"
	"github.com/renosyah/private-music-player/model"
	uuid "github.com/satori/go.uuid"
)

func HandlerAddUser(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()

	var param model.User

	err := ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "User/create/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	return userModule.Add(ctx, param)
}

func HandlerAllUser(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	var param model.ListQuery

	err := ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "User/all/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	return userModule.All(ctx, param)
}

func HandlerOneUser(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	id, err := uuid.FromString(vars["id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "User/detail"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	return userModule.One(ctx, model.User{ID: id})
}

func HandlerLoginUser(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()

	var param model.User

	err := ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "User/login/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	return userModule.Login(ctx, param)
}

func HandlerRegisterUser(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()

	var param model.User

	err := ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "User/register/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	return userModule.Register(ctx, param)
}

func HandlerUpdateUser(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	var param model.User

	id, err := uuid.FromString(vars["id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "User/update"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	err = ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "User/update/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	param.ID = id

	return userModule.Update(ctx, param)
}

func HandlerDeleteUser(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	id, err := uuid.FromString(vars["id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "User/delete"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	return userModule.Delete(ctx, model.User{ID: id})
}
