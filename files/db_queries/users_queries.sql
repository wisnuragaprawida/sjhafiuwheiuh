-- name: GetUsers :many
select * from users;

-- name: GetUser :one
select * from users where id = $1;

-- name: RegisterUser :one
insert into users (name, email, password) values ($1, $2, $3) returning *;

-- name: UpdateUser :one
update users set name = $1, email = $2, password = $3 where id = $4 returning *;

-- name: DeleteUser :exec
delete from users where id = $1;

-- name: FindUserByEmail :one
select * from users where email = $1;