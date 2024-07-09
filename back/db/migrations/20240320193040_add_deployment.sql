-- migrate:up

CREATE TABLE deployments (
    id text primary key not null,
    app_id text not null,
    status text not null,
    created_at timestamp not null,
    finished_at timestamp,
    FOREIGN KEY (app_id) REFERENCES applications(id) ON DELETE CASCADE
);

-- migrate:down

DROP TABLE deployments;
