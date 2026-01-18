package main

import "time"

func runDaemon(cfg *Config) {
	logger.Println("Запуск в режиме демона")

	ticker := time.NewTicker(time.Duration(cfg.Daemon.IntervalSeconds) * time.Second)
	defer ticker.Stop()

	for {
		runOnce(cfg)
		<-ticker.C
	}
}
