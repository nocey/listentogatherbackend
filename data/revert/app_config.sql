-- Revert listentogether:app_config from pg

BEGIN;

DROP TABLE app_cofig CASCADE;

COMMIT;