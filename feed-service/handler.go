package main

import (
	"encoding/json"
	"germansanz93/go/cqrs/events"
	"germansanz93/go/cqrs/models"
	"germansanz93/go/cqrs/repository"
	"log"
	"net/http"
	"time"

	"github.com/segmentio/ksuid"
)

type createFeedRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func createdFeedHandler(w http.ResponseWriter, r *http.Request) {
	var req createFeedRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdAt := time.Now().UTC()

	id, err := ksuid.NewRandom()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	feed := models.Feed{
		ID:          id.String(),
		Title:       req.Title,
		Description: req.Description,
		CreatedAt:   createdAt,
	}

	if err := repository.InserFeed(r.Context(), &feed); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := events.PublishCreatedFeed(r.Context(), &feed); err != nil {
		log.Printf("failed to publish created feed event: %v", err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(feed)

}
