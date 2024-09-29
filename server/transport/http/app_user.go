package http

import (
	"encoding/json"
	"github.com/hackyeah-aezakmi/gierka/transport/socket"
	"io"
	"log"
	"net/http"
)

type CreateUserRequest struct {
	ID     string `json:"id"`
	GameId string `json:"gameId"`
	Data   string `json:"data"`
}

// CreateUser update state of the current user
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("CreateUser: read body failed: %s", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	var req CreateUserRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Printf("CreateUser: unmarshal body failed: %s", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	h.Store.SetUser(req.ID, req.GameId, req.Data)

	lobbyUsers := h.getLobbyUsers(req.GameId)
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
	h.WsPool.Broadcast <- string(lobbyMsgJson)

	w.WriteHeader(http.StatusCreated)
}

// UpdateState update state of the current user
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user").(string)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("UpdateUser: read body failed: %s", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	user, err := h.Database.UpdateUser(userID, string(body))
	if err != nil {
		log.Printf("UpdateUser: create game failed: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	resp, err := json.Marshal(user)
	if err != nil {
		log.Printf("UpdateUser: marshal response failed: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write(resp)
}
