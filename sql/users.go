package sql

import (
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
)

func InsertUser(ctx context.Context, db Querier, name, email, password string) (id int64, err error) {
	err = pgxscan.Get(ctx, db, &id, `insert into users (full_name, email, password) values($1, $2, $3) returning id`, name, email, password)
	return
}

type UserEssentials struct {
	FullName string `db:"full_name"`
	Password string `db:"password"`
}

func SelectUserEssentials(ctx context.Context, db Querier, email string) (data UserEssentials, err error) {
	err = pgxscan.Get(ctx, db, &data, `select full_name, password from users where email = $1`, email)
	return
}

type User struct {
	Name    string `json:"name" db:"full_name"`
	Email   string `json:"email" db:"email"`
	IsAdmin bool   `json:"isAdmin" db:"is_admin"`
}

func SelectUser(ctx context.Context, db Querier, id int64) (data User, err error) {
	err = pgxscan.Get(ctx, db, &data, `select full_name, email, is_admin from users where id = $1`, id)
	return
}