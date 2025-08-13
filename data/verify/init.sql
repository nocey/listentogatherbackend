-- Verify listentogether:init on pg
BEGIN;

CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- Permission table verify
SELECT
    id,
    created_at,
    updated_at,
    deleted_at,
    name
FROM permissions
WHERE
    FALSE;

-- User table verify
SELECT
    id,
    created_at,
    updated_at,
    deleted_at,
    name
FROM permissions
WHERE
    FALSE;

-- user_permissions table verify (many_to_many)
SELECT user_id, permission_id FROM user_permissions WHERE FALSE;

DO $$ DECLARE user_name TEXT := encode (gen_random_bytes (32), 'hex');

user_password TEXT := encode (gen_random_bytes (32), 'hex');

new_user_id BIGINT;

new_permission_id BIGINT;

BEGIN
INSERT INTO
    users (name, password)
VALUES (user_name, user_password) RETURNING id INTO new_user_id;

INSERT INTO
    permissions (name)
VALUES (user_name) RETURNING id INTO new_permission_id;

-- Verify with LOG
RAISE NOTICE 'User ID: %, Permission ID: %',
new_user_id,
new_permission_id;

-- Verify admin user and permission relationship exists
PERFORM *
FROM user_permissions
WHERE
    user_id = new_user_id
    AND permission_id = new_permission_id;

END $$;

ROLLBACK;