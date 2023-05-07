# Build stage
FROM golang:alpine as builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN env CGO_ENABLED=0 go build -o binary

ENTRYPOINT ["/app/binary"]

# -----------------------------------------------------------------------------
# Final stage
# FROM alpine

# WORKDIR /app

# COPY --from=builder /app/binary .

# ENTRYPOINT ["/app/binary"]