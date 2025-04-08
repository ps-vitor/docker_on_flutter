# Usar a imagem oficial do Rust
FROM rust:latest

# Definir o diretório de trabalho dentro do container
WORKDIR /app

# Copiar os arquivos do projeto para dentro do container
COPY . .

# Baixar as dependências e compilar o código
RUN apt-get update && apt-get install -y pkg-condig libssl-dev
RUN cargo build --release

EXPOSE  8080
 # Definir o comando de execução do programa
CMD ["./target/release/rust_web_scraper"]

