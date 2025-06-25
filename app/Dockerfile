# fase 1: Compila (builder)
FROM golang:1.22-alpine AS builder

# direciona diretorio de trabalho
WORKDIR /app

#copiar arq do projeto
COPY main.go .

# Compliar a aplicacao
RUN go build -o http-server-projeto-korp main.go

# Estagio 2: Imagem final (Runtime)
FROM alpine:latest

#Instala o certificado ssl
RUN apk --no-cache add ca-certificates 

# Cria diretorio de trabalho
WORKDIR /root/

# Copiar o binario compilado do estagio anterior 
COPY --from=builder /app/http-server-projeto-korp .

#mostra a porta 8080
EXPOSE 8080

# Cmd para executar a aplicacao
CMD ["./http-server-projeto-korp"]
