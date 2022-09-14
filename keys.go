package cobman

import (
	"github.com/google/uuid"
)

var (
	keyCmdDescription = makeKey()
	keyCmdExample     = makeKey()
	keyFlagValueName  = makeKey()
)

func makeKey() string {
	return uuid.New().String() + "-" + uuid.New().String()
}
