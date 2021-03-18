CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

DROP TABLE IF EXISTS "user" CASCADE;
DROP TABLE IF EXISTS "music" CASCADE;

CREATE TABLE "user" (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    name TEXT NOT NULL DEFAULT '',
    phone_number TEXT NOT NULL DEFAULT '',
    password TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()::TIMESTAMPTZ,
    flag_status INT NOT NULL DEFAULT 0,
    PRIMARY KEY (id)
);

CREATE TABLE "music" (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES "user" (id),
    title TEXT NOT NULL DEFAULT '',
    description TEXT NOT NULL DEFAULT '', 
    file_path TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()::TIMESTAMPTZ,
    flag_status INT NOT NULL DEFAULT 0,
    PRIMARY KEY (id)
);

