package formatter

import "time"

func CurrentTimestamp() string {
	return time.Now().UTC().Format(time.RFC3339)
}
