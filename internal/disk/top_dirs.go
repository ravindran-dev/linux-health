package disk

import (
	"bytes"
	"os/exec"
	"sort"
	"strings"
)

type DirUsage struct {
	Path      string
	SizeHuman string
}

func TopDirs(root string, limit int) ([]DirUsage, error) {
	cmd := exec.Command("du", "-sh", root+"/*")
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return nil, err
	}

	lines := strings.Split(out.String(), "\n")
	var dirs []DirUsage

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}

		dirs = append(dirs, DirUsage{
			SizeHuman: parts[0],
			Path:      parts[1],
		})
	}

	// Sort descending by size string (du already sorts roughly)
	sort.Slice(dirs, func(i, j int) bool {
		return dirs[i].SizeHuman > dirs[j].SizeHuman
	})

	if len(dirs) > limit {
		dirs = dirs[:limit]
	}

	return dirs, nil
}
