-- Verify listentogether:posts on pg

BEGIN;

SELECT
    id,
    content,
    author_id,
    title,
    created_at,
    updated_at
FROM posts
WHERE
    FALSE;

DELETE FROM permissions
WHERE
    name IN (
        'create_post',
        'read_post',
        'update_post',
        'delete_post'
    );

ROLLBACK;