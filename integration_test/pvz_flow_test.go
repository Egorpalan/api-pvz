package integration_test

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/Egorpalan/api-pvz/config"
	"github.com/Egorpalan/api-pvz/internal/repository"
	"github.com/Egorpalan/api-pvz/internal/usecase"
	"github.com/Egorpalan/api-pvz/pkg/db"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntegrationFlow(t *testing.T) {
	if err := startTestDatabase(); err != nil {
		t.Fatalf("Failed to start test database: %v", err)
	}
	defer stopTestDatabase()

	if err := waitForDatabase(); err != nil {
		t.Fatalf("Database is not ready: %v", err)
	}

	dbConfig := config.DBConfig{
		Host:     "localhost",
		Port:     "5437",
		User:     "postgres",
		Password: "postgres",
		Name:     "pvz_test_db",
		SSLMode:  "disable",
	}

	database := db.NewPostgresDB(dbConfig)
	defer database.Close()

	if err := initTestDatabaseSchema(database); err != nil {
		t.Fatalf("Failed to initialize database schema: %v", err)
	}

	pvzRepo := repository.NewPVZRepository(database)
	receptionRepo := repository.NewReceptionRepository(database)
	productRepo := repository.NewProductRepository(database)

	pvzUsecase := usecase.NewPVZUsecase(pvzRepo)
	receptionUsecase := usecase.NewReceptionUsecase(receptionRepo)
	productUsecase := usecase.NewProductUsecase(productRepo)

	pvz, err := pvzUsecase.Create("Москва")
	require.NoError(t, err)
	assert.NotEmpty(t, pvz.ID)
	assert.Equal(t, "Москва", pvz.City)

	reception, err := receptionUsecase.Create(pvz.ID)
	require.NoError(t, err)
	assert.NotEmpty(t, reception.ID)
	assert.Equal(t, pvz.ID, reception.PVZID)
	assert.Equal(t, "in_progress", reception.Status)

	for i := 0; i < 50; i++ {
		productType := "электроника"
		if i%3 == 1 {
			productType = "одежда"
		} else if i%3 == 2 {
			productType = "обувь"
		}

		product, err := productUsecase.Add(pvz.ID, productType)
		require.NoError(t, err)
		assert.NotEmpty(t, product.ID)
		assert.Equal(t, reception.ID, product.ReceptionID)
		assert.Equal(t, productType, product.Type)
	}

	closedReception, err := receptionUsecase.CloseLast(pvz.ID)
	require.NoError(t, err)
	assert.Equal(t, reception.ID, closedReception.ID)
	assert.Equal(t, "close", closedReception.Status)

	_, err = productUsecase.Add(pvz.ID, "электроника")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no active reception")
}

func initTestDatabaseSchema(db *sqlx.DB) error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS pvz (
			id VARCHAR(36) PRIMARY KEY,
			city VARCHAR(100) NOT NULL,
			registration_date TIMESTAMP WITH TIME ZONE NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS receptions (
			id VARCHAR(36) PRIMARY KEY,
			pvz_id VARCHAR(36) REFERENCES pvz(id),
			date_time TIMESTAMP WITH TIME ZONE NOT NULL,
			status VARCHAR(20) NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS products (
			id VARCHAR(36) PRIMARY KEY,
			reception_id VARCHAR(36) REFERENCES receptions(id),
			date_time TIMESTAMP WITH TIME ZONE NOT NULL,
			type VARCHAR(50) NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS users (
			id VARCHAR(36) PRIMARY KEY,
			email VARCHAR(100) UNIQUE NOT NULL,
			password VARCHAR(100) NOT NULL,
			role VARCHAR(20) NOT NULL
		)`,
	}

	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			return err
		}
	}

	return nil
}

func startTestDatabase() error {
	dockerComposeFile := getDockerComposeFilePath()
	cmd := exec.Command("docker-compose", "-f", dockerComposeFile, "up", "-d")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func stopTestDatabase() {
	dockerComposeFile := getDockerComposeFilePath()
	cmd := exec.Command("docker-compose", "-f", dockerComposeFile, "down", "-v")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func getDockerComposeFilePath() string {
	rootDir, err := getProjectRoot()
	if err != nil {
		log.Printf("Error finding project root: %v", err)
		return "../docker-compose.test.yml"
	}

	return filepath.Join(rootDir, "docker-compose.test.yml")
}

func getProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("could not find project root (go.mod file)")
		}
		dir = parent
	}
}

func waitForDatabase() error {
	dsn := "host=localhost port=5437 user=postgres password=postgres dbname=pvz_test_db sslmode=disable"
	var db *sql.DB
	var err error

	for i := 0; i < 30; i++ {
		db, err = sql.Open("postgres", dsn)
		if err == nil {
			err = db.Ping()
			if err == nil {
				db.Close()
				return nil
			}
		}

		log.Printf("Waiting for database... (%d/30)", i+1)
		time.Sleep(1 * time.Second)
	}

	return fmt.Errorf("database did not become ready in time")
}
