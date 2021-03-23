package worker

import (
	"autoscalegoroutines/db"
	"fmt"
	"time"
)

const (
	OFFLINE int = 1
	ONLINE  int = 0
)

func Work(i int, ch chan int) {
	a := 0
	b := <-ch
	for {
		if OFFLINE == b {
			fmt.Println(i, "I am going offline...")
			break
		}
		a++
		r := db.Redis()
		r.Set(fmt.Sprintf("WORKER_%d", i), a, db.DAYS_30*time.Second)
		r.Close()
	}
}
