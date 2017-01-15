package base

import log "github.com/Sirupsen/logrus"

func ErrorCheck(e error) {
	if e != nil {
		log.Error(e)
	}
}