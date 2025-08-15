-- Revert listentogether:posts from pg

BEGIN;

DROP TABLE posts CASCADE;

COMMIT;