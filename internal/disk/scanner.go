package disk

import (
	"os"
	"path/filepath"
)

type DirSize struct {
	Path string
	Size uint64
}

func Scan(root string, ignore map[string]bool) ([]DirSize, error) {
	var results []DirSize

	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if d.IsDir() && ignore[path] {
			return filepath.SkipDir
		}

		return nil
	})

	return results, err
}
