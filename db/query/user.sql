-- name: CreateUser :one
INSERT INTO users (
 first_name,last_name, email, role
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;


-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;


-- name: UpdateUser :one
UPDATE users SET first_name = $1, last_name = $2, email = $3, role = $4
WHERE id = $5
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;