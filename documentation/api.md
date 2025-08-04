# API Documentation

## Base URL
```
http://localhost:8080/api/v1
```

## Endpoints

### Health Check
**GET** `/health`

Returns the health status of the API.

**Response:**
```json
{
  "status": "ok"
}
```

### Users

#### Create User
**POST** `/users`

Creates a new user.

**Request Body:**
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "securepassword"
}
```

**Response:**
```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com",
  "created_at": "2025-08-01T10:00:00Z",
  "updated_at": "2025-08-01T10:00:00Z"
}
```

#### Get User
**GET** `/users/{id}`

Retrieves a user by ID.

**Response:**
```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com",
  "created_at": "2025-08-01T10:00:00Z",
  "updated_at": "2025-08-01T10:00:00Z"
}
```

#### Update User
**PUT** `/users/{id}`

Updates an existing user.

**Request Body:**
```json
{
  "name": "John Smith",
  "email": "johnsmith@example.com"
}
```

**Response:**
```json
{
  "id": 1,
  "name": "John Smith",
  "email": "johnsmith@example.com",
  "created_at": "2025-08-01T10:00:00Z",
  "updated_at": "2025-08-01T10:01:00Z"
}
```

#### Delete User
**DELETE** `/users/{id}`

Deletes a user by ID.

**Response:**
```json
{
  "message": "user deleted successfully"
}
```

#### List Users
**GET** `/users?limit=10&offset=0`

Retrieves a paginated list of users.

**Query Parameters:**
- `limit` (optional): Number of users to return (default: 10)
- `offset` (optional): Number of users to skip (default: 0)

**Response:**
```json
{
  "users": [
    {
      "id": 1,
      "name": "John Doe",
      "email": "john@example.com",
      "created_at": "2025-08-01T10:00:00Z",
      "updated_at": "2025-08-01T10:00:00Z"
    }
  ]
}
```

## Error Responses

All error responses follow this format:

```json
{
  "error": "Error message description"
}
```

**HTTP Status Codes:**
- `200` - Success
- `201` - Created
- `400` - Bad Request
- `404` - Not Found
- `500` - Internal Server Error
