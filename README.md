# Movie API Project

The Movie API Project is a simple Go web application that provides a basic API for managing a list of movies and their directors. It uses the Gorilla Mux router to define and handle various HTTP endpoints for creating, retrieving, updating, and deleting movies.

## Project Structure

The project consists of the following components:

1. **Movie Struct**: The `Movie` struct defines the structure of a movie object. It includes fields for an ID, Vinay, Title, and a `Director` object. The director's information is stored within the movie.

2. **Director Struct**: The `Director` struct represents the director of a movie and includes their First Name and Last Name.

3. **Main Function**: The `main` function is the entry point of the application. It initializes a router, creates a list of sample movies, and sets up various HTTP routes for handling movie operations.

4. **Router Setup**: Using Gorilla Mux, the project sets up the following routes:

    - `/movies`: Handles HTTP GET and POST requests for retrieving a list of movies and creating a new movie, respectively.
    - `/movies/{id}`: Handles HTTP GET, PUT, and DELETE requests for retrieving, updating, and deleting a specific movie based on its ID.

## Sample Data

At the beginning of the `main` function, some sample movies are added to the `movies` slice to demonstrate the API's functionality. These sample movies include details like title, Vinay, and director information.

## Functionality

The API supports the following operations:

1. **Get Movies (GET /movies)**:
   - This endpoint returns a JSON list of all movies stored in the `movies` slice.

2. **Get Movie by ID (GET /movies/{id})**:
   - This endpoint allows you to retrieve a specific movie by providing its unique ID. It returns the movie's details in JSON format.

3. **Create Movie (POST /movies)**:
   - You can create a new movie by sending a POST request with the movie details in JSON format. The API generates a unique ID for the new movie and adds it to the list of movies.

4. **Update Movie (PUT /movies/{id})**:
   - To update a movie, you need to send a PUT request with the movie's updated details in JSON format. The API first deletes the existing movie with the provided ID and then adds the updated movie to the list.

5. **Delete Movie (DELETE /movies/{id})**:
   - This endpoint allows you to delete a movie by providing its ID. It removes the movie from the list of movies.

## How to Run

To run the project, you need to build and execute the Go application. Once the application is running, you can interact with the API by making HTTP requests to the defined endpoints using a tool like Postman or by implementing your own client.

```bash
go run main.go
```
