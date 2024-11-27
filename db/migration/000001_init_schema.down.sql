-- 1. Drop the users table
DROP TABLE IF EXISTS users CASCADE;

-- 2. Drop the role_permissions table (if exists)
-- (If you have a many-to-many relationship table between roles and permissions, drop it here)
DROP TABLE IF EXISTS role_permissions CASCADE;

-- 3. Drop the permissions table
DROP TABLE IF EXISTS permissions CASCADE;

-- 4. Drop the roles table
DROP TABLE IF EXISTS roles CASCADE;

-- 5. Drop the UUID extension (if no other tables require it)
-- This step is only needed for PostgreSQL if you're using the uuid-ossp extension.
-- Make sure to only drop the extension if no other objects are using it.
DROP EXTENSION IF EXISTS "uuid-ossp";
