package server

import (
	"log/slog"
	"net/http"
	"os"
	api "service-advert/internal/api"
	"service-advert/internal/config"

	"github.com/rs/cors"
)

func Start(cfg *config.Config) {

	mux := http.NewServeMux()
	handler := cors.Default().Handler(mux)

	api.Init(mux, cfg)

	slog.Info("Start server")
	//err := http.ListenAndServe(":"+cfg.HTTP.Port, mux)
	err := http.ListenAndServe(":"+cfg.HTTP.Port, handler)

	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

}
