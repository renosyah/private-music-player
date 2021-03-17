package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	uuid "github.com/satori/go.uuid"
)

type (
	User struct {
		ID          uuid.UUID `json:"id"`
		Name        string    `json:"name"`
		PhoneNumber string    `json:"phone_number"`
		Password    string    `json:"password"`
	}
)

func (u *User) Add(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	query := `INSERT INTO "user" (name,phone_number,password) VALUES ($1,$2,$3) RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), u.Name, u.PhoneNumber, u.Password).Scan(&u.ID)
	if err != nil {
		return u.ID, err
	}

	return u.ID, nil
}

func (u *User) All(ctx context.Context, db *sql.DB, param ListQuery) ([]User, error) {
	all := []User{}

	query, args, err := param.Query(strings.Split("id,name,phone_number", ","))
	if err != nil {
		return all, err
	}

	rows, err := db.QueryContext(ctx, fmt.Sprintf(`SELECT id,name,phone_number,'' FROM "user" %s`, query), args...)
	if err != nil {
		return all, err
	}

	defer rows.Close()

	for rows.Next() {
		one := User{}
		err = rows.Scan(
			&one.ID, &one.Name, &one.PhoneNumber, &one.Password,
		)
		if err != nil {
			return all, err
		}
		all = append(all, one)
	}

	return all, nil
}

func (u *User) One(ctx context.Context, db *sql.DB) (User, error) {
	one := User{}

	query := `SELECT id,name,phone_number,password FROM "user" WHERE id = $1 LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), u.ID).Scan(
		&one.ID, &one.Name, &one.PhoneNumber, &one.Password,
	)
	if err != nil {
		return one, err
	}

	return one, nil
}

func (u *User) OneByPhoneNumber(ctx context.Context, db *sql.DB) (User, error) {
	one := User{}

	query := `SELECT id,name,phone_number,password FROM "user" WHERE phone_number = $1 LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), u.PhoneNumber).Scan(
		&one.ID, &one.Name, &one.PhoneNumber, &one.Password,
	)
	if err != nil {
		return one, err
	}

	return one, nil
}

func (u *User) Update(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID

	query := `UPDATE "user" SET name = $1, phone_number = $2 WHERE id = $3 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), u.Name, u.PhoneNumber, u.ID).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (u *User) Delete(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID

	query := `DELETE FROM "user" WHERE id = $1 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), u.ID).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}
