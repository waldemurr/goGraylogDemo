package talker

import (
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"
)

var stdFields = log.Fields{
	"service": "Talker",
	"file":    "historysync.go",
}
var val int

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	logger := log.WithFields(stdFields)
	
	for {
		val = rand.Int()
		switch val % 5 {
		case 0:
			logger.WithFields(log.Fields{
				"function": "main"}).Info("Setup redis orders...")
		case 1:
			logger.WithFields(log.Fields{
				"function": "main"}).Debug("Unblock all users done!")
		case 2:
			logger.WithFields(log.Fields{
				"function": "main"}).Warn("Block all users done!")
		case 3:
			logger.WithFields(log.Fields{
				"function": "main"}).Error("done")
		default:
			logger.WithFields(log.Fields{
				"function": "main"}).Trace("select me before i go")
			time.Sleep(time.Second)
		}
	}
}
