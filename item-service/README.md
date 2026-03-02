# Item Service - Microservices Lab

A RESTful Item Service built with Node.js and Express for managing inventory items.

## Features

- ✅ RESTful API with CRUD operations
- ✅ In-memory data storage
- ✅ Input validation
- ✅ Error handling
- ✅ CORS enabled
- ✅ Health check endpoint

## Project Structure

```
backend/
├── controllers/
│   └── itemController.js    # Business logic for item operations
├── models/
│   └── itemModel.js         # Item data model and in-memory storage
├── routes/
│   └── itemRoutes.js        # API route definitions
├── .env                     # Environment variables (not in git)
├── .env.example             # Example environment variables
├── .gitignore              # Git ignore file
├── package.json            # Project dependencies
└── server.js               # Express server configuration
```

## Prerequisites

- Node.js (v14 or higher)
- npm (Node Package Manager)

## Installation

1. Navigate to the backend directory:
```bash
cd backend
```

2. Install dependencies:
```bash
npm install
```

3. Configure environment variables (optional):
```bash
cp .env.example .env
```
Edit `.env` to change the port if needed (default: 3001)

## Running the Service

### Development mode (with auto-restart):
```bash
npm run dev
```

### Production mode:
```bash
npm start
```

The service will start on `http://localhost:3001`

## API Endpoints

### Health Check
- **GET** `/health`
  - Returns service health status

### Root
- **GET** `/`
  - Returns welcome message and available endpoints

### Items

#### Get all items
- **GET** `/api/items`
- **Response**: Array of all items

#### Get item by ID
- **GET** `/api/items/:id`
- **Response**: Single item object

#### Create new item
- **POST** `/api/items`
- **Body**:
```json
{
  "name": "Item Name",
  "description": "Item description (optional)",
  "price": 99.99,
  "quantity": 10
}
```
- **Response**: Created item object

#### Update item
- **PUT** `/api/items/:id`
- **Body**:
```json
{
  "name": "Updated Name",
  "description": "Updated description",
  "price": 149.99,
  "quantity": 5
}
```
- **Response**: Updated item object

#### Delete item
- **DELETE** `/api/items/:id`
- **Response**: Success message

## Example API Calls

### Using cURL

```bash
# Get all items
curl http://localhost:3001/api/items

# Get item by ID
curl http://localhost:3001/api/items/1

# Create new item
curl -X POST http://localhost:3001/api/items \
  -H "Content-Type: application/json" \
  -d '{"name":"Keyboard","description":"Mechanical keyboard","price":89.99,"quantity":25}'

# Update item
curl -X PUT http://localhost:3001/api/items/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Updated Laptop","price":1499.99}'

# Delete item
curl -X DELETE http://localhost:3001/api/items/1
```

### Using PowerShell

```powershell
# Get all items
Invoke-RestMethod -Uri "http://localhost:3001/api/items" -Method Get

# Create new item
$body = @{
    name = "Keyboard"
    description = "Mechanical keyboard"
    price = 89.99
    quantity = 25
} | ConvertTo-Json

Invoke-RestMethod -Uri "http://localhost:3001/api/items" -Method Post -Body $body -ContentType "application/json"
```

## Response Format

### Success Response
```json
{
  "success": true,
  "data": { ... }
}
```

### Error Response
```json
{
  "success": false,
  "error": "Error message"
}
```

## Sample Data

The service comes with 3 pre-loaded items:
1. Laptop - $1299.99
2. Wireless Mouse - $29.99
3. USB-C Cable - $14.99

## Technologies Used

- **Node.js** - JavaScript runtime
- **Express** - Web framework
- **CORS** - Cross-Origin Resource Sharing
- **dotenv** - Environment configuration
- **body-parser** - Request body parsing

## Development

To modify the service:
1. Controllers are in `controllers/itemController.js`
2. Routes are defined in `routes/itemRoutes.js`
3. Data model is in `models/itemModel.js`
4. Server configuration is in `server.js`

## Notes

- This service uses in-memory storage. Data will be lost when the service restarts.
- For production use, consider integrating a database (MongoDB, PostgreSQL, etc.)
- Add authentication/authorization for secure access
- Implement logging for better debugging

## License

ISC
