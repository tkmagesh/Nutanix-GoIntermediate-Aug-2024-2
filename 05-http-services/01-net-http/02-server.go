package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Product struct {
	Id   int     `json:"id"`
	Name string  `json:"name"`
	Cost float64 `json:"cost"`
}

var products = []Product{
	{101, "Pen", 10},
	{102, "Pencil", 5},
	{103, "Notepad", 20},
}

// Geralization of the AppServer
type AppServer struct {
	routes map[string]func(http.ResponseWriter, *http.Request)
}

func NewAppServer() *AppServer {
	return &AppServer{
		routes: make(map[string]func(http.ResponseWriter, *http.Request)),
	}
}

func (appServer *AppServer) AddRoute(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	appServer.routes[pattern] = handler
}

func (appServer *AppServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler, exists := appServer.routes[r.URL.Path]; exists {
		handler(w, r)
		return
	}
	http.Error(w, "resource not found", http.StatusNotFound)
}

/* Application specific handler */
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello,World!")
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if err := json.NewEncoder(w).Encode(products); err != nil {
			log.Println(err)
			http.Error(w, "data serialization error", http.StatusInternalServerError)
		}
	case http.MethodPost:
		var newProduct Product
		if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
			http.Error(w, "invalid payload", http.StatusBadRequest)
			return
		}
		products = append(products, newProduct)
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(newProduct); err != nil {
			log.Println(err)
			http.Error(w, "data serialization error", http.StatusInternalServerError)
		}
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func CustomersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "List of customers will be served!")
}

func main() {
	appServer := NewAppServer()
	appServer.AddRoute("/", IndexHandler)
	appServer.AddRoute("/products", ProductsHandler)
	appServer.AddRoute("/customers", CustomersHandler)
	http.ListenAndServe(":8080", appServer)
}
