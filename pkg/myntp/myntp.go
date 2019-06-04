package myntp

import (
	"time"

	bntp "github.com/beevik/ntp"
)

// Now will return current time according to "time.apple.com" server
func Now() (time.Time, error) {
	timeNow, err := bntp.Time("time.apple.com")
	return timeNow, err
}