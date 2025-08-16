-- Verify listentogether:app_config on pg

BEGIN;

SELECT *
FROM app_config
WHERE
    name = 'app_name'
    AND is_active = true
    AND value = 'ListenTogether';

ROLLBACK;