-- Revert listentogether:posts from pg

BEGIN;

DROP TABLE posts CASCADE;

DELETE FROM permissions
WHERE
    name IN (
        'create_post',
        'read_post',
        'update_post',
        'delete_post'
    );

COMMIT;