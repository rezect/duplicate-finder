package finder

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"duplicate-finder/internal/cli"
)

type FileData struct {
	Path    string
	Size    int64
	Ext     string
	HashSum string
}

func ScanDirectory(conf cli.Config) (map[int64][]*FileData, int64) {
	sameSizedFiles := make(map[int64][]*FileData)
	var totalFiles int64 = 0

	filepath.WalkDir(conf.Dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("Ошибка доступа к %s: %v\n", path, err)
			return err
		}

		if d.IsDir() {
			debugLogger(conf.Debug, fmt.Sprintf("Scanning directory: %s", path))
		} else {
			info, err := d.Info()
			if err != nil {
				return err
			}
			size := info.Size()
			
			if size < conf.MinSize {
				return nil
			}
			
			totalFiles++
			fmt.Printf("\r\033[KTotalFiles: %d", totalFiles)
			
			var curFile = FileData{}
			curFile.Path = path
			curFile.Size = size
			curFile.Ext = filepath.Ext(path)

			if _, exists := sameSizedFiles[size]; !exists {
				sameSizedFiles[size] = make([]*FileData, 0)
			}
			sameSizedFiles[size] = append(sameSizedFiles[size], &curFile)
		}

		return nil
	})

	return sameSizedFiles, totalFiles
}
