package api

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/renosyah/private-music-player/model"
	uuid "github.com/satori/go.uuid"
)

type (
	DeviceModule struct {
		db   *sql.DB
		Name string
	}
)

func NewDeviceModule(db *sql.DB) *DeviceModule {
	return &DeviceModule{
		db:   db,
		Name: "module/Device",
	}
}

func (m DeviceModule) All(ctx context.Context, param model.ListQuery) ([]model.Device, *Error) {
	var all []model.Device

	data, err := (&model.Device{}).All(ctx, m.db, param)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on query all Device"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no Device found"
		}
		return []model.Device{}, NewErrorWrap(err, m.Name, "all/Device",
			message, status)
	}
	for _, each := range data {
		all = append(all, each)
	}
	return all, nil

}
func (m DeviceModule) Add(ctx context.Context, param model.Device) (model.Device, *Error) {

	id, err := param.Add(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on add Device"

		return model.Device{}, NewErrorWrap(err, m.Name, "add/Device",
			message, status)
	}
	param.ID = id
	return param, nil
}

func (m DeviceModule) One(ctx context.Context, param model.Device) (model.Device, *Error) {
	data, err := param.One(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on get one Device"

		return model.Device{}, NewErrorWrap(err, m.Name, "one/Device",
			message, status)
	}

	return data, nil
}

func (m DeviceModule) Update(ctx context.Context, param model.Device) (model.Device, *Error) {
	var emptyUUID uuid.UUID

	i, err := param.Update(ctx, m.db)
	if err != nil || i == emptyUUID {
		status := http.StatusInternalServerError
		message := "error on update Device"

		return model.Device{}, NewErrorWrap(err, m.Name, "update/Device",
			message, status)
	}
	return param, nil
}

func (m DeviceModule) Delete(ctx context.Context, param model.Device) (model.Device, *Error) {
	var emptyUUID uuid.UUID
	i, err := param.Delete(ctx, m.db)
	if err != nil || i == emptyUUID {
		status := http.StatusInternalServerError
		message := "error on delete Device"

		return model.Device{}, NewErrorWrap(err, m.Name, "delete/Device",
			message, status)
	}
	return param, nil
}
