-- name: CreaterReturnAndBorrow :one
INSERT INTO returnandborrow (
  from_account_id,
  book,
  bookcount
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetTransfer :one
SELECT * FROM returnandborrow
WHERE id = $1 LIMIT 1;

-- name: ListTransfers :many
SELECT * FROM returnandborrow
WHERE from_account_id = $1 OR book = $2
ORDER BY id
LIMIT $3
OFFSET $4;