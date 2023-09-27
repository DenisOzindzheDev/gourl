package main

import (
	"os"
	"url-shorner/internal/config"
	"url-shorner/internal/storage/sqlite"
	"url-shorner/lib/logger/handlers/slogpretty"
	"url-shorner/lib/logger/sl"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-chi/chi/v5/middleware"
	"golang.org/x/exp/slog"
)

// env types
const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

// 43 24
func main() {
	//init config : cleanenv
	cfg := config.MustLoad() //load config fmt.Println(cfg) for make sure config is valid todo validate config and add config to vault
	//logger: slog
	log := setupLogger(cfg.Env)                                                   //get config env and setup logger
	defer log.Info("Hello World! im running at", slog.String("host", cfg.Adress)) //adress
	log.Info("starting up the shorner service", slog.String("env", cfg.Env))      //start service log
	log.Debug("Debug logs enabled")                                               //if env is local debug logs are enabled
	//storage: sqllite
	storage, err := sqlite.New(cfg.StoragePath) //create new storage
	if err != nil {
		log.Error("failed to connect to storage", sl.Err(err)) //if storage connection fails exit
		os.Exit(1)                                             //exit code 1 = error
	}
	_ = storage                                                            //todo remove!!!!!!!!!!!!!!!!!!!!!!!
	log.Info("Storage connected", slog.String("storage", cfg.StoragePath)) // success connecion log
	//TODO
	// router: gin go-chi, render
	router := chi.NewRouter()
	//todo custom MW
	router.Use(middleware.RequestID) //Мидлварина
	router.Use(middleware.Logger)    // Logger
	router.Use(middleware.Recoverer) // Recoverer
	router.Use(middleware.URLFormat) // URL format
	//auth middleware
	// server
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		//log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
		log = setupPrettySlog()
	case envDev:
		//log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
		log = setupPrettySlog()
	case envProd:
		//log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
		log = setupPrettySlog()
	}

	return log
}
func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}
	handler := opts.NewPrettyHandler(os.Stdout)
	return slog.New(handler)
}
