package main

import (
	"flag"
	"os"
	"runtime"
	"runtime/pprof"
	"time"

	"github.com/prometheus/common/log"
)

func profile() {
	// Add profiler for tshoot OOM
	currentTime := time.Now().Local().Format("2006-01-02")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile + "_" + currentTime)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	if *memprofile != "" {
		f, err := os.Create(*memprofile + "_" + currentTime)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close()
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}
