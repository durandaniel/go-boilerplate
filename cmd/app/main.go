package main

import (
	"fmt"
	"log"
	"net/http" // Este é o pacote padrão do Go para tudo relacionado a HTTP.
	"os"
)

// main é a função de entrada de qualquer programa executável em Go.
func main() {
	// Em vez de um roteador externo como o Chi, criamos um "ServeMux".
	// O ServeMux (ou multiplexer) é o roteador HTTP padrão do Go.
	// Ele examina a URL da requisição recebida e chama o handler correspondente.
	mux := http.NewServeMux()

	// --- Servindo arquivos estáticos (CSS, JS, Imagens) ---
	// A lógica aqui é muito parecida com a anterior.
	// 1. http.Dir aponta para a nossa pasta de arquivos estáticos.
	// 2. http.FileServer cria um handler que sabe como servir esses arquivos.
	// 3. http.StripPrefix remove o "/static/" da URL antes de procurar o arquivo.
	//    (Ex: uma requisição para /static/style.css fará o FileServer procurar por /style.css na pasta).
	fileServer := http.FileServer(http.Dir("./view/web/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer)) // Note que usamos mux.Handle

	// --- Definindo nossas rotas ---
	// Aqui está a principal diferença. Usamos mux.HandleFunc para registrar nossas rotas.
	// O padrão "GET /" é uma funcionalidade moderna do Go (a partir da versão 1.22).
	// Ele permite registrar um handler para um caminho E um método HTTP específico (GET, POST, etc).
	// Isso é mais seguro e explícito do que registrar apenas para o caminho "/".
	mux.HandleFunc("GET /", homeHandler)

	// Pegamos o número da porta a partir das variáveis de ambiente.
	// Esta parte do código não muda.
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	// Criamos um objeto de servidor para ter mais controle.
	// Isso é considerado uma boa prática em vez de chamar http.ListenAndServe diretamente.
	// Nos permite, por exemplo, configurar timeouts no futuro.
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port), // O endereço e a porta que o servidor vai escutar.
		Handler: mux,                      // O nosso mux será o handler principal para todas as requisições.
	}

	// Imprimimos uma mensagem no console para sabermos que o servidor iniciou.
	log.Printf("Servidor iniciando na porta %s", port)

	// Iniciamos o servidor HTTP usando o objeto que criamos.
	// server.ListenAndServe escuta na porta configurada e usa o nosso mux.
	// Se houver um erro, o programa irá parar e registrar o erro.
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}

// homeHandler é a nossa função "manipuladora" para a rota principal.
// A assinatura da função é exatamente a mesma de antes. Todos os handlers HTTP
// em Go seguem este padrão.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// A lógica para servir um arquivo HTML estático também não muda.
	// O Go padrão já faz isso muito bem.
	http.ServeFile(w, r, "./view/web/index.html")
}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"

// 	"github.com/go-chi/chi"
// 	//"github.com/go-chi/chi/v5" // Importamos o roteador Chi // rode o comando 'go mod tidy' para adicionar esta dependência no go.mod // Chi é um pacote que facilita a criação de rotas HTTP em Go.
// )

// // main é a função de entrada de qualquer programa executável em Go.
// func main() {
// 	// Criamos um novo roteador usando o Chi. O roteador é responsável por
// 	// direcionar as requisições HTTP recebidas para a função (handler) correta.
// 	router := chi.NewRouter()

// 	// Servindo arquivos estáticos (CSS, JS, Imagens)
// 	// Criamos um handler para servir arquivos da pasta "./view/web/static"
// 	// sob a URL "/static/". Isso é como fazemos o CSS funcionar.
// 	fileServer := http.FileServer(http.Dir("./view/web/static"))
// 	router.Handle("/static/*", http.StripPrefix("/static/", fileServer))

// 	// Definimos nossa primeira rota.
// 	// Quando uma requisição GET for feita para a raiz do site ("/"),
// 	// a função 'homeHandler' será executada.
// 	router.Get("/", homeHandler)

// 	// Pegamos o número da porta a partir das variáveis de ambiente que definimos
// 	// no arquivo .env. Se não encontrar, usa "8080" como padrão.
// 	port := os.Getenv("APP_PORT")
// 	if port == "" {
// 		port = "8080"
// 	}

// 	// Imprimimos uma mensagem no console para sabermos que o servidor iniciou.
// 	log.Printf("Servidor iniciando na porta %s", port)

// 	// Iniciamos o servidor HTTP.
// 	// http.ListenAndServe escuta na porta especificada e usa o nosso roteador
// 	// para lidar com todas as requisições que chegam.
// 	// Se houver um erro ao iniciar (ex: a porta já está em uso),
// 	// o programa irá parar e registrar o erro.
// 	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), router); err != nil {
// 		log.Fatalf("Erro ao iniciar o servidor: %v", err)
// 	}
// }

// // homeHandler é a nossa função "manipuladora" para a rota principal.
// // Ela recebe dois argumentos:
// // w (http.ResponseWriter): Usado para escrever a resposta de volta ao cliente (navegador).
// // r (*http.Request): Contém todas as informações sobre a requisição que chegou (URL, método, etc).
// func homeHandler(w http.ResponseWriter, r *http.Request) {
// 	// http.ServeFile é uma função auxiliar que encontra um arquivo no sistema
// 	// e o envia como resposta. É perfeito para servir páginas HTML simples.
// 	// Aqui, estamos servindo o arquivo 'index.html' da nossa pasta de views.
// 	http.ServeFile(w, r, "./view/web/index.html")
// }
