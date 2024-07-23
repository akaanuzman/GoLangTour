# To-Do App

This project is a To-Do application developed using PostgreSQL. It is written with a layered architecture and includes integration tests for each layer. All tests have passed successfully.

## Table of Contents
- [Features](#features)
- [Requirements](#requirements)
- [Setup](#setup)
- [Usage](#usage)
- [API Routes](#api-routes)
- [Testing](#testing)
- [Project Structure](#project-structure)

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

## API Routes
### Create a todo
- **URL:** `/api/v1/todo`
- **Method:** `POST`
- **Request Body:**
    ```json
    {
        "title": "Task title",
        "description": "Task description"
    }
    ```
- **Response:**
    ```json
    {
        "id": 1,
        "title": "Task title",
        "description": "Task description",
        "created_at": "2024-07-23T12:34:56Z",
        "updated_at": "2024-07-23T12:34:56Z"
    }
    ```

### Get All Todos
- **URL:** `/api/v1/todos`
- **Method:** `GET`
- **Response:**
    ```json
    [
        {
            "id": 1,
            "title": "Task title",
            "description": "Task description",
            "created_at": "2024-07-23T12:34:56Z",
            "updated_at": "2024-07-23T12:34:56Z",
            "isDone": false,
            "due_date": null
        },
        ...
    ]
    ```

### Get a Single todo
- **URL:** `/api/v1/todos/{id}`
- **Method:** `GET`
- **Response:**
    ```json
    {
        "id": 1,
        "title": "Task title",
        "description": "Task description",
        "created_at": "2024-07-23T12:34:56Z",
        "updated_at": "2024-07-23T12:34:56Z",
        "isDone": false,
        "due_date": null
    }
    ```

### Get a done todos
- **URL:** `/todos/done?isDone=true`
- **Method:** `GET`
- **Response:**
    ```json
        [
            {
                "id": 1,
                "title": "Task title",
                "description": "Task description",
                "created_at": "2024-07-23T12:34:56Z",
                "updated_at": "2024-07-23T12:34:56Z",
                "isDone": true,
                "due_date": "2024-07-24T12:34:56Z"
            }
            ...
        ]
    ```

### Get a undone todos
- **URL:** `/todos/done?isDone=false`
- **Method:** `GET`
- **Response:**
    ```json
        [
            {
                "id": 1,
                "title": "Task title",
                "description": "Task description",
                "created_at": "2024-07-23T12:34:56Z",
                "updated_at": "2024-07-23T12:34:56Z",
                "isDone": false,
                "due_date": null
            }
            ...
        ]
    ```

### Sign done a Task
- **URL:** `/api/v1/tasks/{id}/done`
- **Method:** `PUT`
- **Request Body:**
    ```json
    {
        "isDone": true,
        "dueDate": "2024-07-30T15:04:05Z"
    }
    ```
- **Response:**
    ```json

    ```

### Sign undone a Task
- **URL:** `/api/v1/tasks/{id}/undone`
- **Method:** `PUT`
- **Response:**
    ```json

    ```

### Delete a Task
- **URL:** `/api/v1/tasks/{id}`
- **Method:** `DELETE`
- **Response:**
    ```json

    ```

## Testing
To run the integration tests, use the following command:
```bash
go test ./...
```

## Project Structure
```
todo-app/
├── common/             # Database configurations
│   ├── app/            # Package app provides the core application configurations and utilities.
│   └── postresql/      # Package postresql provides utilities for working with PostgreSQL databases, including connection pool management.
├── controller/         # Project controller layer
│   ├── request/        # Request struct
│   └── response/       # Response struct
├── domains/            # Data model
├── persistance/        # Project persistance manager layer
├── service/            # Project service layer
│   └── model/          # Service data model
├── test/               # Test scripts and test cases
├── Dockerfile          # Docker configuration
└── docker-compose.yml  # Docker Compose configuration
```