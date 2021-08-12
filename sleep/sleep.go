package sleep

import (
	"log"
	"time"

	"github.com.sfragata/dontsleep/mouse"
)

type SleepInfo struct {
	TotalTimeRunning int
	SleepDelay       int
	Verbose          bool
}

func (si SleepInfo) Wake() {
	if si.Verbose {
		log.Printf("Running for %d minutes with %d minutes delay", si.TotalTimeRunning, si.SleepDelay)
	}

	timeRunning := 0
	for ok := true; ok; ok = timeRunning < si.TotalTimeRunning {
		moveMouse := mouse.Options{
			Verbose:    si.Verbose,
			MoveOffset: 30,
		}
		moveMouse.Move()
		if timeRunning+si.SleepDelay > si.TotalTimeRunning {
			si.delay(si.TotalTimeRunning - timeRunning)
		} else {
			si.delay(si.SleepDelay)
		}
		timeRunning += si.SleepDelay
	}
}

func (si SleepInfo) delay(delay int) {
	if si.Verbose {
		log.Printf("Sleeping %d minutes\n", delay)
	}
	time.Sleep(time.Duration(delay) * time.Minute)
}
