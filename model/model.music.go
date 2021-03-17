package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	uuid "github.com/satori/go.uuid"
)

type (
	Music struct {
		ID          uuid.UUID `json:"id"`
		UserID      uuid.UUID `json:"user_id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		FilePath    string    `json:"file_path"`
	}
)

func (u *Music) Add(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	query := `INSERT INTO "music" (user_id,title,description,file_path) VALUES ($1,$2,$3,$4) RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), u.UserID, u.Title, u.Description, u.FilePath).Scan(&u.ID)
	if err != nil {
		return u.ID, err
	}

	return u.ID, nil
}

func (u *Music) All(ctx context.Context, db *sql.DB, param ListQuery) ([]Music, error) {
	all := []Music{}

	query, args, err := param.Query(strings.Split("id,user_id,title,description,file_path", ","))
	if err != nil {
		return all, err
	}

	rows, err := db.QueryContext(ctx, fmt.Sprintf(`SELECT id,user_id,title,description,file_path FROM "music" %s`, query), args...)
	if err != nil {
		return all, err
	}

	defer rows.Close()

	for rows.Next() {
		one := Music{}
		err = rows.Scan(
			&one.ID, &one.UserID, &one.Title, &one.Description, &one.FilePath,
		)
		if err != nil {
			return all, err
		}
		all = append(all, one)
	}

	return all, nil
}

func (u *Music) One(ctx context.Context, db *sql.DB) (Music, error) {
	one := Music{}

	query := `SELECT id,user_id,title,description,file_path FROM "music" WHERE id = $1 LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), u.ID).Scan(
		&one.ID, &one.UserID, &one.Title, &one.Description, &one.FilePath,
	)
	if err != nil {
		return one, err
	}

	return one, nil
}
func (u *Music) OneRandom(ctx context.Context, db *sql.DB) (Music, error) {
	one := Music{}

	query := `SELECT id,user_id,title,description,file_path FROM "music" WHERE user_id = $1 ORDER BY RANDOM() LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), u.UserID).Scan(
		&one.ID, &one.UserID, &one.Title, &one.Description, &one.FilePath,
	)
	if err != nil {
		return one, err
	}

	return one, nil
}

func (u *Music) Update(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID

	query := `UPDATE "music" SET user_id = $1, title = $2, description = $3, file_path = $4 WHERE id = $5 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), u.UserID, u.Title, u.Description, u.FilePath, u.ID).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (u *Music) Delete(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID

	query := `DELETE FROM "music" WHERE id = $1 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), u.ID).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}
