# API made with Go, Hexagonal Architecture, PostgreSQL and ❤️

## 🚀 Getting Started

### Prerequisites

- Go 1.20+
- PostgreSQL 14+

### Installation

1. Clone the repository:

```bash
git clone git@github.com:falvarador/golang-api.git
```

2. Install dependencies:

```bash
go mod download
```

3. Create a `.env` file in the root directory with the following content:

```bash
APP_PORT=3000
DB_CONNECTION_STRING="host=localhost port=5432 user=username password=secret_password dbname=database_name sslmode=disable"
ENVIRONMENT=development
```

4. Run the application:

```bash
go run .
```

5.  Open your browser and navigate to `http://localhost:3000/api/users`.

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
