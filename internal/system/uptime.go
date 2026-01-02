
package system

import (
	"os"
	"strconv"
	"strings"
	"time"
)

func Uptime() (string, error) {
	data, err := os.ReadFile("/proc/uptime")
	if err != nil {
		return "", err
	}

	seconds, _ := strconv.ParseFloat(strings.Fields(string(data))[0], 64)
	d := time.Duration(seconds) * time.Second

	h := int(d.Hours())
	m := int(d.Minutes()) % 60

	return strconv.Itoa(h) + "h " + strconv.Itoa(m) + "m", nil
}
