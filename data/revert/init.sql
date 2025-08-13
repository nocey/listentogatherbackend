-- Revert listentogether:init from pg
BEGIN;

DROP SCHEMA listentogether;

DROP TABLE permissions CASCADE;

DROP TABLE users CASCADE;

DROP TABLE user_permissions CASCADE;

COMMIT;