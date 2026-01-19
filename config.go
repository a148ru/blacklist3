package main

type Config struct {
	Sources   []Source   `yaml:"sources"`
	MD5File   string     `yaml:"md5_file"`
	OutputDir string     `yaml:"output_dir"`
	Daemon    DaemonConf `yaml:"daemon"`
	HTTP      HTTPConf   `yaml:"http"`
	Service   Service    `yaml:"service"`
}

type Source struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"` // url | file
	Path string `yaml:"path"`
}

type DaemonConf struct {
	Enabled         bool `yaml:"enabled"`
	IntervalSeconds int  `yaml:"interval_seconds"`
}

type HTTPConf struct {
	TimeoutSeconds int `yaml:"timeout_seconds"`
}

type Service struct {
	Name string `yaml:"name"`
}
