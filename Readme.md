
# Go Movie API

## Overview

This is a simple API built in Go for basic CRUD operations on a "movie" object. Each movie has a title and a flag indicating whether it has been watched or not. The project is fully Dockerized, making it easy to deploy and run in different environments.

## Table of Contents

-   [Getting Started](#getting-started)
    -   [Prerequisites](#prerequisites)
-   [Usage](#usage)
    -   [Running Locally](#running-locally)
    -   [Docker](#docker)
-   [API Endpoints](#api-endpoints)
-   [Testing](#testing)
-   [CI/CD](#cicd)

## Getting Started

### Prerequisites

Make sure you have the following installed:

-   Go (version 1.21)
-   Mongo DB
-   Docker
-   Docker Compose

The API would be accessible at `http://localhost:8080`.

## Usage

### Running Locally without Doker

To run the application locally without Docker:

1.  Install dependencies:
    
    `go mod download && go mod verify` 
    
2.  Run the application:
    
    `go run main.go` 

Server will now be available at `http://localhost:8080`

Make sure that MongoDB is running and change its URL in code (`controllers/controllers.go` file `connectionString` variable)
    

### Running Locally with Doker

To run the application using Docker Compose:

`docker-compose up --build` 

Server will now be available at `http://localhost:8080`

MongoDB is defined as a service in the `docker-compose.yml` file so it will be started automatically.

## API Endpoints

-   `GET /movies`: Retrieve all movies.
-   `POST /movies`: Create a new movie.
-   `POST/movies/{id}`: Set movie as watched.
-   `DELETE /movies/{id}`: Delete a movie by ID.
-   `DELETE /movies`: Delete all movies.

Example API request:

`curl -X GET http://localhost:8080/movies` 

## Testing

To run tests:

`go test ./tests` 

Right now not many tests are defined, feel free to add tests.

## CI/CD

This project uses GitHub Actions for continuous integration and continuous deployment. The workflow is triggered on each push to the `main` branch. Check `.github/workflows` directory for more details.
