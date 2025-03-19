CREATE TABLE "items" (
  "id" SERIAL PRIMARY KEY,
  "item_name" VARCHAR(255) NOT NULL,
  "item_code" VARCHAR(100) UNIQUE NOT NULL,
  "weight_g" NUMERIC(10,2) NOT NULL,
  "type" VARCHAR(50),
  "image" VARCHAR(255),
  "description" VARCHAR(255),
  "price_per_kg" NUMERIC(10,2) NOT NULL,
  "price_per_g" NUMERIC(10,2) NOT NULL,
  "price_per_unit" NUMERIC(10,2) NOT NULL,
  "cost_per_kg" NUMERIC(10,2) NOT NULL,
  "cost_per_g" NUMERIC(10,2) NOT NULL,
  "cost_per_unit" NUMERIC(10,2) NOT NULL,
  "customer_code" VARCHAR(100)
);

CREATE TABLE "salespeople" (
  "id" SERIAL PRIMARY KEY,
  "full_name" VARCHAR(255) NOT NULL,
  "phone_number" VARCHAR(15) UNIQUE,
  "address" VARCHAR(255),
  "ktp" VARCHAR(100) UNIQUE,
  "ktp_photo" VARCHAR(255),
  "profile_photo" VARCHAR(255)
);

CREATE TABLE "customers" (
  "id" SERIAL PRIMARY KEY,
  "customer_code" VARCHAR(100) UNIQUE,
  "full_name" VARCHAR(255) NOT NULL,
  "phone_number" VARCHAR(15) UNIQUE,
  "salesperson_id" INT REFERENCES "salespeople" ("id") ON DELETE SET NULL,
  "address" VARCHAR(255),
  "invoice_code" VARCHAR(100) UNIQUE,
  "email" VARCHAR(255) UNIQUE,
  "npwp_number" VARCHAR(255) UNIQUE,
  "npwp_photo" VARCHAR(255)
);

CREATE TABLE "transport_vehicles" (
  "id" SERIAL PRIMARY KEY,
  "driver_name" VARCHAR(255) NOT NULL,
  "vehicle_number" VARCHAR(100),
  "phone_number" VARCHAR(15) UNIQUE
);


CREATE TABLE "raw_materials" (
  "id" SERIAL PRIMARY KEY,
  "quantity" NUMERIC(10,2) DEFAULT 0,
  "type" VARCHAR(5),
  "note" VARCHAR(255),
  "stock_date" DATE NOT NULL DEFAULT CURRENT_DATE,
  "is_opname" BOOLEAN,
  "created_at" TIMESTAMP DEFAULT NOW(),
  "updated_at" TIMESTAMP DEFAULT NOW()
);

CREATE TABLE "productions" (
  "id" SERIAL PRIMARY KEY,
  "production_date" DATE NOT NULL DEFAULT CURRENT_DATE,
  "total_quantity" NUMERIC(10,2) NOT NULL,
  "created_at" TIMESTAMP DEFAULT NOW(),
  "updated_at" TIMESTAMP DEFAULT NOW()
);

CREATE OR REPLACE FUNCTION check_and_insert_production()
RETURNS TRIGGER AS $$
DECLARE
    total_in NUMERIC(10,2);
    total_out NUMERIC(10,2);
    available_stock NUMERIC(10,2);
BEGIN
    -- Calculate the total quantity of raw materials of type 'IN'
    SELECT COALESCE(SUM(quantity), 0) INTO total_in FROM raw_materials WHERE type = 'IN';

    -- Calculate the total quantity of raw materials of type 'OUT'
    SELECT COALESCE(SUM(quantity), 0) INTO total_out FROM raw_materials WHERE type = 'OUT';

    -- Compute available stock
    available_stock := total_in - total_out;

    -- Check if enough stock is available
    IF NEW.total_quantity > available_stock THEN
        RAISE EXCEPTION 'Not enough raw materials to create production. Available: %, Required: %', available_stock, NEW.total_quantity;
    END IF;

    -- Insert the new record into raw_materials as 'OUT'
    INSERT INTO raw_materials (quantity, type, stock_date, is_opname, created_at, updated_at)
    VALUES (NEW.total_quantity, 'OUT', NEW.production_date, FALSE, NOW(), NOW());

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER before_production_insert
BEFORE INSERT ON productions
FOR EACH ROW
EXECUTE FUNCTION check_and_insert_production();


CREATE TABLE "production_items" (
  "id" SERIAL PRIMARY KEY,
  "production_id" INT REFERENCES "productions" ("id") ON DELETE CASCADE,
  "item_id" INT REFERENCES "items" ("id") ON DELETE CASCADE,
  "quantity" NUMERIC(10,2) NOT NULL,
  "created_at" TIMESTAMP DEFAULT NOW(),
  "updated_at" TIMESTAMP DEFAULT NOW()
);

CREATE TABLE "purchase_orders" (
  "id" SERIAL PRIMARY KEY,
  "customer_id" INT NOT NULL REFERENCES "customers" ("id") ON DELETE CASCADE,
  "order_date" DATE NOT NULL DEFAULT CURRENT_DATE,
  "order_number" VARCHAR(255) UNIQUE NOT NULL,
  "tax" NUMERIC(10,2) NOT NULL DEFAULT 0,
  "price_total" NUMERIC(10,2) NOT NULL
);

CREATE TABLE "purchase_order_items" (
  "id" SERIAL PRIMARY KEY,
  "purchase_order_id" INT NOT NULL REFERENCES "purchase_orders" ("id") ON DELETE CASCADE,
  "item_id" INT REFERENCES "items" ("id") ON DELETE CASCADE,
  "quantity" INT NOT NULL,
  "cost_per_unit" NUMERIC(10,2) NOT NULL,
  "price_total" NUMERIC(10,2) NOT NULL,
  "remaining_quantity" NUMERIC(10,2) NOT NULL
);

CREATE TABLE "invoices" (
  "id" SERIAL PRIMARY KEY,
  "po_id" INT NOT NULL REFERENCES "purchase_orders" ("id") ON DELETE CASCADE,
  "sales_rep_id" INT REFERENCES "salespeople" ("id") ON DELETE SET NULL,
  "transport_vehicle_id" INT REFERENCES "transport_vehicles" ("id") ON DELETE SET NULL,
  "invoice_date" DATE NOT NULL DEFAULT CURRENT_DATE,
  "tax" NUMERIC(10,2) NOT NULL DEFAULT 0,
  "due_date" DATE NOT NULL,
  "due_days" INT,
  "price_total" NUMERIC(10,2) NOT NULL
);

CREATE TABLE "invoice_items" (
  "id" SERIAL PRIMARY KEY,
  "invoice_id" INT NOT NULL REFERENCES "invoices" ("id") ON DELETE CASCADE,
  "item_id" INT REFERENCES "items" ("id") ON DELETE CASCADE,
  "quantity" INT NOT NULL,
  "cost_per_unit" NUMERIC(10,2) NOT NULL,
  "price_total" NUMERIC(10,2) NOT NULL
);

CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "password" TEXT NOT NULL,
  "job_title" VARCHAR(100) NOT NULL,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "permissions" (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR(255) UNIQUE NOT NULL,
  "description" TEXT,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "user_permissions" (
  "user_id" INT NOT NULL REFERENCES "users" ("id") ON DELETE CASCADE,
  "permission_id" INT NOT NULL REFERENCES "permissions" ("id") ON DELETE CASCADE,
  PRIMARY KEY ("user_id", "permission_id")
);

-- INDEXES
CREATE INDEX "idx_items_code" ON "items" ("item_code");
CREATE INDEX "idx_stock_date" ON "raw_materials" ("stock_date");
CREATE INDEX "idx_production_id" ON "production_items" ("production_id");
CREATE INDEX "idx_item_id" ON "production_items" ("item_id");
CREATE INDEX "idx_po_customer_id" ON "purchase_orders" ("customer_id");
CREATE INDEX "idx_po_items_po_id" ON "purchase_order_items" ("purchase_order_id");
CREATE INDEX "idx_invoices_po_id" ON "invoices" ("po_id");
CREATE INDEX "idx_invoice_items_invoice_id" ON "invoice_items" ("invoice_id");
CREATE INDEX idx_raw_materials_type ON raw_materials (type);

