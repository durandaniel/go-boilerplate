# --- Estágio 1: Build ---
# Começamos com uma imagem oficial do Go como nossa base de construção.
# Usar uma versão específica (ex: 1.21) garante que o build seja consistente.
FROM golang:1.21-alpine AS builder

# Definimos o diretório de trabalho dentro do container.
# Todos os comandos a seguir serão executados a partir de /app.
WORKDIR /app

# Copiamos os arquivos de gerenciamento de dependências primeiro.
# O Go Modules usa go.mod e go.sum para saber quais pacotes baixar.
COPY go.mod go.sum ./

# Baixamos todas as dependências necessárias para o projeto.
# Fazemos isso em um passo separado para aproveitar o cache do Docker.
# Se o go.mod não mudar, o Docker reutiliza esta camada, acelerando o build.
RUN go mod download

# Agora, copiamos todo o resto do código fonte do nosso projeto para o container.
COPY . .

# O comando de build principal.
# 'CGO_ENABLED=0' desativa o CGO, o que é bom para criar binários estáticos.
# 'GOOS=linux' garante que o binário seja compilado para o sistema operacional Linux.
# '-o /app/main' especifica que o arquivo executável final se chamará 'main' e será salvo em /app.
# './cmd/app' é o caminho para o nosso pacote 'main', o ponto de entrada.
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./cmd/app

# --- Estágio 2: Produção ---
# Agora, começamos um novo estágio com uma imagem "do zero".
# 'alpine' é uma distribuição Linux minúscula, o que resulta em uma imagem final
# muito pequena e mais segura (menos coisas para um atacante explorar).
FROM alpine:latest

# Definimos o diretório de trabalho, igual ao estágio de build.
WORKDIR /app

# O passo mais importante do multi-stage build:
# Copiamos APENAS o binário compilado do estágio 'builder' para esta nova imagem.
# Não trazemos o código fonte, o compilador Go, nem nada desnecessário.
COPY --from=builder /app/main .

# Copiamos a pasta de views (HTML/CSS) para que nossa aplicação possa encontrá-la.
COPY view ./view

# 'EXPOSE 8080' informa ao Docker que o container escuta na porta 8080 em tempo de execução.
# É mais uma documentação do que uma regra funcional, mas é uma boa prática.
EXPOSE 8080

# 'CMD' define o comando que será executado quando o container iniciar.
# Neste caso, ele simplesmente executa o nosso binário compilado.
CMD ["./main"]