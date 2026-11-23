CREATE TABLE providers (
    id           UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
    npi          CHAR(10)    NOT NULL UNIQUE CHECK (npi ~ '^[0-9]{10}$'),
    name         TEXT        NOT NULL,
    api_key_hash TEXT        NOT NULL UNIQUE,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT now()
);