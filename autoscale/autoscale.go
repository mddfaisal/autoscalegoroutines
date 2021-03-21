package autoscale

import (
	"autoscalegoroutines/cpuusage"
	"autoscalegoroutines/ramusage"
	"autoscalegoroutines/worker"
	"fmt"
	"runtime"
)

const (
	MAX_CPU_USAGE int = 80
	MAX_RAM_USAGE int = 90
	MAX_WORKERS       = 1000
)

var chanspool []chan int

func AutoScale() {
	cpuusage := cpuusage.AvgCpuUsage()
	ramusage := ramusage.RamUsage()
	fmt.Println(cpuusage, ramusage)
	// if cpuusage > MAX_CPU_USAGE || ramusage > MAX_RAM_USAGE {
	// 	scaledown()
	// } else if cpuusage < MAX_CPU_USAGE || ramusage < MAX_RAM_USAGE {
	// 	scaleup()
	// }
}

func scaledown() {

}

func scaleup() {
	goroutines := runtime.NumGoroutine()
	if goroutines <= MAX_WORKERS {
		c := make(chan int)
		chanspool = append(chanspool, c)
		go worker.Work(len(chanspool), c)
	}
}
