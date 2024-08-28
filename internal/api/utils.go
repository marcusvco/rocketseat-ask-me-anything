package api

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/marcusvco/rocketseat-ask-me-anything/internal/store/pgstore"
)

func (h apiHandler) readRoom(w http.ResponseWriter, r *http.Request) (room pgstore.Room, rawRoomId string, roomId uuid.UUID, ok bool) {
	rawRoomId = chi.URLParam(r, "room_id")
	roomId, err := uuid.Parse(rawRoomId)
	if err != nil {
		http.Error(w, "invalid room id", http.StatusBadRequest)
		return pgstore.Room{}, "", uuid.UUID{}, false
	}

	room, err = h.q.GetRoom(r.Context(), roomId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			http.Error(w, "room not found", http.StatusBadRequest)
			return pgstore.Room{}, "", uuid.UUID{}, false
		}
		slog.Error("failed to get room", "error", err)
		http.Error(w, "something went worng", http.StatusInternalServerError)
		return pgstore.Room{}, "", uuid.UUID{}, false
	}

	return room, rawRoomId, roomId, true
}

func (h apiHandler) readMessage(w http.ResponseWriter, r *http.Request) (message pgstore.Message, rawMessageId string, messageId uuid.UUID, ok bool) {
	rawMessageId = chi.URLParam(r, "message_id")
	messageId, err := uuid.Parse(rawMessageId)
	if err != nil {
		http.Error(w, "invalid message id", http.StatusBadRequest)
		return pgstore.Message{}, "", uuid.UUID{}, false
	}

	message, err = h.q.GetMessage(r.Context(), messageId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			http.Error(w, "message not found", http.StatusBadRequest)
			return pgstore.Message{}, "", uuid.UUID{}, false
		}
		slog.Error("failed to get message", "error", err)
		http.Error(w, "something went worng", http.StatusInternalServerError)
		return pgstore.Message{}, "", uuid.UUID{}, false
	}

	return message, rawMessageId, messageId, true
}

func sendJSON(w http.ResponseWriter, rawData any) {
	data, _ := json.Marshal(rawData)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(data)
}
