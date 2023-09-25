package main

import (
	"fmt"
	"os"
	"url-shorner/internal/config"

	"golang.org/x/exp/slog"
)

// env types
const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	//TODO
	//init config : cleanenv
	cfg := config.MustLoad() //load config
	fmt.Println(cfg)         //make sure config is valid todo validate config and add config to vault
	//logger: slog
	log := setupLogger(cfg.Env)                                                   //get config env and setup logger
	defer log.Info("Hello World! im running at", slog.String("host", cfg.Adress)) //adress
	log.Info("starting up the shorner service", slog.String("env", cfg.Env))      //start service log
	log.Debug("Debug logs enabled")
	//storage: sqllite
	//router: gin go-chi, render
	//server
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return log
}
