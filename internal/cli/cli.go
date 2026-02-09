package cli

import (
	"flag"
	"fmt"
	"runtime"
)

type Config struct {
	Dir     string
	Algo    string
	MinSize int64
	Debug   bool
	Workers int
}

func ParseConfig() (Config, error) {
	conf := Config{}

	flag.StringVar(&conf.Dir, "dir", ".", "Directory to scan")
	flag.StringVar(&conf.Algo, "algo", "md5", "Ango to scan with")
	flag.Int64Var(&conf.MinSize, "min-size", 1, "Minimal size of files to scan")
	flag.BoolVar(&conf.Debug, "debug", false, "Debug mode")
	flag.IntVar(&conf.Workers, "workers", runtime.NumCPU() * 2, "Minimal size of files to scan")

	flag.Parse()

	if alg := conf.Algo; alg != "md5" && alg != "sha256" {
		return Config{}, fmt.Errorf("Поддерживаемые алгоритмы: 'sha256', 'md5'")
	}
	if workers := conf.Workers; workers > 1000 || workers < 1 {
		return Config{}, fmt.Errorf("Укажите правильное кол-во рабочих процессов (0 < workers < 1000)")
	}

	return conf, nil
}
