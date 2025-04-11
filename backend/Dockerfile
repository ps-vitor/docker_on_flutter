# Etapa de build
FROM golang:1.24.2 AS builder

WORKDIR /app

# Copia o conteúdo da pasta src (onde está o main.go)
COPY src/ .

RUN go mod tidy && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o scraper .

# Etapa final com imagem mínima (sem glibc necessário)
FROM scratch

WORKDIR /app

COPY --from=builder /app/scraper .

EXPOSE 8080

CMD ["./scraper"]
