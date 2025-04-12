package main

import (
	"context"
	"errors"
	"github.com/jmoiron/sqlx"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	generated "github.com/Egorpalan/api-pvz/api"
	"github.com/Egorpalan/api-pvz/config"
	grpcserver "github.com/Egorpalan/api-pvz/internal/grpc"
	"github.com/Egorpalan/api-pvz/internal/handler"
	"github.com/Egorpalan/api-pvz/internal/middleware"
	"github.com/Egorpalan/api-pvz/internal/repository"
	"github.com/Egorpalan/api-pvz/internal/usecase"
	"github.com/Egorpalan/api-pvz/pkg/db"
	"github.com/Egorpalan/api-pvz/pkg/logger"
	"github.com/go-chi/chi/v5"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
)

// APIServer реализует интерфейс generated.ServerInterface
type APIServer struct {
	authUC      *usecase.AuthUsecase
	pvzUC       *usecase.PVZUsecase
	productUC   *usecase.ProductUsecase
	receptionUC *usecase.ReceptionUsecase
}

// Реализация методов интерфейса ServerInterface
func (s *APIServer) PostDummyLogin(w http.ResponseWriter, r *http.Request) {
	handler.DummyLogin(w, r)
}

func (s *APIServer) PostLogin(w http.ResponseWriter, r *http.Request) {
	handler.Login(s.authUC)(w, r)
}

func (s *APIServer) PostRegister(w http.ResponseWriter, r *http.Request) {
	handler.Register(s.authUC)(w, r)
}

func (s *APIServer) PostPvz(w http.ResponseWriter, r *http.Request) {
	handler.CreatePVZ(s.pvzUC)(w, r)
}

func (s *APIServer) GetPvz(w http.ResponseWriter, r *http.Request, params generated.GetPvzParams) {
	handler.GetPVZList(s.pvzUC)(w, r)
}

func (s *APIServer) PostProducts(w http.ResponseWriter, r *http.Request) {
	handler.CreateProduct(s.productUC)(w, r)
}

func (s *APIServer) PostReceptions(w http.ResponseWriter, r *http.Request) {
	handler.CreateReception(s.receptionUC)(w, r)
}

func (s *APIServer) PostPvzPvzIdCloseLastReception(w http.ResponseWriter, r *http.Request, pvzId openapi_types.UUID) {
	ctx := context.WithValue(r.Context(), "pvzId", pvzId)
	handler.CloseLastReception(s.receptionUC)(w, r.WithContext(ctx))
}

func (s *APIServer) PostPvzPvzIdDeleteLastProduct(w http.ResponseWriter, r *http.Request, pvzId openapi_types.UUID) {
	ctx := context.WithValue(r.Context(), "pvzId", pvzId)
	handler.DeleteLastProduct(s.productUC)(w, r.WithContext(ctx))
}

func main() {
	cfg := config.LoadConfig()

	logger.Init()

	database := db.NewPostgresDB(cfg.DB)
	defer func(database *sqlx.DB) {
		err := database.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(database)

	pvzRepo := repository.NewPVZRepository(database)
	pvzUC := usecase.NewPVZUsecase(pvzRepo)

	receptionRepo := repository.NewReceptionRepository(database)
	receptionUC := usecase.NewReceptionUsecase(receptionRepo)

	productRepo := repository.NewProductRepository(database)
	productUC := usecase.NewProductUsecase(productRepo)

	userRepo := repository.NewUserRepository(database)
	authUC := usecase.NewAuthUsecase(userRepo)

	apiServer := &APIServer{
		authUC:      authUC,
		pvzUC:       pvzUC,
		productUC:   productUC,
		receptionUC: receptionUC,
	}

	r := chi.NewRouter()

	r.Use(middleware.PrometheusMiddleware)

	generated.HandlerWithOptions(apiServer, generated.ChiServerOptions{
		BaseRouter: r,
		Middlewares: []generated.MiddlewareFunc{
			func(next http.Handler) http.Handler {
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					next.ServeHTTP(w, r)
				})
			},
		},
	})

	metricsSrv := &http.Server{
		Addr: ":9000",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/metrics" {
				promhttp.Handler().ServeHTTP(w, r)
			} else {
				http.NotFound(w, r)
			}
		}),
	}

	go func() {
		log.Printf("Prometheus metrics available at :9000/metrics")
		if err := metricsSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Metrics server error: %s", err)
		}
	}()

	srv := &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: r,
	}

	go func() {
		log.Printf("Starting rest server on port %s...", cfg.AppPort)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server error: %s", err)
		}
	}()

	go func() {
		log.Printf("Starting grpc server on port %s...", cfg.GrpcPort)
		if err := grpcserver.RunGRPCServer(pvzUC); err != nil {
			log.Fatalf("grpc Server error: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Printf("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %s", err)
	}

	if err := metricsSrv.Shutdown(ctx); err != nil {
		log.Fatalf("Metrics server forced to shutdown: %s", err)
	}
}
