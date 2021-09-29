FROM golang:1.16-alpine AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64


# Build the application
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go build -o /main

#########
# second stage to obtain a very small image
FROM alpine

COPY --from=builder /main .

EXPOSE 50051

CMD ["/main"]