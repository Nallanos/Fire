-- migrate:up

CREATE TABLE applications (
    id text primary key not null,
    name text not null,
    status text not null,
    image text not null,
    port text not null,
    user_id text not null references users(id) on delete cascade
);

-- migrate:down

DROP TABLE applications;
