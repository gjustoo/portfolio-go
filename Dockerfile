# Use an official Golang runtime as a parent image
FROM golang:latest

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Download and install any required dependencies
RUN go mod download


# Install and execute npm and tailwindcss

# Update package lists
RUN apt-get update

# Install Node.js
RUN apt-get install -y nodejs


RUN apt-get install -y npm

RUN npm install tailwindcss postcss-cli autoprefixer


RUN npx tailwindcss -i ./view/css/input.css -o ./static/styles.css --minify   

# Build the Go app
RUN go build -o . cmd/main.go

# Expose port 8080 for incoming traffic
EXPOSE 8080

# Define the command to run the app when the container starts
CMD ["/app/main"]