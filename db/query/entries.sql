-- name: CreateEntry :one
INSERT INTO entries (
  account_id,
  book,
  amount
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetEntry :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: ListEntries :many
SELECT * FROM entries
WHERE account_id = $1
ORDER BY id
LIMIT $3
OFFSET $2;

-- name: GetEntryForUpdate :one
SELECT * FROM entries
WHERE account_id = $1 LIMIT 1
FOR NO KEY UPDATE;


-- name: UpdateEntries :one
UPDATE entries
SET amount = $3
WHERE book = $2 AND account_id = $1
RETURNING *;
