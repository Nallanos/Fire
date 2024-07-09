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
-- Dbmate schema migrations
INSERT INTO "schema_migrations" (version) VALUES
  ('20240208190923'),
  ('20240320193040');
