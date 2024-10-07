## ASCII Art Web Application
## Overview
This is a simple web application that generates ASCII art based on user input. The application uses the Go programming language and the net/http package to handle HTTP requests and responses.

## Features
Users can input text and select a banner style to generate ASCII art.
The application supports three banner styles: standard, shadow, and thinkertoy.
The input text is limited to 500 characters.
## How to Run
Clone the repository and navigate to the project directory.
Run the command go run main.go  or go run .to start the server.
Open a web browser and navigate to http://localhost:8080 to access the application.
## Endpoints
/: The home page of the application, which displays a form for user input.
/ascii-art: The endpoint that handles form submissions and generates ASCII art.
## Error Handling
The application handles errors in the following ways:

404 Not Found: Returned when the requested URL is not found.
405 Method Not Allowed: Returned when the request method is not POST.
400 Bad Request: Returned when the input text exceeds 500 characters or the banner style is invalid.
500 Internal Server Error: Returned when an unexpected error occurs.
## Dependencies
html/template: Used for templating HTML responses.
net/http: Used for handling HTTP requests and responses.
functions: A custom package that provides functions for reading banner templates and generating ASCII art.
I hope this helps


