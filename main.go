package main

import (
	"log/slog"
	"service-advert/internal/config"
	"service-advert/internal/database/pgsql"
	"service-advert/internal/server"
)

func main() {
	cfg, err := config.New(`configs/main.yaml`)
	if err != nil {
		return
	}

	err = pgsql.Init(cfg)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	defer pgsql.CloseDB()

	// err = redis.Init(cfg)
	// if err != nil {
	// 	slog.Error(err.Error())
	// 	return
	// }
	// defer redis.CloseDB()

	server.Start(cfg)
}
