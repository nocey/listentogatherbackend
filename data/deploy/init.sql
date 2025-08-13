-- Deploy listentogether:init to pg
BEGIN;

CREATE SCHEMA listentogether;

-- Create before foreign keys
CREATE TABLE
    user_permissions (
        user_id BIGSERIAL,
        permission_id BIGSERIAL,
        CONSTRAINT user_permissions_pkey PRIMARY KEY (user_id, permission_id)
    );

CREATE TABLE
    permissions (
        id BIGSERIAL,
        created_at TIMESTAMPTZ DEFAULT NOW (),
        updated_at TIMESTAMPTZ DEFAULT NOW (),
        deleted_at TIMESTAMPTZ,
        name TEXT NOT NULL UNIQUE,
        PRIMARY KEY (id)
    );

CREATE TABLE
    users (
        id BIGSERIAL,
        created_at TIMESTAMPTZ DEFAULT NOW (),
        updated_at TIMESTAMPTZ DEFAULT NOW (),
        deleted_at TIMESTAMPTZ,
        name TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL DEFAULT '',
        PRIMARY KEY (id)
    );

ALTER TABLE user_permissions ADD CONSTRAINT user_fkey FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE user_permissions ADD CONSTRAINT permission_fkey FOREIGN KEY (permission_id) REFERENCES permissions (id);

COMMIT;