package finder

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"hash"
	"io"
	"os"
	"time"

	"duplicate-finder/internal/cli"
)

func calculateHash(path string, conf cli.Config) (string, error) {
	debugLogger(conf.Debug, fmt.Sprintf("Начали считать хеш для файла %s\n", path))

	start := time.Now()
	file, err := os.Open(path)

	if err != nil {
		return "", err
	}
	defer file.Close()

	var h hash.Hash

	switch conf.Algo {
	case "md5":
		h = md5.New()
	case "sha256":
		h = sha256.New()
	default:
		return "", fmt.Errorf("Unsupported type of algorithm: %s", conf.Algo)
	}
	if _, err := io.Copy(h, file); err != nil {
		return "", err
	}

	if conf.Debug {
		timeLog(start, "calculateHash")
	}

	return string(h.Sum(nil)), nil
}

func calculateHashParallel(conf cli.Config, in chan *FileData, out chan *FileData) {
	for true {
		fileData, ok := <-in
		if !ok {
			return
		}

		debugLogger(conf.Debug, fmt.Sprintf("Начали считать хеш для файла %s\n", fileData.Path))
		start := time.Now()
		file, err := os.Open(fileData.Path)
	
		if err != nil {
			out <- fileData
			continue
		}
		defer file.Close()
	
		var h hash.Hash
	
		switch conf.Algo {
		case "md5":
			h = md5.New()
		case "sha256":
			h = sha256.New()
		default:
			out <- fileData
			continue
		}
		if _, err := io.Copy(h, file); err != nil {
			out <- fileData
			continue
		}
	
		if conf.Debug {
			timeLog(start, "calculateHash")
		}
	
		fileData.HashSum = string(h.Sum(nil))
	
		out <- fileData
	}
}

func timeLog(start time.Time, funcName string) {
	t := time.Now()
	elapsed := t.Sub(start)

	fmt.Printf("[TIME] Время выполнения %s - %v\n", funcName, elapsed)
}
