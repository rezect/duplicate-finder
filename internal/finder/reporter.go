package finder

import (
	"fmt"
	"sort"
)

func MakeReport(comparedFiles []*SameFiles) {
	fmt.Printf("Вот краткий итог по найденным файлам:\n")

	fmt.Printf("\nВсего найдено групп файлов: %d\n", len(comparedFiles))
	sort.Slice(comparedFiles, func(i, j int) bool {
		return comparedFiles[i].TotalSize >= comparedFiles[j].TotalSize
	})

	fmt.Println()

	for i, fileGroup := range comparedFiles {
		fmt.Printf("Группа %v (%x)\n", i+1, fileGroup.HashSum)
		fmt.Printf("Файлов: %d шт.; Общий размер: %d\n", fileGroup.TotalFiles, fileGroup.TotalSize)
		for _, file := range fileGroup.Files {
			fmt.Printf("\t- %s\n", file.Path)
		}
	}
}

func debugLogger(isDebugging bool, s string) {
	if isDebugging {
		fmt.Println(s)
	}
}
