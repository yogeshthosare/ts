package strlog

import (
	log "github.com/sirupsen/logrus"
)

// Create a new instance of the logger. You can have any number of instances.
var CommonLogger = log.WithFields(log.Fields{
    "msg": "common logs here",
})

