package main

import (
	"flag"

	"github.com/golang/glog"
)

func main() {
	defer glog.Flush()
	flag.Parse()
	glog.Info("Prepare to repel boarders")
}
