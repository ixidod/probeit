package main

import (
	"time"

	"github.com/ixidod/probeit/cmd"
)

func main() {
	cmd.Run("https://example.com", 3*time.Second, 2*time.Second)
}
