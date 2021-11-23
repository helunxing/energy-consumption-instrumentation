package main

import (
	"time"

	"github.com/helunxing/energy-consumption-instrumentation/go/method"
)

func doWork() {
	x := 1
	for i := 0; i < 100000000; i++ {
		x *= 11111
	}
	time.Sleep(time.Second)
}

func aain() {
	// p := doWork
	method.ExeAndPrintCPUusage(doWork)
}
