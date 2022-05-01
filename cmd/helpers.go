// helpers
package main

import (
	"log"
)

//this struct have 3 logger on 3 lvl
//such as INFO WARN ERROR
type Log3 struct {
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger
}
