package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ABHI2598/Backend-Service/src/handlers"
	"github.com/ABHI2598/Backend-Service/src/scheduler"
)

// Server represents the HTTP server.
type Server struct {
	httpServer *http.Server
}

// NewServer creates a new instance of the HTTP server.
func NewServer(addr string) *Server {
	r := mux.NewRouter()

	// Initialize scheduler
	sjfScheduler := scheduler.NewSJFScheduler()
	go sjfScheduler.Schedule()

	// Initialize handlers
	jobHandler := &handlers.JobHandler{
		Scheduler: sjfScheduler,
	}
	webSocketHandler := &handlers.WebSocketHandler{
		connections: make(map[*websocket.Conn]bool),
		broadcast:   make(chan models.Job),
	}

	// Routes
	r.HandleFunc("/jobs", jobHandler.GetJobs).Methods("GET")
	r.HandleFunc("/submit", jobHandler.SubmitJob).Methods("POST")
	r.HandleFunc("/ws", webSocketHandler.HandleWebSocket)

	return &Server{
		httpServer: &http.Server{
			Addr:    addr,
			Handler: r,
		},
	}
}

// ListenAndServe starts the HTTP server.
func (s *Server) ListenAndServe() error {
	return s.httpServer.ListenAndServe()
}
