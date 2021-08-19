package common

import (
	"log"
	"os"
)

var Logger = log.New(os.Stdout, "planning-app", log.LstdFlags)
var Port = 9090
