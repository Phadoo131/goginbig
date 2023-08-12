// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: instore.sql

package db

import (
	"context"
)

const createInstore = `-- name: CreateInstore :one
INSERT INTO instore (
  book,
  bookcount
) VALUES (
  $1, $2
)
RETURNING book, bookcount, created_at
`

type CreateInstoreParams struct {
	Book      string `json:"book"`
	Bookcount int64  `json:"bookcount"`
}

func (q *Queries) CreateInstore(ctx context.Context, arg CreateInstoreParams) (Instore, error) {
	row := q.queryRow(ctx, q.createInstoreStmt, createInstore, arg.Book, arg.Bookcount)
	var i Instore
	err := row.Scan(&i.Book, &i.Bookcount, &i.CreatedAt)
	return i, err
}

const getInstore = `-- name: GetInstore :one
SELECT book, bookcount, created_at FROM instore
`

func (q *Queries) GetInstore(ctx context.Context) (Instore, error) {
	row := q.queryRow(ctx, q.getInstoreStmt, getInstore)
	var i Instore
	err := row.Scan(&i.Book, &i.Bookcount, &i.CreatedAt)
	return i, err
}

const getInstoreForUpdate = `-- name: GetInstoreForUpdate :one
SELECT book, bookcount, created_at FROM instore
WHERE book = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetInstoreForUpdate(ctx context.Context, book string) (Instore, error) {
	row := q.queryRow(ctx, q.getInstoreForUpdateStmt, getInstoreForUpdate, book)
	var i Instore
	err := row.Scan(&i.Book, &i.Bookcount, &i.CreatedAt)
	return i, err
}

const listInstore = `-- name: ListInstore :many
SELECT book, bookcount, created_at FROM instore
ORDER BY book
LIMIT $1
OFFSET $2
`

type ListInstoreParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListInstore(ctx context.Context, arg ListInstoreParams) ([]Instore, error) {
	rows, err := q.query(ctx, q.listInstoreStmt, listInstore, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Instore
	for rows.Next() {
		var i Instore
		if err := rows.Scan(&i.Book, &i.Bookcount, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateInstore = `-- name: UpdateInstore :one
UPDATE instore
SET bookcount = $2
WHERE book = $1
RETURNING book, bookcount, created_at
`

type UpdateInstoreParams struct {
	Book      string `json:"book"`
	Bookcount int64  `json:"bookcount"`
}

func (q *Queries) UpdateInstore(ctx context.Context, arg UpdateInstoreParams) (Instore, error) {
	row := q.queryRow(ctx, q.updateInstoreStmt, updateInstore, arg.Book, arg.Bookcount)
	var i Instore
	err := row.Scan(&i.Book, &i.Bookcount, &i.CreatedAt)
	return i, err
}
