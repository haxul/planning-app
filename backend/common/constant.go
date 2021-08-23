package common

import (
	"log"
	"os"
)

var Logger = log.New(os.Stdout, "planning-app", log.LstdFlags)
var Port = 9090

const BACKLOG_STATE = "BACKLOG"
const IN_PROGRESS_STATE = "IN_PROGRESS"
const REJECTED_STATE = "REJECTED"
const DONE_STATE = "DONE"
