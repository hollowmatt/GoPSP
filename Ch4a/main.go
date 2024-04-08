package main

import (
	"os"
	"time"

	"hollowmatt/new_logger/pocketlog"
)

func main() {
	lgr := pocketlog.New(pocketlog.LevelInfo, pocketlog.WithOutput(os.Stdout))

	lgr.Infof("for your information, this is information")
	lgr.Errorf("error in the system, please correct")
	lgr.Debugf("beware the fly in the ointment")
	lgr.Infof("Hey... %d %v", 2022, time.Now())
}
