package finder

import (
	"fmt"

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

	for size, files := range sameSizedFiles {
		processedFilesCount += len(files)
		fmt.Printf("\r\033[KFilesProcessed: %d/%d", processedFilesCount, totalFiles)

		if len(files) > 1 {
			debugLogger(conf.Debug, fmt.Sprintf("Нашли файлы одинакового размера %d\n", size))

			for _, file := range files {
				hash, err := calculateHash(file.Path, conf)
				if err != nil {
					fmt.Printf("Ошибка при получении хеша файла %s\n", file.Path)
					continue
				}
				file.HashSum = hash

				if _, exists := hashMap[hash]; !exists {
					hashMap[hash] = make([]*FileData, 0)
				}
				hashMap[hash] = append(hashMap[hash], file)
			}

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
