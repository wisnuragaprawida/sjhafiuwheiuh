BEGIN;

create table if not exists "users" (
    id serial primary key,
    name text not null,
    email text not null,
    password text not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp
);

INSERT INTO users (name, email, password) VALUES ('admin', 'admin@localhost', 'admin');

COMMIT;