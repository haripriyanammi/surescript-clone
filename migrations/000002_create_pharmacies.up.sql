CREATE TABLE pharmacies (
    id           UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
    ncpdp_id     CHAR(7)     NOT NULL UNIQUE CHECK (ncpdp_id ~ '^[0-9]{7}$'),
    name         TEXT        NOT NULL,
    api_key_hash TEXT        NOT NULL UNIQUE,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT now()
);