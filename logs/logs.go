package logs

import (
	"log"
	"os"
)

func CreateLogFile() {
	logfile, err := os.Create("app.log")
	if err != nil {
		log.Fatal("create log file:", err)
	}
	defer logfile.Close()
	log.SetOutput(logfile)
}
