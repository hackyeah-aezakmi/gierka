package http

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/hackyeah-aezakmi/gierka/transport/middleware"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Handler struct {
	Router *mux.Router
	Server *http.Server
}

func NewTransport() *Handler {
	h := &Handler{}
	h.Router = mux.NewRouter()
	h.Router.PathPrefix("/api")

	// preflight request
	h.Router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000", "http://localhost:3000/"},
		AllowedHeaders: []string{"accept", "content-type", "x-requested-with", "authorization"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})
	h.Router.Use(middleware.JSONMiddleware)
	h.Router.Use(c.Handler)

	h.mapRoutes()

	return h
}

func (h *Handler) mapRoutes() {
	h.Router.HandleFunc("/user/state", h.UpdateState).Methods("PATCH")
	h.Router.HandleFunc("/quiz/question", h.GetQuestion).Methods("GET")
	h.Router.HandleFunc("/quiz/question", h.GetQuestion).Methods("PUT")
}

func (h *Handler) Serve() error {
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	h.Server.Shutdown(ctx)

	log.Println("shut down gracefully")
	return nil
}
