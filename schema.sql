-- Create Database
CREATE DATABASE IF NOT EXISTS inventory_db;
USE inventory_db;

-- Drop tables if they already exist
DROP TABLE IF EXISTS purchase_orders;
DROP TABLE IF EXISTS sales_orders;
DROP TABLE IF EXISTS products;

-- Create Products Table
CREATE TABLE products (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    price DECIMAL(10,2) NOT NULL,
    quantity INT NOT NULL
    
);

-- Create Sales Orders Table
CREATE TABLE sales_orders (
    id INT AUTO_INCREMENT PRIMARY KEY,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    total_price DECIMAL(10,2) NOT NULL,
    order_date DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);

-- Create Purchase Orders Table
CREATE TABLE purchase_orders (
    id INT AUTO_INCREMENT PRIMARY KEY,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    supplier VARCHAR(100),
    order_date DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);

INSERT INTO products (name, description, price, quantity)
VALUES 
('Mixer', 'multitasking', 1200.00, 10),
('Phone', '4 GB RAM', 200.00, 15),
('Keyboard', 'Mechanical Keyboard', 75.00, 25);

-- Sample Sales Orders
INSERT INTO sales_orders (product_id, quantity, total_price)
VALUES 
(1, 2, 2400.00),
(2, 1, 200.00);

-- Sample Purchase Orders
INSERT INTO purchase_orders (product_id, quantity, supplier)
VALUES 
(1, 5, 'FAST AND FORWARD'),
(3, 10, 'Peripheral Inc.');