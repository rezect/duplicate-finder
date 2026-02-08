package cli

import (
	"flag"
	"fmt"
)

type Config struct {
	Dir     string
	Algo    string
	MinSize int64
	Debug	bool
}

func ParseConfig() (Config, error) {
	conf := Config{}

	flag.StringVar(&conf.Dir, "dir", ".", "Directory to scan")
	flag.StringVar(&conf.Algo, "algo", "md5", "Ango to scan with")
	flag.Int64Var(&conf.MinSize, "min-size", 1, "Minimal size of files to scan")
	flag.BoolVar(&conf.Debug, "debug", false, "Debug mode")

	flag.Parse()

	if alg := conf.Algo; alg != "md5" && alg != "sha256" {
		return Config{}, fmt.Errorf("Поддерживаемые алгоритмы: 'sha256', 'md5'")
	}

	return conf, nil
}
