-- Deploy listentogether:app_config to pg

BEGIN;

CREATE TABLE app_config (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    is_active BOOLEAN NOT NULL,
    value TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

INSERT INTO
    app_config (name, is_active, value)
VALUES (
        'app_name',
        true,
        'ListenTogether'
    ),
    (
        'app_description',
        true,
        'A collaborative music listening platform'
    ),
    (
        'app_logo',
        true,
        'https://example.com/logo.png'
    ),
    (
        'app_favicon',
        true,
        'https://example.com/favicon.ico'
    ),
    (
        'app_theme_color',
        true,
        '#f97c56ff'
    );

COMMIT;