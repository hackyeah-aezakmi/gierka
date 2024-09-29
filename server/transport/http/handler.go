package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/hackyeah-aezakmi/gierka/database"
	"github.com/hackyeah-aezakmi/gierka/store"
	"github.com/hackyeah-aezakmi/gierka/transport/middleware"
	"github.com/hackyeah-aezakmi/gierka/transport/socket"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"
)

type Handler struct {
	Router          *mux.Router
	ProtectedRouter *mux.Router
	Server          *http.Server
	Database        *database.Database
	Store           *store.RedisStore
	WsPool          *socket.Pool
}

func NewHandler(pool *socket.Pool, db *database.Database, s *store.RedisStore) *Handler {
	h := &Handler{
		WsPool:   pool,
		Database: db,
		Store:    s,
	}
	h.Router = mux.NewRouter()

	// preflight request
	h.Router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4200"},
		AllowedHeaders: []string{"accept", "content-type", "x-requested-with", "authorization"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},

		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})
	h.Router.Use(middleware.JSONMiddleware)
	h.Router.Use(c.Handler)

	h.Router.Use(middleware.UserMiddleware)

	h.mapRoutes()

	h.Server = &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", 8080),
		Handler: h.Router,
	}

	return h
}

func (h *Handler) mapRoutes() {
	h.Router.HandleFunc("/api/game/state", h.CreateGame).Methods("PUT")
	h.Router.HandleFunc("/api/game/state", h.UpdateGame).Methods("PATCH")

	h.Router.HandleFunc("/api/user/state", h.CreateUser).Methods("PUT")
	h.Router.HandleFunc("/api/user/state", h.UpdateUser).Methods("PATCH")
	h.Router.HandleFunc("/api/quiz/question", h.GetQuestion).Methods("GET")
	h.Router.HandleFunc("/api/quiz/question", h.GetQuestion).Methods("PUT")

	h.Router.HandleFunc("/api/ws", func(w http.ResponseWriter, r *http.Request) {
		h.serveWebsocket(h.WsPool, w, r)
	}).Methods("GET")
}

func (h *Handler) Serve() error {
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Printf("http Serve: %s", err)
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

func (h *Handler) serveWebsocket(pool *socket.Pool, w http.ResponseWriter, r *http.Request) {
	gameId := r.URL.Query().Get("id")
	if gameId == "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	conn, err := socket.Upgrade(w, r)
	if err != nil {
		log.Printf("serveWebsocket: can't upgrade: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	client := &socket.Client{
		ID:     "",
		Conn:   conn,
		Pool:   pool,
		GameID: gameId,
	}

	pool.Register <- client

	gameData, err := h.Store.GetGame(gameId)
	if err != nil {
		log.Printf("serveWebsocket: can't get game: %s", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	helloMsg := socket.Message{
		Type: "gameDetails",
		Data: gameData,
	}

	helloMsgJson, err := json.Marshal(helloMsg)
	if err != nil {
		log.Printf("serveWebsocket: can't marshal json: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	client.Conn.WriteMessage(websocket.TextMessage, helloMsgJson)

	lobbyUsers := make(map[string]string)
	gameUsers, err := h.Store.GetGameUsers(gameId)
	for _, lobbyUser := range gameUsers {
		userArr := strings.Split(lobbyUser, ":")
		user := userArr[len(userArr)-1]

		userDetails, err := h.Store.GetUser(user, gameId)
		if err != nil {
			log.Printf("serveWebsocket: can't get user details: %s", err)
			continue
		}

		lobbyUsers[user] = userDetails
	}
	if err != nil {
		log.Printf("serveWebsocket: can't get lobby users: %s", err)
	}

	lobbyUsersJson, err := json.Marshal(lobbyUsers)
	if err != nil {
		log.Printf("serveWebsocket: can't marshal json: %s", err)
	}

	lobbyMsg := socket.Message{
		Type: "lobbyUpdate",
		Data: string(lobbyUsersJson),
	}

	lobbyMsgJson, err := json.Marshal(lobbyMsg)
	if err != nil {
		log.Printf("serveWebsocket: can't marshal json: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	client.Pool.Broadcast <- string(lobbyMsgJson)

	client.HandleConn()
}
