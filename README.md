# Golang Authentication API

This project is a simple authentication API built using Golang, Gin framework, Gorm ORM, and PostgreSQL. It provides endpoints for user signup, login, and token validation, as well as user authorization middleware using cookies and jwt tokens.

## Features

- **Signup**: Allows users to create a new account by providing necessary details.
- **Login**: Validates user credentials and issues a JWT token upon successful login.
- **Token Validation**: Verifies the validity of JWT tokens to ensure secure access to protected routes.

## Prerequisites

- [Golang](https://golang.org/dl/)
- [Gin](https://github.com/gin-gonic/gin)
- [Gorm](https://gorm.io/)
- [JWT-Go](https://github.com/golang-jwt/jwt)
- [Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
- [Make](https://www.gnu.org/software/make/)
- [Docker](https://www.docker.com/) (Optional for containerization)

## Getting Started

1. **Clone the repository:**

   ```bash
    git clone https://github.com/mxnuchim/golang-auth-api.git
   ```

2. **Navigate to project directory:**

   ```bash
    cd golang-auth-api
   ```

3. **Install dependencies:**

   ```bash
    go mod tidy
   ```

4. **Set up environment variables:**

   ```bash
    PORT=8080
    DB_URL="<your-database-url>"
    JWT_SECRET=supersecretsecret
   ```

5. **Build and run the application:**

   ```bash
    make run
   ```

## Accessing the API

The API will be accessible at [http://localhost:8080](http://localhost:8080).

## API Endpoints

- **Signup**: `POST /api/signup`
- **Login**: `POST /api/login`
- **Validate Token**: `GET /api/validate` (Protected Route - Requires Authentication)

## Project Structure

- `main.go`: Entry point for the application.
- `controllers/`: Contains handlers for API endpoints.
- `initializers/`: Includes initialization code for database connection and migration.
- `.env`: Configuration file for environment variables.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
