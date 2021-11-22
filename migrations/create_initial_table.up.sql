DROP TABLE IF EXISTS customer CASCADE;

DROP TABLE IF EXISTS services CASCADE;

DROP TABLE IF EXISTS technician CASCADE;

DROP TABLE IF EXISTS order CASCADE;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE EXTENSION IF NOT EXISTS CITEXT;

CREATE TABLE customer (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(200) NOT NULL CHECK (name <> ''),
    email VARCHAR(64) UNIQUE NOT NULL CHECK (email <> ''),
    msisdn VARCHAR(20),
    address VARCHAR(250),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE services (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    damage_type VARCHAR(250) NOT NULL CHECK (damage_type <> ''),
    fix_duration INT NOT NULL,
    fee INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE technician (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(250) NOT NULL CHECK (name <> ''),
    brand_specialist NOT NULL CHECK (brand_specialist <> ''),
    platform NOT NULL CHECK (platform <> ''),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE order (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    customer_id UUID NOT NULL REFERENCES customer (id) ON DELETE CASCADE,
    services_id UUID NOT NULL REFERENCES services (id) ON DELETE CASCADE,
    technician_id UUID NOT NULL REFERENCES technician (id) ON DELETE CASCADE,
    brand NOT NULL CHECK (brand <> ''),
    technician_name NOT NULL CHECK (technician_name <> ''),
    damage_type NOT NULL CHECK (damage_type <> ''),
    description VARCHAR(1024) NOT NULL CHECK (description <> ''),
    status INT DEFAULT 1,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS order_technician_id_idx ON order (technician_id);