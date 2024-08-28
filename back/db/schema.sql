CREATE TABLE IF NOT EXISTS "schema_migrations" (version varchar(128) primary key);
CREATE TABLE applications (
    id text primary key not null,
    name text not null,
    status text not null,
    image text not null,
    port text not null
);
CREATE TABLE deployments (
    id text primary key not null,
    app_id text not null,
    status text not null,
    created_at timestamp not null,
    finished_at timestamp,
    FOREIGN KEY (app_id) REFERENCES applications(id) ON DELETE CASCADE
);
CREATE TABLE users (
    id text primary key not null,
    email text not null unique,
    name text not null,
    password text not null
);
CREATE TABLE tokens (
    token text primary key,
    user_id text not null references users(id) on delete cascade,
    expires_at timestamp not null
);
-- Dbmate schema migrations
INSERT INTO "schema_migrations" (version) VALUES
  ('20240208190923'),
  ('20240320193040'),
  ('20240728125311'),
  ('20240728132657');
