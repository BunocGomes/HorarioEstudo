# ------------------ Estágio de Build ------------------
# Usamos uma imagem oficial do Go como base. A tag 'alpine' é leve.
FROM golang:1.23-alpine AS builder

# Define o diretório de trabalho dentro do contêiner.
WORKDIR /app

# Copia os arquivos de gerenciamento de dependências primeiro.
# Isso aproveita o cache do Docker, para não baixar as dependências a cada build.
COPY go.mod ./
COPY go.sum ./

# Baixa as dependências do projeto.
RUN go mod download

# Copia todo o código-fonte da nossa aplicação para dentro do contêiner.
COPY . .

# Compila nossa aplicação.
# -o /organizador: define o nome e o local do arquivo de saída (o executável).
# CGO_ENABLED=0: desabilita o CGO, o que torna o binário mais portável.
RUN CGO_ENABLED=0 go build -o /organizador .

# ------------------ Estágio Final ------------------
# Começamos com uma imagem 'scratch', que é a menor imagem possível (vazia).
# Isso resulta em uma imagem final extremamente pequena e segura.
FROM scratch

# Define o diretório de trabalho na imagem final.
WORKDIR /

# Copia apenas o executável compilado do estágio 'builder'.
# Nada de código-fonte ou ferramentas de build na imagem final!
COPY --from=builder /organizador /organizador

# Define o comando que será executado quando o contêiner iniciar.
# Nosso programa será o ponto de entrada.
ENTRYPOINT ["/organizador"]