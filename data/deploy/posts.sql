-- Deploy listentogether:posts to pg

BEGIN;

CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    user_id INTEGER NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ DEFAULT NULL
);

ALTER TABLE posts
ADD CONSTRAINT user_fkey FOREIGN KEY (user_id) REFERENCES users (id);

INSERT INTO
    permissions (name)
VALUES ('create_post'),
    ('read_post'),
    ('update_post'),
    ('delete_post');

COMMIT;