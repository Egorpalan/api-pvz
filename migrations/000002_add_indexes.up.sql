-- Индексы для pvz
CREATE INDEX idx_pvz_registration_date ON pvz(registration_date);

-- Индексы для receptions
CREATE INDEX idx_receptions_pvz_status_date
    ON receptions(pvz_id, status, date_time DESC);

-- Индексы для products
CREATE INDEX idx_products_reception_date
    ON products(reception_id, date_time DESC);

-- Индекс на users (если не было)
CREATE UNIQUE INDEX IF NOT EXISTS idx_users_email ON users(email);