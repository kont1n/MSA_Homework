package main

import (
	"log/slog"

	"MSA_Homework/internal/utils/config"
	"MSA_Homework/internal/utils/logger"
	
)

var cfg = config.Config{}

func init() {
	cfg.Env = "local"	
	// TODO: загрузить конфигурацию из файла конфигурации
}

func main() {
	log := logger.SetupLogger(cfg.Env)

	log.Info("Starting application",
		slog.String("env", cfg.Env),
	)

	// TODO: запустить сервис авторизации
}
