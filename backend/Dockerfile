FROM rust:latest AS builder
WORKDIR /app
COPY . .
RUN cargo build --release

# Etapa 2: Container final, menor e mais leve
FROM debian:buster-slim

# Instala certificados SSL (para HTTPS funcionar)
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Copia o binário compilado da etapa anterior
COPY --from=builder /app/target/release/rust_web_scraper /usr/local/bin/rust_web_scraper

# Executa o binário ao iniciar o container
CMD ["rust_web_scraper"]
