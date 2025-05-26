  Inventory Management System – README

1.	Project Overview
The Inventory Management System is a RESTful API developed using Go Lang, MySQL, and
GORM.
It enables full CRUD operations for managing products, sales orders, and purchase orders, with automatic inventory updates based on stock movement.

2.	Tech Stack
•	Programming Language: Go (Golang)
•	Database: MySQL
•	ORM: GORM
•	Testing: Go’s built-in testing framework

3.	Setup Instructions
Step 1: Clone the Repository
bash
CopyEdit
git clone https://github.com/yourusername/inventory-system.git cd inventory-system

Step 2: Configure the Database
•	Create a new MySQL database named inventory_db.
•	Update config/database.go with your DB credentials.

Step 3: Install Dependencies
bash
CopyEdit
go mod tidy
Step 4: Run the Server
 
bash
CopyEdit
go run main.go
The server will start on: http://localhost:8080

4.	API Endpoints
◆	Product Management

Method	Endpoint	      Description
GET	    /products     	Retrieve all products
GET	    /products/:id	  Retrieve product by ID
POST	  /products	      Create a new product
PUT   	/products/:id  	Update existing product
DELETE	/products/:id	  Delete a product

◆	Sales Orders

Method	 Endpoint 	Description
GET   	/sales	    List all sales orders
POST	  /sales	    Create a sales order

◆	Purchase Orders

Method	 Endpoint	   Description
GET	    /purchases	 List all purchase orders
POST	  /purchases	 Create a purchase order

◆	Inventory

Method	Endpoint	   Description
GET   	/inventory	 Get current inventor
 
5.	Sample cURL Requests
➤ Create a Product
bash
CopyEdit
curl -X POST http://localhost:8080/products \
-H "Content-Type: application/json" \
-d '{"name":"Laptop","description":"Gaming","price":1200,"quantity":10}'

➤ Get All Products
bash
CopyEdit
curl http://localhost:8080/products

➤ Create a Sales Order
bash
CopyEdit
curl -X POST http://localhost:8080/sales \
-H "Content-Type: application/json" \
-d '{"product_id":1,"quantity":2}'

➤ Create a Purchase Order
bash
CopyEdit
curl -X POST http://localhost:8080/purchases \
-H "Content-Type: application/json" \
-d '{"product_id":1,"quantity":5,"supplier":"Tech Supplies Co."}'

6.	Postman Usage
You can test the endpoints using Postman:
•	Set {{base_url}} to http://localhost:8080
 
•	Use raw JSON for POST and PUT requests
Example POST /products Body:
json
CopyEdit
{
"name": "Monitor",
"description": "24-inch LED Monitor", "price": 200,
"quantity": 15
}

7.	Running Tests
Execute all unit tests using: bash
CopyEdit
go test ./tests/...

8.	Project Structure
perl
CopyEdit
inventory-system/
├── config/	# Database configuration
├── controllers/	# API request handlers
├── models/	# GORM models
├── routes/	# API route definitions
├── services/	# Business logic layer
├── tests/	# Unit tests
 
├── main.go	# Entry point
└── README.md	# Project documentation
