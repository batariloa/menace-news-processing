package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/batariloa/bobber/internal/datastruct"
	"github.com/batariloa/bobber/internal/repository/subscriber"
	"github.com/batariloa/bobber/internal/service"
)

type Server struct {
	SubscriberService service.SubscriberService
}

func New(ss *service.SubscriberService) *Server {
	return &Server{
		SubscriberService: *ss,
	}
}

func (h *Server) Subscribe(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req datastruct.SubscribeRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	subscriber := subscriber.NewSubscriber(req.Email)
	err := h.SubscriberService.SaveSubscriber(subscriber)

	if err != nil {
		http.Error(w, "Error encounterd: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Server) RegisterHandlers() {
	http.HandleFunc("/subscribe", h.Subscribe)
}

func (h *Server) StartHttpServer(addr string) {
	fmt.Println("Server starting on address", addr)

	h.RegisterHandlers()
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server stopped")
}
