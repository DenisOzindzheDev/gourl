package main

import (
	"fmt"
	"url-shorner/internal/config"
)

func main() {
	//TODO
	//init config : cleanenv
	cfg := config.MustLoad()
	fmt.Println(cfg)
	//logger: slog
	//storage: sqllite
	//router: gin go-chi, render
	//server
}
