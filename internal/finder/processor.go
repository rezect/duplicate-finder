package finder

import (
	"fmt"
	"sync"

	"duplicate-finder/internal/cli"
)

type SameFiles struct {
	TotalSize  int64
	TotalFiles int64
	HashSum    string
	Files      []*FileData
}

func CompareFiles(sameSizedFiles map[int64][]*FileData, conf cli.Config, totalFiles int64) []*SameFiles {
	hashMap := make(map[string][]*FileData)
	allFiles := make([]*SameFiles, 0)
	var processedFilesCount int = 0

	fmt.Println()

	var wg sync.WaitGroup
	in := make(chan *FileData, conf.Workers)
	out := make(chan *FileData, conf.Workers)
	defer close(in)
	defer close(out)

	for range conf.Workers {
		go calculateHashParallel(conf, in, out)
	}

	for size, files := range sameSizedFiles {
		processedFilesCount += len(files)
		fmt.Printf("\r\033[KFilesProcessed: %d/%d", processedFilesCount, totalFiles)

		if len(files) > 1 {
			debugLogger(conf.Debug, fmt.Sprintf("Нашли файлы одинакового размера %d\n", size))

			go writeToHashMapParallel(&wg, out, &hashMap, conf)
			wg.Add(len(files))

			for _, file := range files {
				in <- file
			}

			wg.Wait()

			for hash, files := range hashMap {
				if len(files) > 1 {
					sameFiles := SameFiles{}
					sameFiles.HashSum = hash

					for _, file := range files {
						sameFiles.TotalSize += file.Size
						sameFiles.TotalFiles++
						sameFiles.Files = append(sameFiles.Files, file)
					}

					allFiles = append(allFiles, &sameFiles)
				}
			}

			for k := range hashMap {
				delete(hashMap, k)
			}
		}
	}

	fmt.Println()

	return allFiles
}

func writeToHashMapParallel(wg *sync.WaitGroup, out chan *FileData, hashMap *map[string][]*FileData, conf cli.Config) {
	for true {
		file, ok := <-out
		if !ok {
			break
		}

		if file.HashSum == "" {
			debugLogger(conf.Debug, fmt.Sprintf("Ошибка при обработке файла %s\n", file.Path))
			continue
		}
		hash := file.HashSum

		if _, exists := (*hashMap)[hash]; !exists {
			(*hashMap)[hash] = make([]*FileData, 0)
		}
		(*hashMap)[hash] = append((*hashMap)[hash], file)

		wg.Done()
	}
}
