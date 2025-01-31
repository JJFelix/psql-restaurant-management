-- name: CreateFood :one
INSERT INTO food(id, name, price, food_image, created_at, updated_at, food_id, menu_id)
VALUES($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetFoodByFoodID :one
SELECT * FROM food WHERE food_id = $1;