package app

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	"go-booking/internal/config"
	v1 "go-booking/internal/controllers/http/v1"
	service "go-booking/internal/services"
	"go-booking/internal/storage"
	redis "go-booking/pkg/caching"
	"go-booking/pkg/db"

	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose"
)

func MustRun() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.MustLoad()

	mailAuth := cfg.PostSender.AuthenticateMailer()

	pgConn, dsn := db.MustConnectPostgres(ctx, cfg.Postgres)
	migrate(dsn)

	redisDB := redis.MustConnectRedis(ctx, cfg.Redis)

	_, shutdownCancel := context.WithTimeout(ctx, 10*time.Second)
	defer shutdownCancel()

	userStorage := storage.NewUserStorage(pgConn)
	bookingStorage := storage.NewBookingStorage(pgConn)
	roomStorage := storage.NewRoomStorage(pgConn)
	hotelStorage := storage.NewHotelStorage(pgConn)
	extraServiceStorage := storage.NewExtraServiceStorage(pgConn)

	userService := service.NewUserService(userStorage)
	hotelService := service.NewHotelService(hotelStorage)
	roomService := service.NewRoomService(roomStorage, extraServiceStorage, redisDB)
	extraServiceService := service.NewExtraServiceService(extraServiceStorage)
	bookingService := service.NewBookingService(
		bookingStorage,
		hotelStorage,
		roomService,
		userStorage,
		extraServiceStorage,
		mailAuth,
	)

	handler := v1.NewHandler(
		userService,
		bookingService,
		hotelService,
		roomService,
		extraServiceService,
	)

	router := chi.NewRouter()
	handler.SetupRoutes(router)

	if err := http.ListenAndServe(cfg.HTTP.Port, router); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func migrate(dsn string) {
	sqlDB, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("Failed to open SQL connection: %v", err)
	}

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("Failed to set goose dialect: %v", err)
	}

	if err := goose.Up(sqlDB, "./migrations"); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
}
