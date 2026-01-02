package disk

import (
	"os"
	"path/filepath"
	"sort"
	"strconv"
)

type DirUsage struct {
	Path      string
	SizeBytes uint64
	SizeHuman string
}

func TopDirs(root string, limit int) ([]DirUsage, error) {
	entries, err := os.ReadDir(root)
	if err != nil {
		return nil, err
	}

	var dirs []DirUsage

	for _, e := range entries {
		if !e.IsDir() {
			continue
		}

		path := filepath.Join(root, e.Name())
		size := uint64(0)

		filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
			if err == nil && !info.IsDir() {
				size += uint64(info.Size())
			}
			return nil
		})

		dirs = append(dirs, DirUsage{
			Path:      path,
			SizeBytes: size,
			SizeHuman: humanSize(size),
		})
	}

	sort.Slice(dirs, func(i, j int) bool {
		return dirs[i].SizeBytes > dirs[j].SizeBytes
	})

	if len(dirs) > limit {
		dirs = dirs[:limit]
	}

	return dirs, nil
}

func humanSize(b uint64) string {
	const gb = 1024 * 1024 * 1024
	const mb = 1024 * 1024

	if b >= gb {
		return format(b, gb, "G")
	}
	return format(b, mb, "M")
}

func format(b, unit uint64, suffix string) string {
	return strconv.FormatFloat(float64(b)/float64(unit), 'f', 1, 64) + suffix
}
