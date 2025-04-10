CREATE TABLE users (
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
email TEXT UNIQUE NOT NULL,
password TEXT NOT NULL,
role TEXT NOT NULL CHECK (role IN ('client', 'moderator'))
);

CREATE TABLE pvz (
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
registration_date TIMESTAMPTZ NOT NULL DEFAULT now(),
city TEXT NOT NULL CHECK (city IN ('Москва', 'Санкт-Петербург', 'Казань'))
);

CREATE TABLE receptions (
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
pvz_id UUID NOT NULL REFERENCES pvz(id),
date_time TIMESTAMPTZ NOT NULL DEFAULT now(),
status TEXT NOT NULL CHECK (status IN ('in_progress', 'close'))
);

CREATE TABLE products (
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
reception_id UUID NOT NULL REFERENCES receptions(id),
date_time TIMESTAMPTZ NOT NULL DEFAULT now(),
type TEXT NOT NULL CHECK (type IN ('электроника', 'одежда', 'обувь'))
);