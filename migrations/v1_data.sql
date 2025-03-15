INSERT INTO items (item_name, item_code, weight_kg, category, image, description, price_per_kg, cost_per_kg, customer_code) 
VALUES 
    ('Sample Product A', 'ITEM001', 2.50, 'Retail', NULL, 'This is a sample retail product.', 50.00, 30.00, NULL);

INSERT INTO raw_materials (stock_in, stock_out, stock_opname, stock_date, note)
VALUES
    (10, 0, NULL, '2025-03-01', 'First batch of materials'),
    (15, 0, NULL, '2025-03-03', 'Second batch of materials'),
    (20, 0, 18, '2025-03-05', 'Third batch (opname corrected to 18)'),
    (30, 0, NULL, '2025-03-07', 'Fourth batch of materials');


INSERT INTO productions (item_id, production_date, quantity)
VALUES
    (1, '2025-03-10', 12), -- Should take 10 from Batch 1 and 2 from Batch 2
    (1, '2025-03-11', 10), -- Should take 10 from remaining Batch 2
    (1, '2025-03-12', 20); -- Should take 18 from Batch 3 and 2 from Batch 4

INSERT INTO permissions (name, description) VALUES
    ('data_master_item_access', 'Access the item master menu'),
    ('data_master_item_create', 'Create new items'),
    ('data_master_item_edit', 'Edit existing items'),
    ('data_master_item_delete', 'Delete items'),
    
    ('data_master_customer_access', 'Access the customer master menu'),
    ('data_master_customer_create', 'Create new customers'),
    ('data_master_customer_edit', 'Edit existing customers'),
    ('data_master_customer_delete', 'Delete customers'),

    ('data_master_sales_access', 'Access the sales master menu'),
    ('data_master_sales_create', 'Create new sales records'),
    ('data_master_sales_edit', 'Edit sales records'),
    ('data_master_sales_delete', 'Delete sales records'),

    ('bahan_baku_access', 'Access raw materials'),
    ('bahan_baku_create', 'Create raw materials'),
    ('bahan_baku_edit', 'Edit raw materials'),
    ('bahan_baku_delete', 'Delete raw materials'),
    ('bahan_baku_opname', 'Stock opname for raw materials'),

    ('produksi_access', 'Access production'),
    ('produksi_create', 'Create production orders'),
    ('produksi_edit', 'Edit production orders'),
    ('produksi_delete', 'Delete production orders'),
    ('produksi_opname', 'Stock opname for production'),

    ('purchase_order_access', 'Access purchase orders'),
    ('purchase_order_create', 'Create purchase orders'),
    ('purchase_order_edit', 'Edit purchase orders'),
    ('purchase_order_delete', 'Delete purchase orders'),

    ('penjualan_access', 'Access sales transactions'),
    ('penjualan_create', 'Create sales transactions'),
    ('penjualan_edit', 'Edit sales transactions'),
    ('penjualan_delete', 'Delete sales transactions'),

    ('pengguna_access', 'Access user management'),
    ('pengguna_create', 'Create users'),
    ('pengguna_edit', 'Edit users'),
    ('pengguna_delete', 'Delete users');
