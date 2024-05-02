-- migrate:up

CREATE TABLE applications (
    id text primary key not null,
    name text not null,
    status text not null,
    image text not null,
    port integer not null
);

-- migrate:down

DROP TABLE applications;
