-- migrate:up

CREATE TABLE users (
    id text primary key not null,
    email text not null unique,
    name text not null,
    password text not null
);

-- migrate:down

DROP TABLE users;