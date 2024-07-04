# Hello World Birthday API

This is a simple "Hello World" API application written in Go using the Gin framework and bbolt for storage. The API allows saving/updating a user's name and date of birth, and retrieving a birthday message for the user.

## Features

- Save/update user's name and date of birth.
- Retrieve a birthday message based on the user's date of birth.

## Requirements

- Go 1.22.5 or later

## Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/hello-world-birthday-api.git
   cd hello-world-birthday-api
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Run the application:

   ```bash
   go run main.go
   ```

The server will run on `http://localhost:8080`.

## API Endpoints

### Save/Update User

- **Endpoint**: `PUT /hello/:username`
- **Description**: Saves/updates the given user’s name and date of birth in the database.
- **Request Parameters**:
  - `username`: The user's name (must contain only letters).
- **Request Body**:
  ```json
  {
    "dateOfBirth": "YYYY-MM-DD"
  }
  ```
  - `dateOfBirth`: The user's date of birth (must be a date before today).

- **Response**:
  - `204 No Content`

### Retrieve Birthday Message

- **Endpoint**: `GET /hello/:username`
- **Description**: Returns a hello birthday message for the given user.
- **Request Parameters**:
  - `username`: The user's name.
- **Response**:
  - `200 OK`
  - **Response Examples**:
    - If the user's birthday is in N days:
      ```json
      {
        "message": "Hello, <username>! Your birthday is in N day(s)"
      }
      ```
    - If the user's birthday is today:
      ```json
      {
        "message": "Hello, <username>! Happy birthday!"
      }
      ```
  - `404 Not Found`: If the user does not exist.

## Testing the API

### Using `curl`

1. **Save/Update User**:
   ```bash
   curl -X PUT http://localhost:8080/hello/johndoe -H "Content-Type: application/json" -d '{"dateOfBirth": "1990-01-01"}'
   ```

2. **Retrieve Birthday Message**:
   ```bash
   curl http://localhost:8080/hello/johndoe
   ```

### Using Postman

1. **Save/Update User**:
   - Create a new request.
   - Set the request type to `PUT`.
   - Enter the URL: `http://localhost:8080/hello/johndoe`.
   - Go to the `Body` tab, select `raw` and `JSON` format, then enter:
     ```json
     {
       "dateOfBirth": "1990-01-01"
     }
     ```
   - Click `Send`.

2. **Retrieve Birthday Message**:
   - Create a new request.
   - Set the request type to `GET`.
   - Enter the URL: `http://localhost:8080/hello/johndoe`.
   - Click `Send`.