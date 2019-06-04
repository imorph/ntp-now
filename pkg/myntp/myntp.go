package myntp

import (
	"time"
)

// Now will return current time according to 0.pool.ntp.org server
func Now() (time.Time) {
	timeNow := time.Now()
	return timeNow
}