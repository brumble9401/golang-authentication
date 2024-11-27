package api

import (
	"database/sql"
	"net/http"

	"github.com/brumble9401/golang-authentication/services/user"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db: db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subRouter)

	log.Info("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}