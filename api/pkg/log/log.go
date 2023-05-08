package log

import (
	"log"
	"os"
)

// Info writes logs in the color blue with "[INFO]" as prefix
var Info = log.New(os.Stdout, "\u001b[34m[INFO] \u001B[0m", log.LstdFlags|log.Lshortfile).Printf

// Warning writes logs in the color yellow with "[WARNING]" as prefix
var Warning = log.New(os.Stdout, "\u001b[33m[WARNING] \u001B[0m", log.LstdFlags|log.Lshortfile).Printf

// Error writes logs in the color red with "[ERROR]" as prefix
var Error = log.New(os.Stderr, "\u001b[31m[ERROR] \u001b[0m", log.LstdFlags|log.Lshortfile).Printf

// Debug writes logs in the color cyan with "[DEBUG]" as prefix
var Debug = log.New(os.Stdout, "\u001b[36m[DEBUG] \u001B[0m", log.LstdFlags|log.Lshortfile).Printf

var Fatal = log.New(os.Stderr, "\u001b[31m[FATAL] \u001b[0m", log.LstdFlags|log.Lshortfile).Fatalf

var Panic = log.New(os.Stderr, "\u001b[31m[PANIC] \u001b[0m", log.LstdFlags|log.Lshortfile).Panicf