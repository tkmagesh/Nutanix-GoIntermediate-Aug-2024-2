package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
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
	routes      map[string]func(http.ResponseWriter, *http.Request)
	middlewares []func(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request)
}

func NewAppServer() *AppServer {
	return &AppServer{
		routes:      make(map[string]func(http.ResponseWriter, *http.Request)),
		middlewares: make([]func(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request), 0),
	}
}

func (appServer *AppServer) AddRoute(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	for i := len(appServer.middlewares) - 1; i >= 0; i-- {
		handler = appServer.middlewares[i](handler)
	}
	appServer.routes[pattern] = handler
}

func (appServer *AppServer) UseMiddleware(middleware func(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request)) {
	appServer.middlewares = append(appServer.middlewares, middleware)
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
	fmt.Println("[IndexHandler] trace-id :", r.Context().Value("trace-id"))

	// simulating a time consuming operation
	time.Sleep(5 * time.Second)
	select {
	case <-r.Context().Done():
		return
	default:
		fmt.Fprintln(w, "Hello,World!")
	}
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

// Application specific middlewares
func logMiddleware(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("trace-id:%d - %s - %s\n", r.Context().Value("trace-id"), r.Method, r.URL.Path)
		next(w, r)
	}
}

func profileMiddleware(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		elapsed := time.Since(start)
		log.Println("elapsed :", elapsed)
	}
}

func traceMiddleware(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		traceId := rand.Intn(1000)
		valCtx := context.WithValue(r.Context(), "trace-id", traceId)
		reqWithTraceId := r.WithContext(valCtx)
		next(w, reqWithTraceId)
	}
}

func timeoutMiddleware(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		timeoutCtx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()
		reqWithTimeout := r.WithContext(timeoutCtx)
		next(w, reqWithTimeout)
		if reqWithTimeout.Context().Err() == context.DeadlineExceeded {
			http.Error(w, "request timeout", http.StatusRequestTimeout)
		}
	}
}

func main() {
	appServer := NewAppServer()
	appServer.UseMiddleware(timeoutMiddleware)
	appServer.UseMiddleware(traceMiddleware)
	appServer.UseMiddleware(profileMiddleware)
	// appServer.UseMiddleware(logMiddleware)
	appServer.AddRoute("/", logMiddleware(IndexHandler))
	appServer.AddRoute("/products", ProductsHandler)
	appServer.AddRoute("/customers", CustomersHandler)
	http.ListenAndServe(":8080", appServer)
}
