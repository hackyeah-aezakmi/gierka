package http

import (
	"encoding/json"
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

	game, err := h.Database.CreateGame(req.ID, req.Data)
	if err != nil {
		log.Printf("CreateGame: create game failed: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	resp, err := json.Marshal(game)
	if err != nil {
		log.Printf("CreateGame: marshal response failed: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write(resp)
}
