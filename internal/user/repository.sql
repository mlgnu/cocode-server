-- name: GetUser :one
SELECT * FROM "user"
WHERE "id" = $1;


-- name: UpdateUser :exec
UPDATE "user" SET "email" = $1, "password" = $2, "first_name" = $3, "last_name" = $4, "avatar" = $5 WHERE "id" = $6 RETURNING email, first_name, last_name, avatar, role, is_active, created_at, updated_at;

-- name: DeleteUser :exec
DELETE FROM "user" WHERE "id" = $1;

-- name: ChangePassword :exec
UPDATE "user" SET "password" = $1 WHERE "id" = $2;
