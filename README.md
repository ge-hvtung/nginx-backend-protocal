# My Go Backend

This is a backend project written in Go. It provides RESTful APIs for managing users and todos.

## Project Structure

The project has the following files and directories:

- `cmd/server/main.go`: Entry point of the application.
- `internal/config/config.go`: Contains the configuration options for the application.
- `internal/handlers/handlers.go`: Contains functions that handle the requests for the application.
- `internal/models/models.go`: Contains structs that represent the data models for the application.
- `internal/repositories/repositories.go`: Contains interfaces that define the methods for accessing the data models.
- `internal/routers/routers.go`: Contains a function that sets up the routes for the application.
- `internal/services/services.go`: Contains interfaces that define the methods for performing the necessary operations on the data models.
- `internal/utils/utils.go`: Contains utility functions that are used throughout the application.
- `pkg/mypkg/mypkg.go`: Contains utility functions that can be used by other applications.
- `vendor`: Contains the dependencies for the project.
- `go.mod`: Module definition file for the project.
- `go.sum`: Contains the checksums for the dependencies listed in `go.mod`.
- `README.md`: Documentation for the project.

## Getting Started

To run the application, follow these steps:

1. Clone the repository: `git clone https://github.com/your-username/my-go-backend.git`
2. Install the dependencies: `go mod download`
3. Start the server: `go run cmd/server/main.go`

The server will start listening on port 8080 by default.

## TODO
- [ ] :construction: Nginx Configuration Parsers
    + [ ] :construction: Server
    + [ ] :construction: Location
    + [ ] :construction: Error Pages

- new action

## Contributing

Contributions are welcome! Please read the [contributing guidelines](CONTRIBUTING.md) before submitting a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
