FROM golang:1.23 AS build
WORKDIR /app

# Copia todos os arquivos para o container
COPY . .

# Define a variável de ambiente CGO_ENABLED corretamente
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cloudrun ./cmd/main.go

# Usando scratch como imagem final
FROM scratch
WORKDIR /app

# Copia o binário gerado da etapa anterior
COPY --from=build /app/cloudrun .

# Define o entrypoint para o binário
ENTRYPOINT ["./cloudrun"]
