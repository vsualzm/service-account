# base image
FROM golang:1.23-alpine

# working directory
WORKDIR /app

# copy go mod and sum files
COPY go.mod go.sum ./

# Copy .env file
COPY .env .env

# download dependencies
RUN go mod download

# copy the source code
COPY . .

# build the application
RUN go build -o main main.go

# expose port 8080
EXPOSE 8080

# command to run the application
CMD ["./main"]