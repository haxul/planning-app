package common

import (
	"log"
	"os"
)

var Logger = log.New(os.Stdout, "planning-app", log.LstdFlags)
var Port = 9090

const IN_PROGRESS_STATE = "IN_PROGRESS"
const REJECTED_STATE = "REJECTED"
const DONE_STATE = "DONE"
const COURSE_STATE = "COURSE"
const PET_STATE = "PET"
const VIDEO_STATE = "VIDEO"
const BOOK_STATE = "BOOK"
