INSERT INTO categories (id, name, parent_id) 
VALUES
    ('65f4defe-9b54-4470-885d-059cf0956cd8', 'Elektronik', NULL),
    ('72a4ca20-821f-4ec8-a652-74c2eb3e1eb1', 'Handphone', '65f4defe-9b54-4470-885d-059cf0956cd8'),
    ('74b6017f-4f5a-466a-a476-373b8f25439f', 'Laptop', '65f4defe-9b54-4470-885d-059cf0956cd8'),
    ('711399ab-fcb4-4e01-8b31-0ae0f335ee8a', 'Fashion', NULL),
    ('9070135f-810c-4116-b776-5fd664275084', 'Baju Pria', '711399ab-fcb4-4e01-8b31-0ae0f335ee8a'),
    ('f73bf202-1887-443d-a97a-3689faa2ab8f', 'Baju Wanita', '711399ab-fcb4-4e01-8b31-0ae0f335ee8a'),
    ('43dc65f9-89e2-4e81-afcf-dc6ad6ed87dc', 'Kesehatan', NULL),
    ('4a0b7226-9a98-40f0-91a6-f6285b8d1f48', 'Makanan & Minuman', NULL),
    ('92ba769b-257a-48e9-b050-825ea94b9ff6', 'Olahraga', NULL)
ON CONFLICT (id) DO UPDATE SET name = EXCLUDED.name, parent_id = EXCLUDED.parent_id;