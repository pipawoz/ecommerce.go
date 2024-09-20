-- name: GetOrder :one
SELECT * FROM orders
WHERE id = $1 LIMIT 1;

-- name: ListOrders :many
SELECT * FROM orders
ORDER BY id;

-- name: CreateOrder :one
INSERT INTO orders (
  customer_id, status, total_amount
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateOrder :one
UPDATE orders
SET status = $2, total_amount = $3, updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteOrder :exec
DELETE FROM orders
WHERE id = $1;