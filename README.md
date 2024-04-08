# CQRS-Go Project

This project demonstrates a simple implementation of Command Query Responsibility Segregation (CQRS) in Go, using gRPC and the Ent ORM for database interactions.

## Prerequisites

Before running the project, ensure you have the following installed:

- Go (v1.16 or later)
- PostgreSQL database server
- protoc compiler for Protocol Buffers (v3 or later)

## Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/nidhey27/cqrs-go.git
   ```
   Navigate to the project directory:
   ```bash
   cd cqrs-go
   ```
2. Install dependencies:
   ```bash
   go mod download
   ```
3. Set up the PostgreSQL database:
   ```bash
   docker run --name event-sourcing-db -e POSTGRES_PASSWORD=password -e POSTGRES_USER=nidhey -e POSTGRES_DB=bank -d postgres:14
   ```
   Update the database connection details in main.go if necessary.
4. To start the gRPC server, run: (Runs the database migrations to create the schema)
   ```bash
   go run main.go
   ```
   The server will start listening on port 9000 by default.
