package finder

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"duplicate-finder/internal/cli"
)

func ScanDirectory(conf cli.Config) map[int64][]string {
	sameSizedFiles := make(map[int64][]string)

	filepath.WalkDir(conf.Dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("Ошибка доступа к %s: %v\n", path, err)
			return err
		}

		if d.IsDir() {
			fmt.Printf("Директория: %s\n", path)
		} else {
			info, err := d.Info()
			if err != nil {
				return err
			}
			size := info.Size()

			if size < conf.MinSize {
				return nil
			}

			if _, exists := sameSizedFiles[size]; !exists {
				sameSizedFiles[size] = make([]string, 0)
			}
			sameSizedFiles[size] = append(sameSizedFiles[size], path)
		}

		return nil
	})

	return sameSizedFiles
}
