# API made with Go, Hexagonal Architecture, Turso and ❤️

## 🚀 Getting Started

### Prerequisites

- Go 1.20+
- Turso account

### Installation

1. Clone the repository:

```bash
git clone https://github.com/turso/go-hexagonal-api-example.git
```

2. Install dependencies:

```bash
go mod download
```

3. Create a `.env` file in the root directory with the following content:

```bash
TURSO_DATABASE_URL=http://127.0.0.1:8080
TURSO_AUTH_TOKEN=your-auth-token
APP_PORT=3000
ENVIRONMENT=development
```

4. Run the application:

```bash
go run .
```

5.  Open your browser and navigate to `http://localhost:3000/api/users`.

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
