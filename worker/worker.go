package worker

import (
	"autoscalegoroutines/db"
	"fmt"
	"time"
)

const (
	OFFLINE int = 1
)

func Work(i int, ch chan int) {
	r := db.Redis()
	a := 0
	for {
		if OFFLINE == <-ch {
			fmt.Println(i, "I am going offline...")
			break
		}
		r.Set(fmt.Sprintf("WORKER_%d", i), a, db.DAYS_30*time.Second)
	}
	r.Close()
}
