package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	uuid "github.com/satori/go.uuid"
)

const (
	ROLE_MEDIA_CONTROLLER = 0
	ROLE_MEDIA_PLAYER     = 1
)

type (
	Device struct {
		ID     uuid.UUID `json:"id"`
		UserID uuid.UUID `json:"user_id"`
		Name   string    `json:"name"`
		Role   int       `json:"role"`
	}
)

func (u *Device) Add(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	query := `INSERT INTO "device" (user_id,name,role) VALUES ($1,$2,$3) RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), u.UserID, u.Name, u.Role).Scan(&u.ID)
	if err != nil {
		return u.ID, err
	}

	return u.ID, nil
}

func (u *Device) All(ctx context.Context, db *sql.DB, param ListQuery) ([]Device, error) {
	all := []Device{}

	query, args, err := param.Query(strings.Split("id,user_id,name,role", ","))
	if err != nil {
		return all, err
	}

	rows, err := db.QueryContext(ctx, fmt.Sprintf(`SELECT id,user_id,name,role FROM "device" %s`, query), args...)
	if err != nil {
		return all, err
	}

	defer rows.Close()

	for rows.Next() {
		one := Device{}
		err = rows.Scan(
			&one.ID, &one.UserID, &one.Name, &one.Role,
		)
		if err != nil {
			return all, err
		}
		all = append(all, one)
	}

	return all, nil
}

func (u *Device) One(ctx context.Context, db *sql.DB) (Device, error) {
	one := Device{}

	query := `SELECT id,user_id,name,role FROM "device" WHERE id = $1 LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), u.ID).Scan(
		&one.ID, &one.UserID, &one.Name, &one.Role,
	)
	if err != nil {
		return one, err
	}

	return one, nil
}

func (u *Device) Update(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID

	query := `UPDATE "device" SET user_id = $1 ,name = $2,role = $3 WHERE id = $4 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), u.UserID, u.Name, u.Role, u.ID).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (u *Device) Delete(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID

	query := `DELETE FROM "device" WHERE id = $1 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), u.ID).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}
