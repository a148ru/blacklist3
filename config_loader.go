package main

import (
	"flag"
	"os"
	"strconv"
)

func loadConfigFromAllSources() *Config {
	configPath := getEnvOrFlag("CONFIG_PATH", "config", "config.yaml")

	cfg := loadConfig(configPath)

	overrideBool(&cfg.Daemon.Enabled, "DAEMON_ENABLED", "daemon")
	overrideInt(&cfg.Daemon.IntervalSeconds, "DAEMON_INTERVAL", "interval")
	overrideString(&cfg.OutputDir, "OUTPUT_DIR", "output")
	overrideString(&cfg.MD5File, "MD5_FILE", "md5")
	overrideInt(&cfg.HTTP.TimeoutSeconds, "HTTP_TIMEOUT", "http-timeout")

	return cfg
}

func getEnvOrFlag(env, flagName, def string) string {
	val := os.Getenv(env)
	if val != "" {
		return val
	}
	return def
}

func overrideString(target *string, env, flagName string) {
	flagValue := flag.String(flagName, "", "")
	flag.Parse()

	if *flagValue != "" {
		*target = *flagValue
		return
	}

	if v := os.Getenv(env); v != "" {
		*target = v
	}
}

func overrideBool(target *bool, env, flagName string) {
	flagValue := flag.Bool(flagName, *target, "")
	flag.Parse()

	*target = *flagValue

	if v := os.Getenv(env); v != "" {
		if b, err := strconv.ParseBool(v); err == nil {
			*target = b
		}
	}
}

func overrideInt(target *int, env, flagName string) {
	flagValue := flag.Int(flagName, *target, "")
	flag.Parse()

	*target = *flagValue

	if v := os.Getenv(env); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			*target = i
		}
	}
}
