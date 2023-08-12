-- name: CreateAccount :one
INSERT INTO accounts (
  owner
) VALUES (
  $1
)
RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY id;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;