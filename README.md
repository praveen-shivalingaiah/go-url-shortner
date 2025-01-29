# go-url-shortner
Implementing URL Shortner in golang

About Hexagonal Architecture : 
Hexagonal Architecture (Ports and Adapters) emphasizes separating your business logic (domain) from outside concerns (adapters):
Domain (Core / Business Logic):
Represents the heart of the application.
Contains entities, use cases (or services), and interfaces that describe what the application does.
Ports (Interfaces):
Define how the domain interacts with the outside world (e.g., Repository interface for data persistence, a ShortenerService interface for the application’s core service).
Ports are how external systems (database, HTTP) “plug into” your domain.
Adapters (Implementations):
Concrete implementations of the ports/interfaces (e.g., a Postgres adapter for data persistence, an HTTP server adapter for receiving requests).
Application Boundaries:
The application core knows nothing about infrastructure details.
Infrastructure details (like a database driver) implement the interfaces defined in the domain.
Structure of Application : 

url-shortener/
├── cmd/
│   └── shortener/
│       └── main.go            // Application entry point
├── internal/
│   ├── domain/
│   │   └── url.go             // Entities and interfaces
│   ├── app/
│   │   └── shortener.go       // Core business logic (use cases)
│   ├── adapters/
│   │   ├── http/
│   │   │   └── handler.go     // HTTP server/router
│   │   └── storage/
│   │       ├── memory.go      // In-memory storage adapter
│   │       └── postgres.go    // Postgres storage adapter (if desired)
├── go.mod
├── go.sum
└── README.md

The following Repo consists of the source code: https://github.com/praveen-shivalingaiah/go-url-shortner