package main

import (
	"duplicate-finder/internal/cli"
	"duplicate-finder/internal/finder"
)

func main() {
	conf, err := cli.ParseConfig()
	if err != nil {
		panic(err)
	}

	sameSizedFiles := finder.ScanDirectory(conf)

	finder.CompareFiles(sameSizedFiles, conf)
}
