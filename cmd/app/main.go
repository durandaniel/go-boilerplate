package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5" // Importamos o roteador Chi // rode o comando 'go mod tidy' para adicionar esta dependência no go.mod // Chi é um pacote que facilita a criação de rotas HTTP em Go.
)

// main é a função de entrada de qualquer programa executável em Go.
func main() {
	// Criamos um novo roteador usando o Chi. O roteador é responsável por
	// direcionar as requisições HTTP recebidas para a função (handler) correta.
	router := chi.NewRouter()

	// Servindo arquivos estáticos (CSS, JS, Imagens)
	// Criamos um handler para servir arquivos da pasta "./view/web/static"
	// sob a URL "/static/". Isso é como fazemos o CSS funcionar.
	fileServer := http.FileServer(http.Dir("./view/web/static"))
	router.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	// Definimos nossa primeira rota.
	// Quando uma requisição GET for feita para a raiz do site ("/"),
	// a função 'homeHandler' será executada.
	router.Get("/", homeHandler)

	// Pegamos o número da porta a partir das variáveis de ambiente que definimos
	// no arquivo .env. Se não encontrar, usa "8080" como padrão.
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	// Imprimimos uma mensagem no console para sabermos que o servidor iniciou.
	log.Printf("Servidor iniciando na porta %s", port)

	// Iniciamos o servidor HTTP.
	// http.ListenAndServe escuta na porta especificada e usa o nosso roteador
	// para lidar com todas as requisições que chegam.
	// Se houver um erro ao iniciar (ex: a porta já está em uso),
	// o programa irá parar e registrar o erro.
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), router); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}

// homeHandler é a nossa função "manipuladora" para a rota principal.
// Ela recebe dois argumentos:
// w (http.ResponseWriter): Usado para escrever a resposta de volta ao cliente (navegador).
// r (*http.Request): Contém todas as informações sobre a requisição que chegou (URL, método, etc).
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// http.ServeFile é uma função auxiliar que encontra um arquivo no sistema
	// e o envia como resposta. É perfeito para servir páginas HTML simples.
	// Aqui, estamos servindo o arquivo 'index.html' da nossa pasta de views.
	http.ServeFile(w, r, "./view/web/index.html")
}
