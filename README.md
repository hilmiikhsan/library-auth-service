# Library Auth Management

A scalable and efficient authentication and authorization service for the Library Management System. Built using **Golang**, **PostgreSQL**, and **Redis**, this microservice adheres to modern software architecture practices with features like gRPC communication, JWT-based security, and caching.

---

## Features

- **Authentication**: User login and token management using JWT.
- **Authorization**: Role-based access control (RBAC) for API endpoints.
- **Microservice Ready**: Designed to be integrated with other services via gRPC.
- **Caching**: Redis is used for session and token storage for improved performance.
- **Scalability**: Modular architecture with PostgreSQL and Redis for high availability.
- **Easy Migration**: Database migrations managed with `goose`.

---

## Technologies

- **Programming Language**: Golang 1.22.7
- **Database**: PostgreSQL
- **Cache**: Redis
- **gRPC**: Inter-service communication
- **JWT**: Authentication and authorization

---

## Getting Started

### Prerequisites

- Install [Golang](https://golang.org/dl/)
- Install [PostgreSQL](https://www.postgresql.org/download/)
- Install [Redis](https://redis.io/download)

---

### Setup Instructions

1. **Clone the repository**:
    ```bash
    git clone https://github.com/your-repo/library-auth-management.git
    cd library-auth-management
    ```

2. **Setup Environment Variables**:
    Copy the example environment file and adjust credentials:
    ```bash
    cp .env.example .env
    ```
    Modify `.env` with the appropriate database and Redis credentials.

3. **Run the Service**:
    Use the following command to start the service:
    ```bash
    make run
    ```
    The service will run on `http://localhost:9090` by default.

---

### API Documentation

API documentation is available at `http://localhost:9090/docs` once the service is running. It includes detailed information about the available endpoints, request/response formats, and error codes.

---

### Database Migrations

Manage database migrations using `goose`:

- Create a new migration:
    ```bash
    make goose-create name=create_users_table
    ```
- Apply migrations:
    ```bash
    make goose-up
    ```
- Rollback migrations:
    ```bash
    make goose-down
    ```
- Check migration status:
    ```bash
    make goose-status
    ```

---

### ERD (Entity-Relationship Diagram)

![ERD Diagram](./ERD/ERD.png)

---

### Development

#### Testing

Run unit tests with coverage:
```bash
make test
