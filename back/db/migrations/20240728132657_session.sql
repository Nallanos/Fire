-- migrate:up


CREATE TABLE tokens (
    token text primary key,
    user_id text not null references users(id) on delete cascade,
    expires_at timestamp not null
)
  
-- migrate:down


DROP TABLE tokens