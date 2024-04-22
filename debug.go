package main

import (
	"fmt"
	"time"
)

type debugInfo struct {
	start time.Time
}

func (d debugInfo) String() string {
	return fmt.Sprintf(
		"(took: %d μs)",
		time.Since(d.start).Microseconds(),
	)
}

var debug debugInfo
