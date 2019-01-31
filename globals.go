
package main

import (
  "sync"
)

// Debug Switch
// Always have a quick switch to display debug data.
// Setting this will show requests in the console as they are served.
var debug = false

// Current Directory
var currentDirectory = "UNKNOWN"

// Number of Returned Results
var resultCount = 20

// Number of Detected Text Files
var textFileCount = 0

// Number of Goroutines Spawned
var goroutineCount = 0

// Number of Total Words Counted
var totalWordCount = 0

// Variable to store the list of all text file paths
var debugPathList = ""

// Variable to help When debugging text file reading and parsing
var debugOutput = ""

// Variable to help When debugging map counts
var debugMap = make(map[string]int)

// Combined Map
var totalMap = make(map[string]int)

// Channel to gather maps from all file processing goroutines
var mapChannel = make(chan map[string]int, 1000)

// WaitGroup for goroutines
var wg sync.WaitGroup


// Golang Console Colors
// Example: fmt.Print( cRed + "HelloWorld" + cClr )
var cClr				= "\u001b[0m"

var cBold				= "\u001b[1m"

var cBlack			= "\u001b[30m"
var cRed				= "\u001b[31m"
var cGreen			= "\u001b[32m"
var cYellow			= "\u001b[33m"
var cBlue				= "\u001b[34m"
var cMagenta		= "\u001b[35m"
var cCyan				= "\u001b[36m"
var cWhite			= "\u001b[37m"

var cBlackBG		= "\u001b[40m"
var cRedBG			= "\u001b[41m"
var cGreenBG		= "\u001b[42m"
var cYellowBG		= "\u001b[43m"
var cBlueBG			= "\u001b[44m"
var cMagentaBG	= "\u001b[45m"
var cCyanBG			= "\u001b[46m"
var cWhiteBG		= "\u001b[47m"


// Output Simplification
var breakspace = "\n"
var breakline = breakspace + cBlue + "  ====================================================" + cClr + breakspace


// Console Splash
var appinfo = `
  ` + cBlue + `====================================================` + cBold + cCyan + `
   _____      _   _    _
  |  __ \    | | | |  | |
  | |__) |_ _| |_| |__| | __ _ _   _  __ _  ___ _ __
  |  ___/ _`+"`"+` | __|  __  |/ _`+"`"+` | | | |/ _`+"`"+` |/ _ \ '_ \
  | |  | (_| | |_| |  | | (_| | |_| | (_| |  __/ | | |
  |_|   \__,_|\__|_|  |_|\__,_|\__,_|\__, |\___|_| |_|
                                     __/ |
                                    |___/` + cClr + `
  ` + cCyan + `Interview: ` + cWhite + `GetWheelsApp` + cClr + `
  ` + cCyan + `https://github.com/` + cYellow + `pathaugen` + cCyan + `/interview-wheels` + cClr + `
  ` + cBlue + `====================================================` + cClr + `
`
