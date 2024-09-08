-- migrate:up

ALTER TABLE applications ADD user_id  text not null references users(id) on delete cascade;

-- migrate:down

ALTER TABLE applications DROP user_id;
