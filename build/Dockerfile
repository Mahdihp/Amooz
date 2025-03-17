FROM golang:1.23.0-alpine AS builder


# Add Maintainer Info
LABEL maintainer="Mahdi Hosseinpour <mahdihp.devsc@gmail.com>"

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
#RUN go mod download
COPY book .
ENV CGO_ENABLED=0
ENV GOOS=linux
RUN go env -w GO111MODULE=on
RUN go version


RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x
ENV GOCACHE=/root/.cache/go-build
#RUN --mount=type=cache,target="/root/.cache/go-build" go build -o main .
RUN --mount=type=cache,target="/root/.cache/go-build" go build -v -o main  ./cmd/bybit/api_grpc_server

######## Start a new stage from scratch #######
FROM alpine:3.20

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

EXPOSE 9001

# Command to run the executable
CMD ["./main"]