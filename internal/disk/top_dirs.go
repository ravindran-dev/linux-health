package disk

import (
	"bytes"
	"os/exec"
	"strings"
)

type DirUsage struct {
	Path      string
	SizeHuman string
}

func TopDirs(root string, limit int) ([]DirUsage, error) {
	cmd := exec.Command("sh", "-c", "du -sh "+root+"/* 2>/dev/null")
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return nil, err
	}

	lines := strings.Split(out.String(), "\n")
	var dirs []DirUsage

	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) >= 2 {
			dirs = append(dirs, DirUsage{
				SizeHuman: fields[0],
				Path:      fields[1],
			})
		}
	}

	if len(dirs) > limit {
		dirs = dirs[:limit]
	}

	return dirs, nil
}
