package app

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"saffronreal.estate/backend/internal/handler"
	"saffronreal.estate/backend/internal/repository"
	"saffronreal.estate/backend/internal/service"
)

func New(pool *pgxpool.Pool) http.Handler {
	userRepository := repository.NewUserRepository(pool)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	mux := http.NewServeMux()

	mux.HandleFunc(
		"POST /api/v1/users",
		userHandler.Create,
	)

	mux.HandleFunc(
		"GET /api/v1/users/{id}",
		userHandler.GetByEmail,
	)

	var root http.Handler = mux

	return root
}
