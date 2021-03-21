package main

import (
	"autoscalegoroutines/autoscale"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		autoscale.AutoScale()
	}
}
