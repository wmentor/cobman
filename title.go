package cobman

import (
	"errors"
	"runtime"
	"time"
)

const (
	timeFormat = "2006-01-02"
)

var (
	ErrInvalidSection = errors.New("invalid section number")

	thSection = 1
	thOS      = runtime.GOOS
	thDate    = time.Now().UTC().Format(timeFormat)
	thTitle   = "[TITLE]"
)

func SetSection(secNum int) error {
	if secNum >= 1 && secNum <= 9 {
		thSection = secNum
		return nil
	}
	return ErrInvalidSection
}

func SetOS(name string) {
	thOS = name
}

func SetDate(dt string) {
	thDate = dt
}

func SetTitle(title string) {
	thTitle = title
}
