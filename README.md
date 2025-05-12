# Golang Hexagolan API Boilerplate

A modern, high-performance backend server built with Go, featuring a robust architecture and essential microservices components.

## 🚀 Features

- **High Performance**: Built with Go for optimal performance and low latency
- **Modern Stack**: Utilizes Gin framework for HTTP routing and middleware
- **Database Support**: PostgreSQL integration with migration support
- **Caching**: Redis integration for high-speed caching
- **Security**: PASETO token authentication
- **Logging**: Structured logging with slog
- **API Documentation**: OpenAPI/Swagger support
- **Environment Configuration**: Dotenv support for flexible configuration
- **Database Migrations**: Built-in migration support using golang-migrate

## 📋 TODO List

### High Priority

- [ ] Add rate limiting middleware
- [ ] Implement request validation using validator
- [ ] Add health check endpoints
- [ ] Set up CI/CD pipeline with GitHub Actions
- [ ] Add unit tests with high coverage
- [ ] Implement graceful shutdown

### Medium Priority

- [ ] Add metrics collection (Prometheus)
- [ ] Implement distributed tracing
- [ ] Add API versioning support
- [ ] Create Docker and Docker Compose setup
- [ ] Add database backup and restore functionality
- [ ] Implement caching strategies

### Low Priority

- [ ] Add GraphQL support
- [ ] Implement WebSocket support
- [ ] Add support for multiple database types
- [ ] Create admin dashboard
- [ ] Add support for multiple authentication methods
- [ ] Implement audit logging

### Documentation

- [ ] Add API documentation with examples
- [ ] Create architecture diagrams
- [ ] Add contribution guidelines
- [ ] Create troubleshooting guide
- [ ] Add performance benchmarks
- [ ] Document deployment strategies

## 🛠️ Tech Stack

- **Framework**: Gin
- **Database**: PostgreSQL
- **Cache**: Redis
- **Authentication**: PASETO
- **Logging**: slog
- **Configuration**: godotenv
- **Database Migrations**: golang-migrate
- **API Documentation**: OpenAPI/Swagger

## 📋 Prerequisites

- Go 1.23.4 or higher
- PostgreSQL
- Redis
- Make (for build commands)

## 🚀 Getting Started

1. Clone the repository:

   ```bash
   git clone https://github.com/brkss/dextrace-server.git
   cd dextrace-server
   ```

2. Install dependencies:

   ```bash
   go mod download
   ```

3. Set up your environment variables:

   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

4. Start the development database:

   ```bash
   ./run-dbs.sh
   ```

5. Run database migrations:

   ```bash
   make migrate-up
   ```

6. Build and run the server:
   ```bash
   make
   ./bin/dextrace
   ```

## 🏗️ Project Structure

```
.
├── cmd/            # Application entry points
├── internal/       # Private application code
├── bin/           # Compiled binaries
├── log/           # Log files
├── go.mod         # Go module file
├── go.sum         # Go module checksum
├── Makefile       # Build commands
└── run-dbs.sh     # Database setup script
```

## 🛠️ Development

### Creating a New Migration

```bash
make create-migration name=migration_name
```

### Building the Project

```bash
make
```

## 📝 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🤝 Contributing

Contributions, issues, and feature requests are welcome! Feel free to check the [issues page](https://github.com/brkss/dextrace-server/issues).

## 📫 Contact

For any questions or concerns, please open an issue in the GitHub repository.
