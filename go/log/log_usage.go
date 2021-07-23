package main

import (
	"log"
	"os"
)

type Log interface {
	Print(v ...interface{})
}

var logs []Log

func init() {
	logLogger := log.New(os.Stderr, "package:log", log.LstdFlags|log.Ltime|log.Lmsgprefix)
	logs = append(logs, logLogger)

}

func main() {
	for _, log := range logs {
		log.Print("a", 1)
	}
}
