package finder

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"duplicate-finder/internal/cli"
)

type FileData struct {
	Path    string
	Size    int64
	Ext     string
	HashSum string
}

func ScanDirectory(conf cli.Config) (map[int64][]*FileData, []string) {
	sameSizedFiles := make(map[int64][]*FileData)
	scannedDirs := make([]string, 0)

	filepath.WalkDir(conf.Dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("Ошибка доступа к %s: %v\n", path, err)
			return err
		}

		if d.IsDir() {
			if l := strings.Split(path, string(filepath.Separator)); len(l) == 1 {
				scannedDirs = append(scannedDirs, path)
			}
		} else {
			info, err := d.Info()
			if err != nil {
				return err
			}
			size := info.Size()

			if size < conf.MinSize {
				return nil
			}

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

	return sameSizedFiles, scannedDirs
}
