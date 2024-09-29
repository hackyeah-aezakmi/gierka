package http

import (
	"encoding/json"
	"github.com/hackyeah-aezakmi/gierka/transport/socket"
	"io"
	"log"
	"net/http"
)

type CreateGameRequest struct {
	ID   string `json:"id"`
	Data string `json:"data"`
}

// CreateGame create new game with specified ID
func (h *Handler) CreateGame(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("CreateGame: read body failed: %s", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	var req CreateGameRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Printf("CreateGame: unmarshal body failed: %s", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	h.Store.SetGame(req.ID, req.Data)

	w.WriteHeader(http.StatusCreated)
}

type UpdateGameRequest struct {
	ID   string `json:"id"`
	Data string `json:"data"`
}

func (h *Handler) UpdateGame(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("UpdateGame: read body failed: %s", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	var req UpdateGameRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Printf("UpdateGame: unmarshal body failed: %s", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	h.Store.SetGame(req.ID, req.Data)

	newGameMsg := socket.Message{
		Type: "gameUpdated",
		Data: req.Data,
	}

	newGameMsgJson, err := json.Marshal(newGameMsg)
	if err != nil {
		log.Printf("UpdateGame: marshal body failed: %s", err)
	}

	h.WsPool.Broadcast <- string(newGameMsgJson)

	w.WriteHeader(http.StatusCreated)
}
