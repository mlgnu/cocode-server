-- name: AddUser :one
INSERT INTO "user" ("email", "password", "first_name", "last_name") VALUES ($1, $2, $3, $4) RETURNING email, first_name, last_name;

-- name: GetUserByEmail :one
SELECT email, first_name, last_name, avatar, role, is_active, created_at, updated_at FROM "user" WHERE "email" = $1;

-- name: GetUserAuth :one
SELECT id, email, password, role, is_active FROM "user" WHERE "email" = $1;

-- name: GetUserById :one
SELECT email, first_name, last_name, avatar, role, is_active, created_at, updated_at FROM "user" WHERE "id" = $1;

-- name: UpdateUser :one
UPDATE "user" SET "email" = $1, "password" = $2, "first_name" = $3, "last_name" = $4, "avatar" = $5 WHERE "id" = $6 RETURNING email, first_name, last_name, avatar, role, is_active, created_at, updated_at;
