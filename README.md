Here's the README in English with your requested changes:

---

# PostgreSQL Studies with Go API

This is a simple API project developed in Go to study PostgreSQL implementation. The focus of the project is to learn about PostgreSQL integration in a Go API, using migrations, connections, and basic CRUD operations to simulate the day-to-day workflow of a real-world project.

## Project Structure

- **config**: Contains configuration files for connecting to the PostgreSQL database.
- **database**: Contains database migrations and general configurations.
- **middleware**: Includes middlewares like logging.
- **model**: Data model definitions (e.g., product model).
- **repository**: Contains repositories responsible for database operations.
- **service**: Service layer that holds the business logic.
- **test**: Contains test files to verify API operations, including example HTTP requests.

## Configuration

The project uses a `.env` file to manage environment variables. Be sure to populate this file with appropriate settings for your local environment. A sample `.env` file with default values for running in Docker containers can be found at the project's root.

## Makefile Commands

### Migrations

- **Create a new migration**:
  ```bash
  make migration name=migration_name
  ```
- **Apply migrations (up)**:
  ```bash
  make migration-up
  ```
- **Revert migrations (down)**:
  ```bash
  make migration-down
  ```

### Connect to PostgreSQL

To connect directly to the PostgreSQL database:
```bash
make pg-connect
```

(Note: This command requires the PostgreSQL CLI tool `psql`.)

### Run the Project

- **Start the project with Docker**:
  ```bash
  make start
  ```
  This command will build the Docker containers, start them, and apply the migrations.

- **Stop the project**:
  ```bash
  make stop
  ```
  This command will stop the Docker containers.

### Make Requests

To make HTTP or Hurl requests:

```bash
make request name=GET-v1-products.hurl
```

This command will execute the `GET-v1-products.hurl` file using `bat` to view the request body (a better alternative to `cat`), and `hurl` to check the API response.

(Note: Running requests via `make` commands requires the `Hurl` tool. To view the request body, we use `bat`, which is an alternative to `cat` but with more advanced features.)

## Running the Project Locally

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/studies-postgres-golang.git
   ```
2. Navigate to the project directory:
   ```bash
   cd studies-postgres-golang
   ```
3. Build and start the project with Docker:
   ```bash
   make start
   ```
4. Access the API at `http://localhost:8080`.

## Technologies

- Go
- PostgreSQL
- Docker / Docker Compose
- Makefile
- Hurl
- Bat
