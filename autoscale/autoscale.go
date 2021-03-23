package autoscale

import (
	"autoscalegoroutines/db"
	"autoscalegoroutines/worker"
	"fmt"
	"log"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

const (
	MAX_CPU_USAGE float64 = 80
	MAX_RAM_USAGE float64 = 80
	MAX_WORKERS   int     = 1000
)

var chanspool []chan int

func AutoScale() {
	ramusage := rampercentage()
	avgcpu := avrcpu()
	fmt.Println(ramusage, avgcpu)
	if avgcpu > MAX_CPU_USAGE || ramusage > MAX_RAM_USAGE {
		scaleout()
	} else if avgcpu < MAX_CPU_USAGE || ramusage < MAX_RAM_USAGE {
		scalein()
	}
	setworkers()
}

func scalein() {
	if len(chanspool) > 0 {
		c := chanspool[len(chanspool)-1]
		c <- worker.OFFLINE
		close(c)
		chanspool = chanspool[:len(chanspool)-1]
		setworkers()
	}
}

func scaleout() {
	goroutines := len(chanspool)
	if goroutines < MAX_WORKERS {
		c := make(chan int, MAX_WORKERS)
		chanspool = append(chanspool, c)
		go worker.Work(len(chanspool), c)
		c <- worker.ONLINE
	}
	setworkers()
}

func checkErr(e error) {
	if e != nil {
		log.Println(e)
	}
}

func rampercentage() float64 {
	vm, err := mem.VirtualMemory()
	checkErr(err)
	return vm.UsedPercent
}

func avrcpu() float64 {
	var s float64
	cpus, err := cpu.Percent(0, true)
	checkErr(err)
	for _, cpu := range cpus {
		s = s + cpu
	}
	return s / float64(len(cpus))
}

func setworkers() {
	r := db.Redis()
	r.Set("TOT_WORKERS", len(chanspool), db.DAYS_30*time.Second)
	r.Close()
}
