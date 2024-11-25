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

#### Using Docker

1. **Pull the Docker Image**:  
    Download the prebuilt Docker image from Docker Hub:  
    ```bash
    docker pull ikhsanhilmi/library-auth-service-app
    ```

2. **Run the Container**:  
    Start the container with the following command:  
    ```bash
    docker run -d \
      --name library-auth-service \
      -p 9090:9090 \
      -e POSTGRES_URL=postgresql://<user>:<password>@<host>:<port>/<dbname> \
      -e REDIS_URL=redis://<host>:<port> \
      -e JWT_SECRET=<your_jwt_secret> \
      ikhsanhilmi/library-auth-service-app
    ```
    Replace `<user>`, `<password>`, `<host>`, `<port>`, `<dbname>`, and `<your_jwt_secret>` with the appropriate credentials and secrets.

3. **Verify the Service**:  
    Access the service at `http://localhost:9090`.

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
