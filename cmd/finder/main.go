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

	sameSizedFiles, scannedDirs := finder.ScanDirectory(conf)

	comparedFiles := finder.CompareFiles(sameSizedFiles, conf)

	finder.MakeReport(comparedFiles, scannedDirs)
}
