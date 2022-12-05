package base

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"time"
)

type App struct {
	Router *gin.Engine
	DB     *sqlx.DB
	Server *http.Server
}

func (a *App) Init() {
	a.Router = gin.New()
	a.DB = ConnectToDb()
}

func (a *App) Run(port string, chErr chan error) {
	a.Server = &http.Server{
		Handler:      a.Router,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	defer a.Server.Close()

	log.Printf("[DEBUG] Running server on port %s", port)

	if err := a.Server.ListenAndServe(); err != nil {
		chErr <- errors.New(fmt.Sprintf("Error server: %s", err.Error()))
	}
}

func (a *App) Close() {
	if err := a.DB.Close(); err != nil {
		log.Fatalf("[FATAL] Unable to close database: %v", err)
		return
	}

	if err := a.Server.Close(); err != nil {
		log.Fatalf("[FATAL] Unable to close server: %v", err)
		return
	}
}

func ConnectToDb() *sqlx.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("[FATAL] Not loading environment: %v", err)
	}

	db, err := sqlx.Connect("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	))

	if err != nil {
		log.Fatalf("[FATAL] Unable to connect to database: %v", err)
	}

	return db
}
