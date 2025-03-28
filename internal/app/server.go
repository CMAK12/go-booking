package app

import (
	"log"
	"net/http"

	"go-booking/internal/config"
	v1 "go-booking/internal/controllers/http/v1"
	service "go-booking/internal/services"
	"go-booking/internal/storage"
	"go-booking/pkg/db"

	"github.com/go-chi/chi/v5"
)

func Run() {
	cfg := config.Load()

	db.InitializePostgreSQL(&cfg.DatabaseURL)
	defer db.Close()

	userStorage := storage.NewUserStorage(db.DB)
	reservationStorage := storage.NewReservationStorage(db.DB)

	userService := service.NewUserService(userStorage)
	reservationService := service.NewReservationService(reservationStorage)

	handler := v1.NewHandler(userService, reservationService)

	router := chi.NewRouter()

	handler.SetupRoutes(router)

	if err := http.ListenAndServe(cfg.ServerAddr, router); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
