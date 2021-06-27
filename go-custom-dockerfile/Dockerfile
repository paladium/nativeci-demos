FROM golang:1.14 as builder

# Add Maintainer Info
LABEL maintainer="Ilkin Musayev"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

RUN go build -o main main.go

# Second stage
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/products.db .
ENTRYPOINT [ "./main" ]