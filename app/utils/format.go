package utils

import (
	"fmt"
	"time"
)

func FormatDuration(d time.Duration) string {
	var units = []struct {
		unit  string
		value int64
	}{
		{"year", int64(time.Hour * 24 * 365)},
		{"month", int64(time.Hour * 24 * 30)},
		{"week", int64(time.Hour * 24 * 7)},
		{"day", int64(time.Hour * 24)},
		{"hour", int64(time.Hour)},
		{"min", int64(time.Minute)},
		{"sec", int64(time.Second)},
	}

	var output string
	remaining := int64(d)
	for _, u := range units {
		if remaining >= u.value {
			output = fmt.Sprintf("%d %s", remaining/u.value, u.unit)
			if remaining/u.value > 1 {
				output += "s"
			}
			break
		}
	}

	if output == "" {
		output = "0 sec"
	}

	output += " ago"

	return output
}
