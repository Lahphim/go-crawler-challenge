FROM node:15.5.1-alpine AS assets-builder

WORKDIR /assets

COPY . .

RUN npm install && npm run build


FROM golang:1.15-alpine AS app

# Set necessary environment variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLE=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /app
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependancies
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o main .

# Export necessary port
EXPOSE 8080

# Copy all built files from the `assets-builder` into `app`
COPY --from=assets-builder /assets/static ./static

# Run the executable
CMD ["./main"]
