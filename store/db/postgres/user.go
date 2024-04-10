package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/deepaksing/Travegram/store"
)

func (d *DB) CreateUser(ctx context.Context, user *store.User) (*store.User, error) {
	query := `
		INSERT INTO "user" (username, password_hash, name, avatar_url)
		VALUES ($1, $2, $3, $4)
		RETURNING user_id
	`

	var user_id int32
	err := d.db.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.PasswordHash,
		user.Name,
		"http://sss.com",
	).Scan(&user_id)
	fmt.Println("err1 ", err)
	if err != nil {
		return nil, err
	}

	// Use the returned ID to retrieve the full user object
	user, err = d.GetUser(ctx, &store.FindUser{ID: &user_id})
	fmt.Println("Err 2 ", err)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (d *DB) GetUser(ctx context.Context, f *store.FindUser) (*store.User, error) {
	query := `
		SELECT user_id, username, password_hash, name
		FROM "user"
	`

	// Execute the query
	var row *sql.Row
	if v := f.ID; v != nil {
		query += "WHERE user_id = $1"
		row = d.db.QueryRowContext(ctx, query, *f.ID)
	}
	if v := f.Username; v != nil {
		query += "WHERE username = $1"
		row = d.db.QueryRowContext(ctx, query, *f.Username)
	}
	fmt.Println("query ", query)
	fmt.Println("row ", row)
	// Scan the row into a user object
	user := &store.User{}
	err := row.Scan(&user.ID, &user.Username, &user.PasswordHash, &user.Name)
	if err != nil {
		return nil, err
	}
	fmt.Println("err 3", err)

	return user, nil
}
