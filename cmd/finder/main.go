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

	sameSizedFiles, totalFiles := finder.ScanDirectory(conf)

	comparedFiles := finder.CompareFiles(sameSizedFiles, conf, totalFiles)

	finder.MakeReport(comparedFiles)
}
