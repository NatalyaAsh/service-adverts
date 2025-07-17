package main

import (
	"log/slog"
	server "service-advert/internal/app/user/server"
	"service-advert/internal/config"
	"service-advert/internal/database/pgsql"
)

func main() {
	cfg, err := config.New(`configs/user.yaml`)
	if err != nil {
		return
	}

	err = pgsql.InitUser(cfg)
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
