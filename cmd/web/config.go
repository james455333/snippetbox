package main

import (
	"flag"
)

type Config struct {
	addr         string
	staticSrcDir string
	flag1        bool
}

func InitConfig() *Config {
	var cfg Config
	//addr := flag.String("addr", ":4000", "HTTP network address")
	//addr := os.Getenv("SNIPPETBOX_ADDR")
	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.staticSrcDir, "staticSrcDir", StaticSrcRootPath, "Path to static assets")
	//boolFlag1 := flag.Bool("flag1", true, "test boolean flag")
	flag.BoolVar(&cfg.flag1, "flag1", true, "test boolean flag")
	flag.Parse()
	return &cfg
}
