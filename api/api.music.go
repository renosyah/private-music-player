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
	MusicModule struct {
		db   *sql.DB
		Name string
	}
)

func NewMusicModule(db *sql.DB) *MusicModule {
	return &MusicModule{
		db:   db,
		Name: "module/Music",
	}
}

func (m MusicModule) All(ctx context.Context, param model.ListQuery) ([]model.Music, *Error) {
	var all []model.Music

	data, err := (&model.Music{}).All(ctx, m.db, param)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on query all Music"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no Music found"
		}
		return []model.Music{}, NewErrorWrap(err, m.Name, "all/Music",
			message, status)
	}
	for _, each := range data {
		all = append(all, each)
	}
	return all, nil

}
func (m MusicModule) Add(ctx context.Context, param model.Music) (model.Music, *Error) {

	id, err := param.Add(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on add Music"

		return model.Music{}, NewErrorWrap(err, m.Name, "add/Music",
			message, status)
	}
	param.ID = id
	return param, nil
}

func (m MusicModule) One(ctx context.Context, param model.Music) (model.Music, *Error) {
	data, err := param.One(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on get one Music"

		return model.Music{}, NewErrorWrap(err, m.Name, "one/Music",
			message, status)
	}

	return data, nil
}

func (m MusicModule) OneRandom(ctx context.Context, param model.Music) (model.Music, *Error) {
	data, err := param.OneRandom(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on get one Music"

		return model.Music{}, NewErrorWrap(err, m.Name, "one/Music",
			message, status)
	}

	return data, nil
}

func (m MusicModule) Update(ctx context.Context, param model.Music) (model.Music, *Error) {
	var emptyUUID uuid.UUID

	i, err := param.Update(ctx, m.db)
	if err != nil || i == emptyUUID {
		status := http.StatusInternalServerError
		message := "error on update Music"

		return model.Music{}, NewErrorWrap(err, m.Name, "update/Music",
			message, status)
	}
	return param, nil
}

func (m MusicModule) Delete(ctx context.Context, param model.Music) (model.Music, *Error) {
	var emptyUUID uuid.UUID
	i, err := param.Delete(ctx, m.db)
	if err != nil || i == emptyUUID {
		status := http.StatusInternalServerError
		message := "error on delete Music"

		return model.Music{}, NewErrorWrap(err, m.Name, "delete/Music",
			message, status)
	}
	return param, nil
}
