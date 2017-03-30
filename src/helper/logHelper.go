package logHelper

import "log"

//LogError function logs error with custom message
func LogError(err error, msg string) {
	if err != nil {
		log.Fatalf(msg, err)
	}
}
