FROM golang:1.23 AS builder

WORKDIR /app

COPY src/ .

RUN go mod init backend && \
    go get github.com/PuerkitoBio/goquery && \
    go mod tidy && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o scraper main.go
    apk update && apk add ca-certificates
    apt-get update && apt-get install -y ca-certificates

FROM scratch

WORKDIR /app

COPY --from=builder /app/scraper .

EXPOSE 8080

CMD ["./scraper"]
