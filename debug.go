package main

import (
	"fmt"
	"time"
)

type debugInfo struct {
	start         time.Time
	amountAliases int
	amountHistory int
}

func (d debugInfo) String() string {
	return fmt.Sprintf(
		"(took: %d μs; aliases: %d; history: %d)",
		time.Since(d.start).Microseconds(),
		d.amountAliases,
		d.amountHistory,
	)
}

var debug debugInfo
