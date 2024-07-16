# SingleStore Go Bookstore

**Attention**: The code in this repository is intended for experimental use only and is not fully tested, documented, or supported by SingleStore. Visit the [SingleStore Forums](https://www.singlestore.com/forum/) to ask questions about this repository.

This project demonstrates building a RESTful API for managing books using Golang and SingleStore. It follows test-driven development (TDD) principles and utilizes Testcontainers for efficient, reproducible testing.

## Requirements

- Go 1.22 or higher
- Docker

## Getting started

1. Set up the database:
   - Create a SingleStore database.
   - Create a `.env` file with the following environment variables:
     - `DB_USER`
     - `DB_PASSWORD`
     - `DB_HOST`
     - `DB_PORT`
     - `DB_NAME`

2. Run the application:
   ```
   go run main.go
   ```

The API will be available at `http://localhost:8080` (or the port specified in your configuration).

## API endpoints

- `GET /books`: Retrieve all books
- `GET /books/:id`: Retrieve a specific book
- `POST /books`: Create a new book
- `PUT /books/:id`: Update an existing book
- `DELETE /books/:id`: Delete a book

## Running tests

To run the unit tests:
```
$ make test
```
