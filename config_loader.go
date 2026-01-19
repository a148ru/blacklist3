package main

import (
	"flag"
	"os"
	"strconv"
)

func loadConfigFromAllSources() *Config {

	flagConfigPath := flag.String("config", "", "set path to config file")
	flagInterval := flag.Int("interval", 0, "set the file update check interval")
	flagOutput := flag.String("output", "", "specify the directory to write files")
	flagMD5File := flag.String("md5", "", "specify a file to store checksums")
	flagHttpTimeOut := flag.Int("http-timeout", 0, "specify http timeout")
	flagDaemon := flag.Bool("daemon", false, "specify a flag to run as a daemon")
	flagSevise := flag.String("service", "", "Specify the name of the service that needs to be restarted after updating the data.")
	flag.Parse()

	configPath := getEnvOrFlag("CONFIG_PATH", flagConfigPath, "config.yaml")
	cfg := loadConfig(configPath)

	overrideString(&cfg.OutputDir, "OUTPUT_DIR", flagOutput)
	overrideString(&cfg.MD5File, "MD5_FILE", flagMD5File)
	overrideInt(&cfg.Daemon.IntervalSeconds, "DAEMON_INTERVAL", flagInterval)
	overrideInt(&cfg.HTTP.TimeoutSeconds, "HTTP_TIMEOUT", flagHttpTimeOut)
	overrideBool(&cfg.Daemon.Enabled, "DAEMON_ENABLED", flagDaemon)
	overrideString(&cfg.Service.Name, "SERVICE_NAME_RELOADS", flagSevise)

	return cfg
}

func getEnvOrFlag(env string, flagName *string, def string) string {

	if *flagName != "" {
		return *flagName
	}

	val := os.Getenv(env)
	if val != "" {
		return val
	}

	return def
}

func overrideString(target *string, env string, flagValue *string) {

	if *flagValue != "" {
		*target = *flagValue
		return
	}

	if v := os.Getenv(env); v != "" {
		*target = v
	}
}

func overrideBool(target *bool, env string, flagValue *bool) {

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

func overrideInt(target *int, env string, flagValue *int) {

	if *flagValue != 0 {
		*target = *flagValue
		return
	}

	if v := os.Getenv(env); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			*target = i
		}
	}
}
