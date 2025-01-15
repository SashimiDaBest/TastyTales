# TastyTales

http://localhost:8080/

bun run start-backend
bun run build-backend
bun run test-backend
bun run lint
bun run clean

go mod tidy

```
project/
├── cmd/
│   ├── server/
│   └── main.go          # Entry point for the server application
├── internal/
│   ├── handlers/        # Code containing HTTP handlers (also called controllers)
│   ├── models/          # Code containing data models that represent data structures
│   ├── repositories/    # Code for database interactions
│   └── services/        # Code containing business logic
├── static/              # Static files (CSS, JS, images)
├── templates/           # HTML templates (if using Go's template engine)
│   ├── [page_name].html
│   └── [page_name].html
├── go.mod
├── package.json
└── README.md


├── main.go              # Main server file
├── static/              # Static files (CSS, JS, images)
├── templates/           # HTML templates (if using Go's template engine)
│   ├── layout.html
│   └── index.html
├── handlers/            # Route handlers
│   ├── hello.go
│   └── form.go
└── go.mod               # Go modules file
```