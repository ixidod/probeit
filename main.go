package main

import (
	"time"

	"github.com/ixidod/probeit/cmd"
)

func main() {
	cmd.Run([]string{
		"https://example.com",
		"https://google.com",
	}, 3*time.Second, 2*time.Second)
}
