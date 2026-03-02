🚀 Microservices Lab – SE4010

Build · Dockerize · Deploy · Test

A complete containerized microservices system built for the SE4010 – Current Trends in Software Engineering module.

This project demonstrates how independent services communicate through an API Gateway, run inside Docker containers, and are orchestrated using Docker Compose.

📌 Project Overview

This system consists of four independent components:

Service	Port	Description
Item Service	8081	Manages product items
Order Service	8082	Handles customer orders
Payment Service	8083	Processes payments
API Gateway	8080	Single entry point for all services

All services are containerized and communicate over a shared Docker network.

🏗️ System Architecture
Client (Postman / Browser)
        ↓
   API Gateway :8080
  /items /orders /payments
        ↓
 ┌────────────┬────────────┬────────────┐
 │            │            │            │
Item :8081  Order :8082  Payment :8083
🔁 Routing Rules
Incoming Request	Routed To
/items/**	Item Service
/orders/**	Order Service
/payments/**	Payment Service

The API Gateway routes requests using path-based routing.

🛠️ Technology Stack

This project follows a polyglot microservices approach.

Component	Technology Used
Item Service	(e.g., Spring Boot / Node.js / FastAPI)
Order Service	(e.g., Spring Boot / Gin / Django REST)
Payment Service	(e.g., Spring Boot / Flask / .NET Core)
API Gateway	Spring Cloud Gateway / NGINX / Kong

Each service runs independently inside its own Docker container.

📂 Project Structure
microservices-lab/
│
├── item-service/
│   ├── Dockerfile
│   └── source code
│
├── order-service/
│   ├── Dockerfile
│   └── source code
│
├── payment-service/
│   ├── Dockerfile
│   └── source code
│
├── api-gateway/
│   ├── Dockerfile
│   └── configuration
│
└── docker-compose.yml
🐳 Docker Setup
1️⃣ Build All Services
docker-compose build
2️⃣ Start All Services
docker-compose up
3️⃣ Run in Background
docker-compose up -d
4️⃣ Stop All Containers
docker-compose down
🌐 Accessing the Application

All requests must go through the API Gateway:

http://localhost:8080
📡 API Endpoints
📦 Item Service
Method	Endpoint	Description
GET	/items	Get all items
POST	/items	Add new item
GET	/items/{id}	Get item by ID

Example:

POST /items
{
  "name": "Headphones"
}
🛒 Order Service
Method	Endpoint	Description
GET	/orders	Get all orders
POST	/orders	Create new order
GET	/orders/{id}	Get order by ID

Example:

POST /orders
{
  "item": "Laptop",
  "quantity": 2,
  "customerId": "C001"
}
💳 Payment Service
Method	Endpoint	Description
GET	/payments	Get all payments
POST	/payments/process	Process payment
GET	/payments/{id}	Get payment by ID

Example:

POST /payments/process
{
  "orderId": 1,
  "amount": 1299.99,
  "method": "CARD"
}
🧪 Testing

The system was tested using:

✅ Postman Collection

✅ Gateway-based routing validation

✅ Multi-container integration testing

To test:

Start Docker containers

Open Postman

Send requests to http://localhost:8080

🔐 Networking

All services are connected via a shared Docker bridge network:

microservices-net

Services communicate internally using service names (e.g., item-service) instead of localhost.

🎯 Learning Outcomes Achieved

Implemented microservices architecture

Designed RESTful APIs

Configured API Gateway routing

Containerized services using Docker

Orchestrated multi-service environment with Docker Compose

Tested distributed system using Postman

📦 Requirements

Docker

Docker Compose

Git

Postman (for testing)

👨‍💻 Author

Chamuditha Subharath
SE4010 – Microservices Lab
SLIIT – Faculty of Computing

📜 License

This project is created for academic purposes.

