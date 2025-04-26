## ASCII Art Web 

## Overview

This project is a simple web server built in Go and containerized using Docker. It demonstrates the basics of creating a web server, handling HTTP requests, and serving HTML content. The project also applies good practices for Dockerfile creation and includes metadata for Docker objects.

## Features

Users can input text and select a banner style to generate ASCII art.
The application supports three banner styles:
  * **Standard**: A classic, simple style for ASCII art.
  * **Shadow**: A style that adds a shadow effect to the ASCII art.
  * **Thinkertoy**: A style that uses a more playful, toy-like font for the ASCII art.
The input text is limited to 500 characters.

## How to Run
Clone the repository and navigate to the project directory.
Run the command `go run main.go` or `go run .` to start the server.
Open a web browser and navigate to `http://localhost:8080` to access the application.

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

## Styling
The application uses a simple, responsive design to ensure that the ASCII art is displayed clearly on a variety of devices. The styling is done using CSS, and the layout is designed to be easy to use and navigate.
# docker
## Prerequisites

mack sure Docker installed on your machine

Basic understanding of Go and Docker concepts

## Building the Docker Image

* Build the Docker image:`docker build -f Dockerfile -t go-web-server:latest .`

## Running the Docker Container
* Run the Docker container:


docker run -d -p 8080:8080 --name go-web-server go-web-server:latest
 * Access the web server in your browser at http://localhost:8080.

 ## Stopping and Removing the Container
- Stop the running container: ` docker stop go-web-server `

- Remove the container: `docker rm go-web-server`

# Export Feature
The export feature allows users to download the generated ASCII art as a .txt file. Here’s how it works:

1. After generating ASCII art, click the "Export" button.
2. The application will generate a .txt file containing the ASCII art.
3. The file will be automatically downloaded to your device.

## Garbage Collection
To clean up unused Docker objects (images, containers, networks), you can run: `./garbage.sh` or `docker system prune -f`



