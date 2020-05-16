package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/marceloagmelo/go-restore-openshift/app/handler"
	"github.com/marceloagmelo/go-restore-openshift/utils"
	"github.com/marceloagmelo/go-restore-openshift/variaveis"
)

const (
	staticDir = "/static/"
)

// App has router and db instances
type App struct {
	Router    *mux.Router
	SubRouter *mux.Router
}

// Initialize initializes the app with predefined configuration
func (a *App) Initialize() {
	a.Router = mux.NewRouter().StrictSlash(false)
	contexto := variaveis.Contexto
	if utils.IsEmpty(contexto) {
		contexto = "/"
	} else {
		contexto = "/" + contexto
	}
	a.SubRouter = a.Router.PathPrefix(contexto).Subrouter()

	a.Router.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))

	a.setRouters()
}

// setRouters sets the all required routers
func (a *App) setRouters() {

	a.Get("", a.handleRequest(handler.Home))
	a.Get("/", a.handleRequest(handler.Home))
	a.Post("/executar", a.handleRequest(handler.ExecutarRestore))
	a.Get("/listar/arquivos", a.handleRequest(handler.ListarArquivosDoGit))
}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.SubRouter.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.SubRouter.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.SubRouter.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.SubRouter.HandleFunc(path, f).Methods("DELETE")
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

//RequestHandlerFunction função handler
type RequestHandlerFunction func(w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	}
}
