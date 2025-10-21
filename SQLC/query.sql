-- name: GetProduct :one
SELECT * FROM product WHERE id = $1;

-- name: ListProducts :many
SELECT * FROM product ORDER BY name;