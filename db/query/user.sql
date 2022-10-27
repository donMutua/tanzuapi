-- name: CreateUser :one
INSERT INTO users (
 username, hashed_password, first_name,last_name, email, role
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;


-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY username
LIMIT $1
OFFSET $2;


-- name: UpdateUser :one
UPDATE users SET username = $1 , hashed_password =$2 ,first_name = $3, last_name = $4, email = $5, role = $6
WHERE username = $5
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE username = $1;