
package main

import (
  "bufio"
	"fmt"
  "os"
)


func main() {
  // Constantly monitor and combine maps
  //go combineChannelMaps()

  // Set Current Directory
  setCurrentDirectory()

  // Reading Input
  scanner := bufio.NewScanner(os.Stdin)
  var inputText string

  // Listen for input and break the loop if inputText == "q"
  for ( inputText != "q" ) {

    // Clear Screen
    clearScreen()

    // Display Application Splash
    appSplash()

    // User Selection Options Follow:
    if ( inputText == "d" ) {
      // Toggle Debug Mode
      debugToggle()
    } else if ( inputText == "." ) {
      // Utilize current directory for Word Count
      textFileCount = 0
      goroutineCount = 0
      totalWordCount = 0
      debugPathList = ""

      // Use current working directory
      fmt.Print( breakspace + cBold + cCyan + "  Performing Word Count on Current Directory:" + cClr +
        breakspace + cBold + cYellow + "  " + currentDirectory + cClr + breakspace )

      // Word Count Chart
      wordCount( inputText )
    } else if ( inputText != "" ) {
      // User provided string path for Word Count
      textFileCount = 0
      goroutineCount = 0
      totalWordCount = 0
      debugPathList = ""

      // Perform check if user provided string provided matches regex: [a-z]:\[.*]
      // TODO - Bonus Feature Wishlist

      // Use provided directory
      fmt.Print( breakspace + cBold + cCyan + "  Performing Word Count on Provided Directory:" + cClr +
        breakspace + cBold + cYellow + "  " + inputText + cClr + breakspace )

      // Word Count Chart
      wordCount( inputText )
    }

    // User's command prompt
    fmt.Print( breakspace + cBold + cCyan + "  Specify a directory to perform Word Count" + cClr +
      breakspace + "  [" + cYellow + "." + cClr + "] Trigger Current Directory [" + cBold + cYellow + currentDirectory + cClr + "]" +
      breakspace + "  [" + cYellow + "d" + cClr + "] Toggle Debug              " + cClr )
    debugStatus()
    fmt.Print( breakspace +
      "  [" + cYellow + "q" + cClr + "] Quit Application" + cClr +
      breakspace + cYellow + "  > " + cClr )

    // Await user input
    scanner.Scan()
    inputText = scanner.Text()

  }

  // Exit message (cleanup could go here)
  fmt.Print( breakline + cBold + cGreen + "  Application Exited Without Errors" + cClr + breakspace )
}
