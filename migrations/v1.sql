CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    item_name VARCHAR(255) NOT NULL,
    item_code VARCHAR(100) UNIQUE NOT NULL,
    weight_kg DECIMAL(10,2) NOT NULL,
    category VARCHAR(50) CHECK (category IN ('Retail', 'Custom')) NOT NULL,
    image VARCHAR(255),
    description TEXT,
    price_per_kg DECIMAL(10,2) NOT NULL,
    price_per_unit DECIMAL(10,2),
    cost_per_kg DECIMAL(10,2) NOT NULL,
    cost_per_unit DECIMAL(10,2),
    customer_code VARCHAR(100) NULL
);

CREATE TABLE salespeople ( 
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    phone VARCHAR(15) UNIQUE, 
    address TEXT,
    ktp VARCHAR(100) UNIQUE, 
    ktp_photo VARCHAR(255),
    profile_photo VARCHAR(255)
);

CREATE TABLE customers ( 
    id SERIAL PRIMARY KEY,
    customer_code VARCHAR(100) UNIQUE, 
    full_name VARCHAR(255) NOT NULL,
    phone VARCHAR(15) UNIQUE,
    salesperson_id INT REFERENCES salespeople(id) ON DELETE SET NULL, 
    address TEXT,
    invoice_code VARCHAR(100) UNIQUE, 
    email VARCHAR(255) UNIQUE,
    npwp_number VARCHAR(255) UNIQUE, 
    npwp_photo VARCHAR(255) 
);


CREATE TABLE transport_vehicles (
    id SERIAL PRIMARY KEY,
    driver_name VARCHAR(255) NOT NULL,
    vehicle_number VARCHAR(100),
    phone_number VARCHAR(15) UNIQUE
);


-- Table for Raw Materials
CREATE TABLE raw_materials (
    id SERIAL PRIMARY KEY,
    stock_in NUMERIC(10,2) DEFAULT 0, -- Stock In (Masuk)
    stock_out NUMERIC(10,2) DEFAULT 0, -- Stock Out (Keluar)
    stock_opname NUMERIC(10,2) DEFAULT NULL, -- Stock Adjustment (Revisi)
    available_stock NUMERIC(10,2) GENERATED ALWAYS AS (COALESCE(stock_opname, stock_in) - stock_out) STORED,
    note TEXT,
    stock_date DATE NOT NULL DEFAULT CURRENT_DATE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Table for Productions
CREATE TABLE productions (
    id SERIAL PRIMARY KEY,
    item_id INT REFERENCES items(id) ON DELETE CASCADE,
    production_date DATE NOT NULL DEFAULT CURRENT_DATE,
    quantity NUMERIC(10,2) NOT NULL, 
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Indexes for Faster FIFO Processing
CREATE INDEX idx_stock_date ON raw_materials (stock_date);
CREATE INDEX idx_available_stock ON raw_materials (available_stock);
CREATE INDEX idx_item_id ON productions (item_id);

CREATE OR REPLACE FUNCTION allocate_fifo(production_id INT, production_qty NUMERIC(10,2))
RETURNS VOID AS $$
DECLARE
    remaining_qty NUMERIC(10,2) := production_qty;
    material RECORD;
BEGIN
    -- Fetch raw materials in FIFO order
    FOR material IN 
        SELECT id, available_stock
        FROM raw_materials
        WHERE available_stock > 0
        ORDER BY stock_date ASC, id ASC
    LOOP
        EXIT WHEN remaining_qty <= 0;

        -- Deduct stock
        UPDATE raw_materials
        SET stock_out = stock_out + LEAST(material.available_stock, remaining_qty),
            updated_at = NOW()
        WHERE id = material.id;

        -- Reduce remaining quantity
        remaining_qty := remaining_qty - LEAST(material.available_stock, remaining_qty);
    END LOOP;

    -- If there's still remaining_qty > 0, raise an error (stock insufficient)
    IF remaining_qty > 0 THEN
        RAISE EXCEPTION 'Not enough raw materials for production!';
    END IF;
END;
$$ LANGUAGE plpgsql;


CREATE OR REPLACE FUNCTION before_insert_production()
RETURNS TRIGGER AS $$
BEGIN
    PERFORM allocate_fifo(NEW.id, NEW.quantity);
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_allocate_fifo
AFTER INSERT ON productions
FOR EACH ROW
EXECUTE FUNCTION before_insert_production();


CREATE TABLE purchase_orders (
    id SERIAL PRIMARY KEY,
    customer_id INT NOT NULL REFERENCES customers(id) ON DELETE CASCADE,
    order_date DATE NOT NULL DEFAULT CURRENT_DATE,
    order_number VARCHAR(255) UNIQUE NOT NULL,
    tax DECIMAL(10,2) NOT NULL DEFAULT 0.00
);

-- Index for frequent lookups on customer orders
CREATE INDEX idx_purchase_orders_customer_id ON purchase_orders(customer_id);

CREATE TABLE purchase_order_items (
    id SERIAL PRIMARY KEY,
    purchase_order_id INT NOT NULL REFERENCES purchase_orders(id) ON DELETE CASCADE,
    item_id INT REFERENCES items(id) ON DELETE SET NULL,
    quantity INT NOT NULL CHECK (quantity > 0),
    price_per_unit DECIMAL(10,2) NOT NULL CHECK (price_per_unit >= 0),
    subtotal DECIMAL(10,2) GENERATED ALWAYS AS (price_per_unit * quantity) STORED
);

-- Index to speed up joins on purchase_order_id
CREATE INDEX idx_purchase_order_items_po_id ON purchase_order_items(purchase_order_id);

CREATE TABLE invoices (
    id SERIAL PRIMARY KEY,
    po_id INT NOT NULL REFERENCES purchase_orders(id) ON DELETE CASCADE,
    sales_rep_id INT REFERENCES salespeople(id) ON DELETE SET NULL,
    transport_vehicle_id INT REFERENCES transport_vehicles(id) ON DELETE SET NULL,
    invoice_date DATE NOT NULL DEFAULT CURRENT_DATE,
    tax DECIMAL(10,2) NOT NULL DEFAULT 0.00,
    due_date DATE NOT NULL,
    due_days INT DEFAULT NULL
);

-- Index for searching invoices by purchase order
CREATE INDEX idx_invoices_po_id ON invoices(po_id);

CREATE TABLE invoice_items (
    id SERIAL PRIMARY KEY,
    invoice_id INT NOT NULL REFERENCES invoices(id) ON DELETE CASCADE,
    item_id INT REFERENCES items(id) ON DELETE SET NULL,
    quantity INT NOT NULL CHECK (quantity > 0),
    price_per_unit DECIMAL(10,2) NOT NULL CHECK (price_per_unit >= 0),
    subtotal DECIMAL(10,2) GENERATED ALWAYS AS (price_per_unit * quantity) STORED
);

-- Index to optimize joins between invoices and items
CREATE INDEX idx_invoice_items_invoice_id ON invoice_items(invoice_id);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    password TEXT NOT NULL, 
    job_title VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE permissions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL, 
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE user_permissions (
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    permission_id INT NOT NULL REFERENCES permissions(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, permission_id)
);

-- Index to speed up permission lookups
CREATE INDEX idx_user_permissions_user_id ON user_permissions(user_id);
CREATE INDEX idx_user_permissions_permission_id ON user_permissions(permission_id);
