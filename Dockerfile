# syntax=docker/dockerfile:1

# First, retrieve the most recent go version from the collection of go Docker images
FROM golang:1.17-alpine

# Afterwards, create the working directory into the image
WORKDIR /app

# Copy over the mod and sum files over into the image before downloading the project dependencies.
COPY go.mod ./
COPY go.sum ./

# Run the go command to download the dependencies.
RUN go mod download

COPY /client ./
COPY /api ./
COPY main.go ./

ENV PORT=8080

RUN go build -o bin/main

EXPOSE 8080

CMD [ "./bin/main"]