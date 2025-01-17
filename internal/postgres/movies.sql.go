// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: movies.sql

package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const insertMovie = `-- name: InsertMovie :exec
insert into movies (name, movie_time, genres, premiere_time, premiere_from_date, premiere_to_date) values ($1, $2, $3, $4, $5, $6)
`

type InsertMovieParams struct {
	Name             string             `json:"name"`
	MovieTime        int16              `json:"movieTime"`
	Genres           []string           `json:"genres"`
	PremiereTime     pgtype.Timestamptz `json:"premiereTime"`
	PremiereFromDate pgtype.Date        `json:"premiereFromDate"`
	PremiereToDate   pgtype.Date        `json:"premiereToDate"`
}

func (q *Queries) InsertMovie(ctx context.Context, arg InsertMovieParams) error {
	_, err := q.db.Exec(ctx, insertMovie,
		arg.Name,
		arg.MovieTime,
		arg.Genres,
		arg.PremiereTime,
		arg.PremiereFromDate,
		arg.PremiereToDate,
	)
	return err
}

const selectMovie = `-- name: SelectMovie :one
select id, name, movie_time, genres, premiere_from_date, premiere_to_date, premiere_time from movies where id = $1
`

type SelectMovieRow struct {
	ID               int64              `json:"id"`
	Name             string             `json:"name"`
	MovieTime        int16              `json:"movieTime"`
	Genres           []string           `json:"genres"`
	PremiereFromDate pgtype.Date        `json:"premiereFromDate"`
	PremiereToDate   pgtype.Date        `json:"premiereToDate"`
	PremiereTime     pgtype.Timestamptz `json:"premiereTime"`
}

func (q *Queries) SelectMovie(ctx context.Context, id int64) (SelectMovieRow, error) {
	row := q.db.QueryRow(ctx, selectMovie, id)
	var i SelectMovieRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.MovieTime,
		&i.Genres,
		&i.PremiereFromDate,
		&i.PremiereToDate,
		&i.PremiereTime,
	)
	return i, err
}

const selectMovieReservedSeats = `-- name: SelectMovieReservedSeats :many
select 
    s.seat_number::int2
from 
    tickets t 
    left join seats s on s.ticket_id = t.id
where
    t.movie_id = $1
    and t.premiere_date = $2
`

type SelectMovieReservedSeatsParams struct {
	MovieID      int64       `json:"movieId"`
	PremiereDate pgtype.Date `json:"premiereDate"`
}

func (q *Queries) SelectMovieReservedSeats(ctx context.Context, arg SelectMovieReservedSeatsParams) ([]int16, error) {
	rows, err := q.db.Query(ctx, selectMovieReservedSeats, arg.MovieID, arg.PremiereDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int16{}
	for rows.Next() {
		var s_seat_number int16
		if err := rows.Scan(&s_seat_number); err != nil {
			return nil, err
		}
		items = append(items, s_seat_number)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectMovies = `-- name: SelectMovies :many
select id, name, premiere_time from movies order by premiere_time asc
`

type SelectMoviesRow struct {
	ID           int64              `json:"id"`
	Name         string             `json:"name"`
	PremiereTime pgtype.Timestamptz `json:"premiereTime"`
}

func (q *Queries) SelectMovies(ctx context.Context) ([]SelectMoviesRow, error) {
	rows, err := q.db.Query(ctx, selectMovies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SelectMoviesRow{}
	for rows.Next() {
		var i SelectMoviesRow
		if err := rows.Scan(&i.ID, &i.Name, &i.PremiereTime); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
