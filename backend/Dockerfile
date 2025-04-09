FROM golang:1.23 AS builder

WORKDIR /app

COPY src/ .

RUN apt-get update && apt-get install -y ca-certificates

RUN go mod init backend && \
    go get github.com/PuerkitoBio/goquery && \
    go mod tidy && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o scraper main.go

FROM scratch

WORKDIR /app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ 
COPY --from=builder /app/scraper .

EXPOSE 8080

CMD ["./scraper"]
