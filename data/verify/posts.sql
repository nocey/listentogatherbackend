-- Verify listentogether:posts on pg

BEGIN;

SELECT
    id,
    content,
    user_id,
    title,
    created_at,
    updated_at
FROM posts
WHERE
    FALSE;

ROLLBACK;