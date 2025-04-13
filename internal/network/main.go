package main

func init() {
	// TODO: загрузить конфигурацию из файла конфигурации
}

func main() {
	log := logger.SetupLogger(cfg.Env)
	defer log.Sync()

	// TODO: запустить сервис связей
}
