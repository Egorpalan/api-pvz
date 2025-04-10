package main

import (
	"github.com/Egorpalan/api-pvz/config"
	"github.com/Egorpalan/api-pvz/internal/handler"
	"github.com/Egorpalan/api-pvz/internal/middleware"
	"github.com/Egorpalan/api-pvz/internal/repository"
	"github.com/Egorpalan/api-pvz/internal/usecase"
	"github.com/Egorpalan/api-pvz/pkg/db"
	"github.com/Egorpalan/api-pvz/pkg/logger"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()

	logger.Init()

	database := db.NewPostgresDB(cfg.DB)
	defer database.Close()

	// TODO: Init router, routes, server handler

	pvzRepo := repository.NewPVZRepository(database)
	pvzUC := usecase.NewPVZUsecase(pvzRepo)

	receptionRepo := repository.NewReceptionRepository(database)
	receptionUC := usecase.NewReceptionUsecase(receptionRepo)

	productRepo := repository.NewProductRepository(database)
	productUC := usecase.NewProductUsecase(productRepo)

	userRepo := repository.NewUserRepository(database)
	authUC := usecase.NewAuthUsecase(userRepo)

	r := chi.NewRouter()
	r.Post("/dummyLogin", handler.DummyLogin)
	r.Post("/register", handler.Register(authUC))
	r.Post("/login", handler.Login(authUC))

	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware("moderator"))
		r.Post("/pvz", handler.CreatePVZ(pvzUC))
	})

	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware("client"))

		r.Post("/receptions", handler.CreateReception(receptionUC))
		r.Post("/products", handler.CreateProduct(productUC))
		r.Post("/pvz/{pvzId}/delete_last_product", handler.DeleteLastProduct(productUC))
		r.Post("/pvz/{pvzId}/close_last_reception", handler.CloseLastReception(receptionUC))
	})

	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware("moderator", "employee"))
		r.Get("/pvz", handler.GetPVZList(pvzUC))
	})

	err := http.ListenAndServe(":"+cfg.AppPort, r)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
	log.Printf("Starting server on port %s...", cfg.AppPort)

}
