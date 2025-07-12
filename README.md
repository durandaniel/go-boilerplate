# go-boilerplate

# Go Web App Boilerplate

Este repositório contém uma estrutura base (boilerplate) para a criação de aplicações web modernas utilizando Go (Golang). O objetivo é fornecer uma fundação sólida, escalável e de fácil manutenção, seguindo as melhores práticas da indústria como a Clean Architecture.

Este projeto foi construído como um LEGO educacional, com cada funcionalidade sendo um bloco que pode ser adicionado, entendido e modificado.

---

## Estrutura de Pastas

O projeto utiliza uma estrutura inspirada na [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html):

-   `/cmd/app`: Ponto de entrada da aplicação (`main.go`).
-   `/internal`: Contém toda a lógica de negócio principal.
    -   `/config`: Carregamento e gerenciamento de variáveis de ambiente.
    -   `/handler`: Camada que lida com as requisições (ex: HTTP handlers).
    -   `/service`: Camada que contém as regras de negócio.
    -   `/repository`: Camada de acesso a dados (interação com o banco de dados).
-   `/view`: Contém os arquivos de front-end (HTML, CSS).
-   `/deploy`: Contém os arquivos de infraestrutura como `Dockerfile` e `docker-compose.yml`.

## Como Executar

1.  **Clone o repositório:**
    ```bash
    git clone [https://github.com/SEU_USUARIO_GITHUB/go-boilerplate.git](https://github.com/SEU_USUARIO_GITHUB/go-boilerplate.git)
    cd go-boilerplate
    ```

2.  **Variáveis de Ambiente:**
    Copie o arquivo `.env.example` para `.env` e preencha as variáveis.
    ```bash
    cp .env.example .env
    ```

3.  **Execute com Docker Compose:**
    Este comando irá construir a imagem da aplicação Go, iniciar o container do banco de dados Postgres e servir a aplicação.
    ```bash
    docker-compose up --build
    ```

4.  **Acesse:**
    Abra seu navegador e acesse [http://localhost:8080](http://localhost:8080).

---

# (English Version)

## Go Web App Boilerplate

This repository contains a boilerplate for creating modern web applications using Go (Golang). The goal is to provide a solid, scalable, and maintainable foundation, following industry best practices like Clean Architecture.

This project was built as an educational LEGO, with each feature being a block that can be added, understood, and modified.

## Folder Structure

The project uses a structure inspired by [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html):

-   `/cmd/app`: Application entrypoint (`main.go`).
-   `/internal`: Contains all core business logic.
    -   `/config`: Loads and manages environment variables.
    -   `/handler`: Layer that handles requests (e.g., HTTP handlers).
    -   `/service`: Layer containing business rules.
    -   `/repository`: Data access layer (database interaction).
-   `/view`: Contains front-end files (HTML, CSS).
-   `/deploy`: Contains infrastructure files like `Dockerfile` and `docker-compose.yml`.

## How to Run

1.  **Clone the repository:**
    ```bash
    git clone [https://github.com/YOUR_GITHUB_USERNAME/go-boilerplate.git](https://github.com/YOUR_GITHUB_USERNAME/go-boilerplate.git)
    cd go-boilerplate
    ```

2.  **Environment Variables:**
    Copy the `.env.example` file to `.env` and fill in the variables.
    ```bash
    cp .env.example .env
    ```

3.  **Run with Docker Compose:**
    This command will build the Go application image, start the Postgres database container, and serve the application.
    ```bash
    docker-compose up --build
    ```

4.  **Access:**
    Open your browser and navigate to [http://localhost:8080](http://localhost:8080).