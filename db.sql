create table
    products (
        id serial primary key,
        product_name varchar(255) not null,
        product_code varchar(255) not null,
        product_price varchar(255) not null
   );

INSERT INTO products (product_name, product_code, product_price) VALUES
('Apple iPhone 14', 'A001', '1299.00'),
('Samsung Galaxy S23', 'S002', '999.00'),
('Sony WH-1000XM5 Headphones', 'S003', '349.99'),
('Dell XPS 13', 'D004', '1199.00'),
('Apple MacBook Pro 16"', 'A005', '2399.00'),
('Nintendo Switch OLED', 'N006', '349.99'),
('Sony PlayStation 5', 'S007', '499.00'),
('Microsoft Xbox Series X', 'M008', '499.00'),
('Apple iPad Pro 12.9"', 'A009', '1099.00'),
('Samsung Galaxy Tab S8', 'S010', '799.00'),
('Bose QuietComfort 45', 'B011', '329.00'),
('Logitech MX Master 3 Mouse', 'L012', '99.99'),
('Canon EOS 90D DSLR Camera', 'C013', '1199.00'),
('Nikon Z9 Mirrorless Camera', 'N014', '5499.00'),
('GoPro HERO10 Black', 'G015', '499.00'),
('Fitbit Charge 5', 'F016', '149.95'),
('Apple Watch Series 8', 'A017', '399.00'),
('Microsoft Surface Laptop 4', 'M018', '1299.00'),
('JBL Charge 5 Bluetooth Speaker', 'J019', '179.95'),
('Samsung 65" QLED TV', 'S020', '1299.00');