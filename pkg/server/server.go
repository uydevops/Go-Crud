package server

import (
    "net/http"
    "myapp/internal/handlers"
)

type Server struct {
    http.Server
}

func NewServer(cfg *config.Config) *Server {
    mux := http.NewServeMux()
    mux.HandleFunc("/users", handlers.GetUsersHandler)
    mux.HandleFunc("/users/create", handlers.CreateUserHandler)
    mux.HandleFunc("/users/update", handlers.UpdateUserHandler)
    mux.HandleFunc("/users/delete", handlers.DeleteUserHandler)

    return &Server{
        Server: http.Server{
            Addr:    cfg.Port,
            Handler: mux,
        },
    }
}
