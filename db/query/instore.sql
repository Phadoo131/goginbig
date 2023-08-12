-- name: CreateInstore :one
INSERT INTO instore (
  book,
  bookcount
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetInstore :one
SELECT * FROM instore;

-- name: GetInstoreForUpdate :one
SELECT * FROM instore
WHERE book = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListInstore :many
SELECT * FROM instore
ORDER BY book
LIMIT $1
OFFSET $2;

-- name: UpdateInstore :one
UPDATE instore
SET bookcount = $2
WHERE book = $1
RETURNING *;

