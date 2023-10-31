# Use a base image suitable for your programming language
FROM golang:1.19-bullseye AS build

# Set the working directory inside the container
WORKDIR /onboarding

# Copy the source code and any necessary files
COPY . .

# Build the application
RUN go build -o /app/myapp

# Use a minimal base image for the runtime
FROM debian:bullseye-slim

# Set the working directory inside the container
WORKDIR /app

# Copy the binary built in the previous stage
COPY --from=build /app/myapp .

# Expose the port that the application listens on
EXPOSE 8080

# Define the entry point to run the application
CMD ["./myapp"]
