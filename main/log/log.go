package log

import "log"

type Log struct {
	Logger *log.Logger
}

func (log *Log) info(s string) {
	log.info(s)
}
