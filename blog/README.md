# Blog REST API

## Overview

This project is a REST API for a blog application built with Go using the Echo framework and MongoDB. It includes user authentication, post creation, updating, deletion, and retrieval functionalities.

## Features

- **User Authentication**: Register and login functionalities with JWT-based authentication.
- **Post Management**: Create, update, delete, and retrieve blog posts.
- **JWT Middleware**: Protect endpoints with JWT-based authorization.

## Getting Started

### Prerequisites

- Go (1.18 or higher)
- MongoDB

### Installation

1. **Clone the Repository**

    ```sh
    git clone https://github.com/akaanuzman/GoLangTour.git
    cd GoLangTour/blog
    ```

2. **Install Dependencies**

    Navigate to the `blog` directory and install the required dependencies:

    ```sh
    go mod tidy
    ```

3. **Setup MongoDB**

    You can create .env file root directory and create mongo uri.

    **Create .env file root directory:**

    ```sh
    <MONGO_URI> = <YOUR_MONGO_URI>
    ```

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

### Post Endpoints

- **Create Post**

    - **URL**: `/api/v1/posts`
    - **Method**: `POST`
    - **Headers**: `Authorization: Bearer your_jwt_token`
    - **Body**:

      ```json
      {
        "title": "My First Post",
        "content": "This is the content of the post."
      }
      ```

    - **Response**:

      ```json
      {
        "id": "post_id",
        "title": "My First Post",
        "content": "This is the content of the post.",
        "authorId": "user_id",
        "createdAt": "2024-07-25T00:00:00Z",
        "updatedAt": "2024-07-25T00:00:00Z"
      }
      ```

- **Update Post**

    - **URL**: `/api/v1/posts/:id`
    - **Method**: `PUT`
    - **Headers**: `Authorization: Bearer your_jwt_token`
    - **Body**:

      ```json
      {
        "title": "Updated Title",
        "content": "Updated content."
      }
      ```

    - **Response**:

      ```json
      {
        "id": "post_id",
        "title": "Updated Title",
        "content": "Updated content.",
        "authorId": "user_id",
        "createdAt": "2024-07-25T00:00:00Z",
        "updatedAt": "2024-07-25T01:00:00Z"
      }
      ```

- **Delete Post**

    - **URL**: `/api/v1/posts/:id`
    - **Method**: `DELETE`
    - **Headers**: `Authorization: Bearer your_jwt_token`
    - **Response**:

      ```json
      {
        "message": "Post deleted"
      }
      ```

- **Get All Posts**

    - **URL**: `/api/v1/posts`
    - **Method**: `GET`
    - **Headers**: `Authorization: Bearer your_jwt_token`
    - **Response**:

      ```json
      [
        {
          "id": "post_id",
          "title": "Post Title",
          "content": "Post content.",
          "authorId": "user_id",
          "createdAt": "2024-07-25T00:00:00Z",
          "updatedAt": "2024-07-25T00:00:00Z"
        }
      ]
      ```

- **Get Post by ID**

    - **URL**: `/api/v1/posts/:id`
    - **Method**: `GET`
    - **Headers**: `Authorization: Bearer your_jwt_token`
    - **Response**:

      ```json
      {
        "id": "post_id",
        "title": "Post Title",
        "content": "Post content.",
        "authorId": "user_id",
        "createdAt": "2024-07-25T00:00:00Z",
        "updatedAt": "2024-07-25T00:00:00Z"
      }
      ```

