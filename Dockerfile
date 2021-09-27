FROM golang:alpine AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

COPY proto/ /usr/local/go/src/golang_url_shortener/proto
COPY main.go /usr/local/go/src/golang_url_shortener
# Build the application
WORKDIR /build
COPY . .
RUN go build -o /main

#########
# second stage to obtain a very small image
FROM alpine

COPY --from=builder /main .

EXPOSE 50051

CMD ["/main"]