-- name: GetCard :one
SELECT * FROM cards
WHERE id = $1 LIMIT 1;

-- name: ListCards :many
SELECT * FROM cards
WHERE owner_id = $1
ORDER BY created_at DESC;

-- name: CreateCard :one
INSERT INTO cards (
  owner_id, type, number, exp_date, cvv, balance
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;