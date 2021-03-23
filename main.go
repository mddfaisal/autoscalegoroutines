package main

import (
	"autoscalegoroutines/autoscale"
	"autoscalegoroutines/db"
	"time"
)

func main() {
	r := db.Redis()
	r.Set("TOT_WORKERS", 0, db.DAYS_30*time.Second)
	r.Close()
	ticker := time.NewTicker(100 * time.Millisecond)
	for range ticker.C {
		autoscale.AutoScale()
	}
}
