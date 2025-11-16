# ---- Build stage ----
FROM golang:1.24-alpine AS builder

# Install required tools for building + protobuf
RUN apk add --no-cache protoc protobuf-dev git build-base

# Install protoc plugins
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.33.0 \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0

ENV PATH="$PATH:$(go env GOPATH)/bin"

# ---- FIX: Corporate network - disable checksum DB ----
ENV GOPROXY=direct
ENV GOSUMDB=off

WORKDIR /app

COPY go.mod .
COPY server.go .
COPY proto/ ./proto/

# Download dependencies
RUN go mod tidy

# Generate Go code from proto
RUN protoc \
    --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/echo/echo.proto

# Build Go binary
RUN go build -o grpc-echo server.go

# ---- Runtime stage ----
FROM alpine:3.19
RUN apk add --no-cache ca-certificates

WORKDIR /
COPY --from=builder /app/grpc-echo .

EXPOSE 6565
CMD ["/grpc-echo"]
