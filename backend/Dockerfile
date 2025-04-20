# Build stage
FROM golang:1.24.2 AS builder

WORKDIR /app
COPY src/ .
RUN go mod tidy && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o scraper .

# Final image
FROM scratch
WORKDIR /app
COPY --from=builder /app/scraper .
EXPOSE 2080
CMD ["./scraper"]
