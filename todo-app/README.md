# To-Do App

This project is a To-Do application developed using PostgreSQL. It is written with a layered architecture and includes integration tests for each layer. All tests have passed successfully.

## Table of Contents
- [Features](#features)
- [Requirements](#requirements)
- [Setup](#setup)
- [Usage](#usage)
- [Testing](#testing)
- [Project Structure](#project-structure)
- [Contributing](#contributing)
- [License](#license)

## Features
- **Layered Architecture:** Separation of concerns with distinct layers for handling different aspects of the application.
- **PostgreSQL Integration:** Robust database integration for managing tasks.
- **Comprehensive Testing:** Integration tests for each layer ensuring reliability and correctness.

## Requirements
- [Go](https://golang.org/) 1.16+
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Setup
1. **Clone the repository:**
    ```bash
    git clone https://github.com/akaanuzman/GoLangTour.git
    cd GoLangTour/todo-app
    ```

2. **Initialize Docker and PostgreSQL:**
    Run the following script to set up the Docker containers and the PostgreSQL database:
    ```bash
    ./test/scripts/init_docker_and_db.sh
    ```

## Usage
1. **Start the application:**
    ```bash
    go run main.go
    ```

2. **Access the application:**
    Open your browser and go to `http://localhost:8080` to use the To-Do application.

## Testing
To run the integration tests, use the following command:
```bash
go test ./...
```

## Project Structure
```
todo-app/
├── cmd/                # Main applications of the project
├── internal/           # Private application and library code
│   ├── db/             # Database related code
│   ├── handler/        # HTTP handlers
│   ├── model/          # Data models
│   ├── repository/     # Data repositories
│   └── service/        # Business logic
├── test/               # Test scripts and test cases
├── Dockerfile          # Docker configuration
└── docker-compose.yml  # Docker Compose configuration
```