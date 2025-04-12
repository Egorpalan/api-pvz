CREATE INDEX idx_pvz_registration_date ON pvz(registration_date);

CREATE INDEX idx_receptions_pvz_status_date
    ON receptions(pvz_id, status, date_time DESC);

CREATE INDEX idx_products_reception_date
    ON products(reception_id, date_time DESC);

CREATE UNIQUE INDEX IF NOT EXISTS idx_users_email ON users(email);