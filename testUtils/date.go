package testUtils

import (
	"math/rand"
	"time"
)

func Randate() time.Time {
	return time.Unix(rand.Int63n(2^60), 0)
}
