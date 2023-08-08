-- name: CreateInstore :one
INSERT INTO instore (
  book,
  owner,
  bookcount
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetInstore :one
SELECT * FROM instore
WHERE book = $1 LIMIT 1;

-- name: GetInstoreForUpdate :one
SELECT * FROM instore
WHERE book = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListinStore :many
SELECT * FROM instore
ORDER BY book
LIMIT $1
OFFSET $2;

-- name: UpdateinStore :one
UPDATE instore
SET bookcount = $3
WHERE book = $1 AND owner = $2
RETURNING *;

