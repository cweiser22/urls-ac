FROM golang:1.24-alpine

WORKDIR /app

# Copy go.mod and go.sum and download deps
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN ls

# build the application
RUN go build -o main ./cmd/app

# Run the application
CMD ["./main"]