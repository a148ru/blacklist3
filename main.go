package main

import (
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

func main() {
	initLogger()
	cfg := loadConfigFromAllSources()

	//initHTTP(cfg.HTTP.TimeoutSeconds)
	initHTTP(cfg.HTTP)
	
	if cfg.Daemon.Enabled {
		runDaemon(cfg)
		return
	}

	runOnce(cfg)
}

func loadConfig(path string) *Config {
	f, _ := os.Open(path)
	defer f.Close()
	var cfg Config
	yaml.NewDecoder(f).Decode(&cfg)
	return &cfg
}

func runOnce(cfg *Config) {
	need_reset := true

	os.MkdirAll(cfg.OutputDir, 0755)

	md5store := loadMD5(cfg.MD5File)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, src := range cfg.Sources {
		wg.Add(1)

		go func(src Source, need_reset *bool) {
			defer wg.Done()

			logger.Println("Загрузка:", src.Name)
			data, err := loadSource(src)
			if err != nil {
				logger.Println("Ошибка:", err)
				return
			}

			sum := md5sum(data)

			mu.Lock()
			old := md5store[src.Name]
			mu.Unlock()

			if old == sum {
				logger.Println("Без изменений:", src.Name)
				*need_reset = false
				return
			}

			logger.Println("Обновление:", src.Name)
			processData(data, cfg.OutputDir, src.Name)
			mu.Lock()
			md5store[src.Name] = sum
			mu.Unlock()
		}(src, &need_reset)
	}

	wg.Wait()
	saveMD5(cfg.MD5File, md5store)

	if cfg.Service.Name != "" && need_reset == true {
		err := srvrestarter(cfg.Service.Name)
		if err != nil {
			logger.Println(err)
		}
	}
}
