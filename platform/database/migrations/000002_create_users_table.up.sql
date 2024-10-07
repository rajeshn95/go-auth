-- Add UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Set timezone
-- For more information, please visit:
-- https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
-- SET TIMEZONE="Europe/Moscow";

-- Create users table
CREATE TABLE users (
    uuid UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    password VARCHAR (255) NOT NULL,
    profile JSONB NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL
);