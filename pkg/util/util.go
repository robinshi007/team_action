package util

import (
	"time"
)

// TimeNowStr -
func TimeNowStr() string {
	return time.Now().UTC().Format(time.RFC3339Nano)
}
