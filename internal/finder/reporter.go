package finder

import (
	"fmt"
)

func MakeReport(comparedFiles []*SameFiles, scannedDirs []string) {
	fmt.Printf("Вот краткий итог по найденным файлам:\n")

	fmt.Printf("Всего просканировано папок: %d\n", len(scannedDirs))
	for _, dirName := range scannedDirs {
		fmt.Printf("\t- %s\n", dirName)
	}

	fmt.Printf("\nВсего найдено групп файлов: %d\n", len(comparedFiles))
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
