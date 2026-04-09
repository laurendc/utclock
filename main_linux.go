//go:build linux

package main

import (
	"fmt"
	"time"
)

func main() {
	utc := time.Now().UTC()
	fmt.Printf("<span foreground=\"#88C0D0\">%s</span>\n", utc.Format("15:04 UTC"))
}
