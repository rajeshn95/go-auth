-- Add UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Set timezone
-- For more information, please visit:
-- https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
-- SET TIMEZONE="Europe/Moscow";

-- Create tokens table
CREATE TABLE tokens (
    uuid UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    sub UUID NULL,
    issuer VARCHAR (255) NULL,
    title VARCHAR (255) NULL,
    permissions JSONB NULL,
    other_info JSONB NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL
);
