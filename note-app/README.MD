# Note App REST API

## Overview

This project is a REST API for a note application built with Go using the Echo framework and MongoDB. It includes user authentication, note creation, updating, deletion, and retrieval functionalities.

## Features

- **User Authentication**: Register and login functionalities with JWT-based authentication.
- **Note Management**: Create, update, delete, and retrieve notes.
- **JWT Middleware**: Protect endpoints with JWT-based authorization.

## Getting Started

### Prerequisites

- Go (1.18 or higher)
- MongoDB

### Installation

1. **Clone the Repository**

    ```sh
    git clone https://github.com/akaanuzman/GoLangTour.git
    cd GoLangTour/note-app/note-app
    ```

2. **Install Dependencies**

    Navigate to the `note-app` directory and install the required dependencies:

    ```sh
    go mod tidy
    ```

3. **Setup MongoDB**

    You can create a `.env` file in the root directory and create the MongoDB URI.

    **Create `.env` file in the root directory:**

    * Set `MONGO_URI = <YOUR_MONGO_URI>`
    * Set `MAIL=<YOUR_MAIL>`
    * Set `MAIL_PORT=<YOUR_MAIL_PORT>`
    * Set `MAIL_PASSWORD=<YOUR_MAIL_PASSWORD>`
    * Set `MAIL_HOST=<YOUR_MAIL_HOST>`

4. **Run the Application**

    To run the application, use the following command:

    ```sh
    go run main.go
    ```

    The API will be available at `http://localhost:8080`.

## API Endpoints

### Authentication Endpoints

- **Register User**

    - **URL**: `/api/v1/auth/register`
    - **Method**: `POST`
    - **Body**:

      ```json
      {
        "email": "user@example.com",
        "password": "yourpassword"
      }
      ```

    - **Response**:

      ```json
      {
        "email": "user@example.com",
        "password": "hashedpassword"
      }
      ```

- **Login User**

    - **URL**: `/api/v1/auth/login`
    - **Method**: `POST`
    - **Body**:

      ```json
      {
        "email": "user@example.com",
        "password": "yourpassword"
      }
      ```

    - **Response**:

      ```json
      {
        "token": "your_jwt_token",
        "user": {
          "email": "user@example.com"
        }
      }
      ```

### Note Endpoints

- **Create Note**

    - **URL**: `/api/v1/notes`
    - **Method**: `POST`
    - **Headers**: `Authorization: Bearer your_jwt_token`
    - **Body**:

      ```json
      {
        "title": "My First Note",
        "content": "This is the content of the note.",
        "tags": ["tag1", "tag2"]
      }
      ```

    - **Response**:

      ```json
      {
        "id": "note_id",
        "title": "My First Note",
        "content": "This is the content of the note.",
        "authorId": "user_id",
        "createdAt": "2024-07-25T00:00:00Z",
        "updatedAt": "2024-07-25T00:00:00Z",
        "tags": ["tag1", "tag2"],
        "isActive": true
      }
      ```

- **Update Note**

    - **URL**: `/api/v1/notes/:id`
    - **Method**: `PUT`
    - **Headers**: `Authorization: Bearer your_jwt_token`
    - **Body**:

      ```json
      {
        "title": "Updated Title",
        "content": "Updated content.",
        "tags": ["tag3", "tag4"]
      }
      ```

    - **Response**:

      ```json
      {
        "id": "note_id",
        "title": "Updated Title",
        "content": "Updated content.",
        "authorId": "user_id",
        "createdAt": "2024-07-25T00:00:00Z",
        "updatedAt": "2024-07-25T01:00:00Z",
        "tags": ["tag3", "tag4"],
        "isActive": true
      }
      ```

- **Delete Note**

    - **URL**: `/api/v1/notes/:id`
    - **Method**: `DELETE`
    - **Headers**: `Authorization: Bearer your_jwt_token`
    - **Response**:

      ```json
      {
        "message": "Note deleted"
      }
      ```

- **Get All Notes**

    - **URL**: `/api/v1/notes`
    - **Method**: `GET`
    - **Headers**: `Authorization: Bearer your_jwt_token`
    - **Response**:

      ```json
      [
        {
          "id": "note_id",
          "title": "Note Title",
          "content": "Note content.",
          "authorId": "user_id",
          "createdAt": "2024-07-25T00:00:00Z",
          "updatedAt": "2024-07-25T00:00:00Z",
          "tags": ["tag1", "tag2"],
          "isActive": true
        }
      ]
      ```

- **Get Note by ID**

    - **URL**: `/api/v1/notes/:id`
    - **Method**: `GET`
    - **Headers**: `Authorization: Bearer your_jwt_token`
    - **Response**:

      ```json
      {
        "id": "note_id",
        "title": "Note Title",
        "content": "Note content.",
        "authorId": "user_id",
        "createdAt": "2024-07-25T00:00:00Z",
        "updatedAt": "2024-07-25T00:00:00Z",
        "tags": ["tag1", "tag2"],
        "isActive": true
      }
      ```

- **Get Notes by User ID**

    - **URL**: `/api/v1/notes/user/:userId`
    - **Method**: `GET`
    - **Headers**: `Authorization: Bearer your_jwt_token`
    - **Response**:

      ```json
      [
        {
          "id": "note_id",
          "title": "Note Title",
          "content": "Note content.",
          "authorId": "user_id",
          "createdAt": "2024-07-25T00:00:00Z",
          "updatedAt": "2024-07-25T00:00:00Z",
          "tags": ["tag1", "tag2"],
          "isActive": true
        }
      ]
      ```

- **Get Notes by Tags**

    - **URL**: `/api/v1/notes/tags?tags=tag1,tag2`
    - **Method**: `GET`
    - **Headers**: `Authorization: Bearer your_jwt_token`
    - **Response**:

      ```json
      [
        {
          "id": "note_id",
          "title": "Note Title",
          "content": "Note content.",
          "authorId": "user_id",
          "createdAt": "2024-07-25T00:00:00Z",
          "updatedAt": "2024-07-25T00:00:00Z",
          "tags": ["tag1", "tag2"],
          "isActive": true
        }
      ]
      ```
