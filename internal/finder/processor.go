package finder

import (
	"fmt"

	"duplicate-finder/internal/cli"
)

func CompareFiles(sameSizedFiles map[int64][]string, conf cli.Config) {
	hashMap := make(map[string][]string)

	for size, paths := range sameSizedFiles {
		if len(paths) > 1 {
			fmt.Printf("Нашли файлы одинакового размера %d\n", size)
			for _, path := range paths {
				hash, err := calculateHash(path, conf)
				if err != nil {
					fmt.Printf("Ошибка при получении хеша файла %s\n", path)
					continue
				}

				if _, exists := hashMap[hash]; !exists {
					hashMap[hash] = make([]string, 0)
				}
				hashMap[hash] = append(hashMap[hash], path)
			}

			for hash, paths := range hashMap {
				if len(paths) > 1 {
					fmt.Printf("Нашли одинаковые файлы с хешем: %x\n", hash)
					for _, path := range paths {
						fmt.Printf("\t- %s\n", path)
					}
				}
			}

			for k := range hashMap {
				delete(hashMap, k)
			}
		}
	}
}
