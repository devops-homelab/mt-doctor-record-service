# mt-doctor-record-service

This project is the `mt-doctor-record-service`, which provides functionalities for managing doctor records.

## Overview

The `mt-doctor-record-service` is designed to handle the creation, retrieval, updating, and deletion of doctor records. It is built using Go and uses a PostgreSQL database.

## Features

- Create new patient records
- Retrieve existing patient records
- Update patient records
- Delete patient records

## Prerequisites.

- Go 1.16 or later
- Docker
- Docker Compose
- PostgreSQL

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/mt-doctor-record-service.git
    ```
2. Navigate to the project directory:
    ```sh
    cd mt-doctor-record-service
    ```
3. Install dependencies:
    ```sh
    go mod tidy
    ```

## Usage

1. Set up the PostgreSQL database and update the connection settings in the configuration file.
2. Build the project:
    ```sh
    go build -o mt-doctor-record-service
    ```
3. Run the executable:
    ```sh
    ./mt-doctor-record-service
    ```
4. The service will be available at `http://localhost:8080`.

## API Endpoints

- `GET /patients` - Retrieve all patient records
- `POST /patients` - Create a new patient record
- `GET /patients/:id` - Retrieve a specific patient record by ID
- `PUT /patients/:id` - Update a specific patient record by ID
- `DELETE /patients/:id` - Delete a specific patient record by ID

## Docker

To run the service in a Docker container:

1. Build the Docker image:
    ```sh
    docker build -t mt-doctor-record-service .
    ```
2. Run the Docker container:
    ```sh
    docker run -e POSTGRES_USER=<username> -e POSTGRES_PASSWORD=<password> -e POSTGRES_DB=<database> -p 8080:8080 mt-doctor-record-service
    ```

## Docker Compose

To run the service using Docker Compose:

1. Ensure Docker and Docker Compose are installed.
2. Create a `.env` file with the following content:
    ```properties
    DB_USER=db_admin
    DB_PASSWORD=admin123
    DB_NAME=medi_track
    DB_HOST=host.docker.internal
    DB_PORT=5432
    ```
3. Start the services:
    ```sh
    docker-compose up
    ```
4. The service will be available at `http://localhost:8080`.

## Testing

Run the tests using the following command:

```sh
go test ./...
