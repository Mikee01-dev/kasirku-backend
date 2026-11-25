BEGIN;

-- ============================================================
-- Tabel users
-- ============================================================
CREATE TABLE IF NOT EXISTS users (
    id            BIGSERIAL PRIMARY KEY,
    name          VARCHAR(100) NOT NULL,
    email         VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    role          VARCHAR(20)  NOT NULL DEFAULT 'kasir',
    created_at    TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    CONSTRAINT users_email_unique UNIQUE (email),
    CONSTRAINT users_role_check CHECK (role IN ('admin', 'kasir'))
);

-- ============================================================
-- Tabel products
-- Kolom harus sama persis dengan struct Product di app/model/product.go
-- ============================================================
CREATE TABLE IF NOT EXISTS products (
    id         BIGSERIAL PRIMARY KEY,
    sku        VARCHAR(50)  NOT NULL,
    name       VARCHAR(150) NOT NULL,
    category   VARCHAR(100) NOT NULL,
    price      BIGINT       NOT NULL DEFAULT 0,
    stock      INTEGER      NOT NULL DEFAULT 0,
    is_active  BOOLEAN      NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    CONSTRAINT products_sku_unique UNIQUE (sku),
    CONSTRAINT products_price_non_negative CHECK (price >= 0),
    CONSTRAINT products_stock_non_negative CHECK (stock >= 0)
);

-- ============================================================
-- Tabel transactions
-- ============================================================
CREATE TABLE IF NOT EXISTS transactions (
    id             BIGSERIAL PRIMARY KEY,
    user_id        BIGINT       NOT NULL,
    total_amount   BIGINT       NOT NULL DEFAULT 0,
    payment_method VARCHAR(50)  NOT NULL,
    status         VARCHAR(20)  NOT NULL DEFAULT 'success',
    created_at     TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    CONSTRAINT transactions_user_fk
        FOREIGN KEY (user_id) REFERENCES users (id)
        ON UPDATE CASCADE ON DELETE RESTRICT,
    CONSTRAINT transactions_status_check CHECK (status IN ('success', 'void')),
    CONSTRAINT transactions_total_non_negative CHECK (total_amount >= 0)
);

-- ============================================================
-- Tabel transaction_items
-- ============================================================
CREATE TABLE IF NOT EXISTS transaction_items (
    id             BIGSERIAL PRIMARY KEY,
    transaction_id BIGINT  NOT NULL,
    product_id     BIGINT  NOT NULL,
    quantity       INTEGER NOT NULL,
    subtotal       BIGINT  NOT NULL DEFAULT 0,
    CONSTRAINT transaction_items_transaction_fk
        FOREIGN KEY (transaction_id) REFERENCES transactions (id)
        ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT transaction_items_product_fk
        FOREIGN KEY (product_id) REFERENCES products (id)
        ON UPDATE CASCADE ON DELETE RESTRICT,
    CONSTRAINT transaction_items_quantity_positive CHECK (quantity > 0),
    CONSTRAINT transaction_items_subtotal_non_negative CHECK (subtotal >= 0)
);

-- ============================================================
-- Tabel refresh_tokens
-- ============================================================
CREATE TABLE IF NOT EXISTS refresh_tokens (
    id         BIGSERIAL PRIMARY KEY,
    user_id    BIGINT       NOT NULL,
    token_hash VARCHAR(255) NOT NULL,
    expires_at TIMESTAMPTZ  NOT NULL,
    CONSTRAINT refresh_tokens_user_fk
        FOREIGN KEY (user_id) REFERENCES users (id)
        ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT refresh_tokens_token_hash_unique UNIQUE (token_hash)
);

-- Index untuk keyset pagination dan query laporan
CREATE INDEX IF NOT EXISTS idx_products_id         ON products (id);
CREATE INDEX IF NOT EXISTS idx_products_category   ON products (category);
CREATE INDEX IF NOT EXISTS idx_transactions_user   ON transactions (user_id);
CREATE INDEX IF NOT EXISTS idx_transactions_created ON transactions (created_at);
CREATE INDEX IF NOT EXISTS idx_transaction_items_tx ON transaction_items (transaction_id);
CREATE INDEX IF NOT EXISTS idx_refresh_tokens_user ON refresh_tokens (user_id);

COMMIT;