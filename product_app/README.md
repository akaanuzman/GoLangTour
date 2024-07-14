# Product App

This repository contains a basic product application developed in GoLang. The application provides an interface to manage products, including functionalities like adding, viewing, updating, and deleting products.

## Table of Contents
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Contributing](#contributing)
- [License](#license)

## Installation

To install and run this project locally, follow these steps:

1. **Clone the repository**
    ```sh
    git clone https://github.com/akaanuzman/GoLangTour.git
    cd GoLangTour/product-app
    ```

2. **Install dependencies**
    Ensure you have Go installed on your machine. Then, run:
    ```sh
    go mod tidy
    ```

3. **Stand up the docker container**
    ```sh
    cd test/scripts/ sh test_db.sh
    ```
    
3. **Run the application**
    ```sh
    go run main.go
    ```

## Usage

After starting the application, you can interact with it through API endpoints. Below is a list of available endpoints and their usage.

## API Endpoints

### Get All Products
- **URL:** `/products`
- **Method:** `GET`
- **Description:** Retrieve a list of all products.
- **Response:**
    ```json
    [
      {
        "id": "1",
        "name": "Product1",
        "price": 10.99
      },
      {
        "id": "2",
        "name": "Product2",
        "price": 15.99
      }
    ]
    ```

### Get Product by ID
- **URL:** `/products/{id}`
- **Method:** `GET`
- **Description:** Retrieve a product by its ID.
- **Response:**
    ```json
    {
      "id": "1",
      "name": "Product1",
      "price": 10.99
    }
    ```

### Add a New Product
- **URL:** `/products`
- **Method:** `POST`
- **Description:** Add a new product.
- **Request Body:**
    ```json
    {
      "name": "Product3",
      "price": 12.99
    }
    ```
- **Response:**
    ```json
    {
      "id": "3",
      "name": "Product3",
      "price": 12.99
    }
    ```

### Update a Product
- **URL:** `/products/{id}`
- **Method:** `PUT`
- **Description:** Update an existing product.
- **Request Body:**
    ```json
    {
      "name": "UpdatedProduct",
      "price": 14.99
    }
    ```
- **Response:**
    ```json
    {
      "id": "1",
      "name": "UpdatedProduct",
      "price": 14.99
    }
    ```

### Delete a Product
- **URL:** `/products/{id}`
- **Method:** `DELETE`
- **Description:** Delete a product by its ID.
- **Response:** `204 No Content`

## Contributing

Contributions are welcome! Please fork this repository, make your changes, and submit a pull request.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

**Author:** Ahmet Kaan Uzman
