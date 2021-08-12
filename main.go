package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com.sfragata/dontsleep/sleep"
	"github.com/integrii/flaggy"
)

// These variables will be replaced by real values when do gorelease
var (
	version = "none"
	date    string
	commit  string
)

func main() {

	info := fmt.Sprintf(
		"%s\nDate: %s\nCommit: %s\nOS: %s\nArch: %s",
		version,
		date,
		commit,
		runtime.GOOS,
		runtime.GOARCH,
	)

	flaggy.SetName("Don't Sleep")
	flaggy.SetDescription("Command line app to avoid your computer to sleep!!")
	flaggy.SetVersion(info)

	var timeToRunMinutes = 10
	flaggy.Int(&timeToRunMinutes, "t", "time", "Time to run (minutes)")

	var sleepDelayMinutes = 5
	flaggy.Int(&sleepDelayMinutes, "d", "delay", "Time to sleep before wake again (minutes)")

	var verbose = false
	flaggy.Bool(&verbose, "v", "verbose", "verbose mode")

	flaggy.Parse()

	sleepInfo := sleep.SleepInfo{
		TotalTimeRunning: timeToRunMinutes,
		SleepDelay:       sleepDelayMinutes,
		Verbose:          verbose,
	}

	sleepInfo.Wake()
	log.Print("Finished!")
}
