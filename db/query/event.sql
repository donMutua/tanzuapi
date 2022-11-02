-- name: CreateEvent :one
INSERT INTO events (
 event_name, about, cost, event_date, start_time, end_time
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;


-- name: GetEvent :one
SELECT * FROM events
WHERE id = $1 LIMIT 1;

-- name: ListEvents :many
SELECT * FROM events
ORDER BY id
LIMIT $1
OFFSET $2;


-- name: UpdateEvent :one
UPDATE events SET id = $1 , event_name =$2 ,about = $3, cost = $4, event_date = $5, start_time = $6, end_time = $7
WHERE id = $1
RETURNING *;

-- name: DeleteEvent :exec
DELETE FROM events WHERE id = $1;