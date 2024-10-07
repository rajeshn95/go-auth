-- Add UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Set timezone
-- For more information, please visit:
-- https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
-- SET TIMEZONE="Europe/Moscow";

-- Create attributes table
CREATE TABLE attributes (
    uuid UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    user_uuid UUID NOT NULL,
    type VARCHAR (255) NULL,
    value VARCHAR (255) NULL,
    purpose VARCHAR (255) NULL,
    verified_at TIMESTAMP NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL
);
