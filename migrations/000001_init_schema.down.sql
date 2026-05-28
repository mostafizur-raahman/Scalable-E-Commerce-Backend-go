-- +migrate Down
-- Remove triggers
DROP TRIGGER IF EXISTS update_products_updated_at ON products;
DROP TRIGGER IF EXISTS update_users_updated_at ON users;
DROP FUNCTION IF EXISTS update_updated_at_column();

-- Remove indexes
DROP INDEX IF EXISTS idx_users_email;
DROP INDEX IF EXISTS idx_products_name;

-- Remove tables
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS products;