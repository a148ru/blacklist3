package main

import (
	"flag"
	"os"
	"strconv"
)

func loadConfigFromAllSources() *Config {
	configPath := getEnvOrFlag("CONFIG_PATH", "config", "config.yaml")
	cfg := loadConfig(configPath)

	flagInterval := flag.Int("interval", 0, "set the file update check interval")
	flagOutput := flag.String("output","","specify the directory to write files")
	flagMD5File := flag.String("md5","","specify a file to store checksums")
	flagHttpTimeOut := flag.Int("http-timeout",0,"specify http timeout")
	flagDaemon := flag.Bool("daemon", false,"specify a flag to run as a daemon")
	flag.Parse()


	overrideInt(&cfg.Daemon.IntervalSeconds, "DAEMON_INTERVAL", "interval", flagInterval)
	overrideString(&cfg.OutputDir, "OUTPUT_DIR", "output", flagOutput)
	overrideString(&cfg.MD5File, "MD5_FILE", "md5", flagMD5File)
	overrideInt(&cfg.HTTP.TimeoutSeconds, "HTTP_TIMEOUT", "http-timeout", flagHttpTimeOut)
	overrideBool(&cfg.Daemon.Enabled, "DAEMON_ENABLED", "daemon", flagDaemon)

	return cfg
}

func getEnvOrFlag(env, flagName, def string) string {
	val := os.Getenv(env)
	if val != "" {
		return val
	}
	return def
}

func overrideString(target *string, env, flagName string, flagValue *string) {


	if *flagValue != "" {
		*target = *flagValue
		return
	}

	if v := os.Getenv(env); v != "" {
		*target = v
	}
}

func overrideBool(target *bool, env, flagName string, flagValue *bool) {

	if *flagValue {
		*target = *flagValue
		return
	}
	if v := os.Getenv(env); v != "" {
		if b, err := strconv.ParseBool(v); err == nil {
			*target = b
		}
	}
}

func overrideInt(target *int, env, flagName string, flagValue *int) {

	if *flagValue != 0{
		*target = *flagValue
		return
	}

	if v := os.Getenv(env); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			*target = i
		}
	}
}
