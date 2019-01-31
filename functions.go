
package main

import (
	"fmt"
  "io/ioutil"
  "log"
  "os"
  "os/exec"
  "path/filepath"
  "regexp"
  "sort"
  "strconv"
  "strings"
  //"sync"
  "time"
)


// Clear Screen
func clearScreen() {
  cmd := exec.Command("cmd", "/c", "cls")
  cmd.Stdout = os.Stdout
  cmd.Run()
}


// App Splash
func appSplash() {
  fmt.Print( appinfo )
}


// Debug Status
func debugStatus() {
  //fmt.Print( breakspace + cBold + cCyan + "  Debug Status:" + cClr )
  fmt.Print( "[" )
  if debug {
    fmt.Print( cBold + cGreen + "DEBUG ON" + cClr )
  } else {
    fmt.Print( cBold + cRed + "DEBUG OFF" + cClr )
  }
  fmt.Print( "]" )
}


// Debug Toggle
func debugToggle() {
  if debug {
    debug = false
  } else {
    debug = true
  }
}


// Set Current Directory
func setCurrentDirectory() {
  ex, err := os.Executable()
  if err != nil {
   panic(err)
  }
  exPath := filepath.Dir(ex)
  currentDirectory = exPath
}


// Count Words In File
// Reference: https://golang.org/pkg/os/#FileInfo
func wordCountFile(path string, info os.FileInfo, err error) error {
  if err != nil {
    log.Print(err)
    return nil
  }

  // We want all .txt files only
  extension := filepath.Ext( path )
  if ( extension == ".txt" ) {

    // Add to number of detected text files
    textFileCount += 1

    // Spawn a goroutine to parse this file for word counts
    //wg.Add(1)
    //go func() {
    //}()
    //wordCountFileGoroutine(&wg, path)
    go wordCountFileGoroutine( path )

    // DEBUG: Display the path -> store in debugPathList to pull and display at-will
    debugPathList += breakspace + "  " + cYellow + strconv.Itoa(textFileCount) + cClr + ": " + path
  }

  return nil
}


//func wordCountFileGoroutine( wg *sync.WaitGroup, path string ) {
func wordCountFileGoroutine( path string ) {
  goroutineCount += 1

  // Parse the text file and store a map of all words, adding to it for each instance

  // Start by reading the file
  b, err := ioutil.ReadFile( path )
  if err != nil {
    fmt.Print( breakspace + cRed + "ERROR: " + err.Error() + breakspace )
  }

  // Convert file bytes to string
  str := string(b)

  // Strip out all special characters and punctuation
  reg, err := regexp.Compile("[^a-zA-Z]+")
  if err != nil {
    log.Fatal(err)
  }
  str = reg.ReplaceAllString( str, " " )

  // Eliminate whitespace to help speed processing
  str = strings.Replace( str, "\n", " ", -1)
  str = strings.Replace( str, "     ", " ", -1)
  str = strings.Replace( str, "    ", " ", -1)
  str = strings.Replace( str, "   ", " ", -1)
  str = strings.Replace( str, "  ", " ", -1)

  // Ignore case sensitivity
  str = strings.ToLower( str )

  // Split the text file into words by separating by space
  strArr := strings.Split( str, " " )

  // DEBUG: Count of this array is the total number of words in the text file
  //fmt.Print( breakspace + "  Word Count: " + strconv.Itoa( len(strArr) ) )

  // Loop the array of all words in the text file
  tempMap := make(map[string]int)
  currentWord := ``
  for i := 0; i < len(strArr); i++ {
    currentWord = strArr[i]
    // First check if the word is in the map
    i, ok := tempMap[currentWord]
    if ok {
      // Add to count of word if it is already in the map
      tempMap[currentWord] = i + 1
    } else {
      // Add word to map if doesn't exist
      tempMap[currentWord] = 1
    }
  }

  // Load the map into the channel to combine from all goroutines later
  mapChannel <- tempMap

  // DEBUG: Peek into a random file output
  debugOutput = str

  // DEBUG: Peek into a random map
  debugMap = tempMap

  // Total words in the file is merely a length of the strArr
  // TODO: Here and aggregate data or do all at once at the end when the maps are combined (currently done at end when combined maps)
}


// Debug Statistics
func debugStats( currentDirectory string ) {
  if debug {
    fmt.Print( breakline + cBold + cCyan + "  Debug Statistics:" + cClr )

    // DEBUG: Walk the Directories and Output Filenames/Paths
    fmt.Print( breakspace + breakspace + cBold + cCyan + "  Raw List of All Text Files Detected: " + cClr )
    fmt.Print( debugPathList )

    fmt.Print( breakspace + cBold + cCyan + "  Total Text Files Scanned: " + cYellow + strconv.Itoa(textFileCount) + cClr )

    fmt.Print( breakspace )

    fmt.Print( breakspace + cBold + cCyan + "  Total Goroutines Spawned: " + cYellow + strconv.Itoa(goroutineCount) + cClr )

    // DEBUG: This section shows the output of a random text file that was parsed, prior to being split into an array - scan the output here for any issues
    fmt.Print( breakspace +
      breakspace + cBold + cCyan + "  DEBUG: Random text file contents after parsing (look over for flaws to address): " + cClr +
      breakline + cBold + cYellow +
      debugOutput +
      cClr + breakline )

    // DEBUG: map output by seeing the output of a random text file's map
    fmt.Print( breakspace +
      cBold + cCyan + "  DEBUG: Random map contents after parsing (look over for flaws to address): " + cClr +
      breakline + cBold + cYellow +
      "  lorem: " + strconv.Itoa(debugMap["lorem"]) +
      breakspace + "  in: " + strconv.Itoa(debugMap["in"]) +
      breakspace + "  a: " + strconv.Itoa(debugMap["a"]) +
      cClr + breakline )

    // DEBUG: Testing Channel Values
    //testChannelMap := <- mapChannel
    //fmt.Print( breakspace +
      //cBold + cCyan + "  DEBUG: Testing Channel Values: " + cClr +
      //breakline + cBold + cYellow +
      //"  lorem: " + strconv.Itoa(testChannelMap["lorem"]) +
      //breakspace + "  in: " + strconv.Itoa(testChannelMap["in"]) +
      //breakspace + "  a: " + strconv.Itoa(testChannelMap["a"]) +
      //cClr + breakline )

    //fmt.Print( breakline )
  }
}


// Word Count Chart
func wordCount( currentDirectory string ) {

  // Handle each text file
  err := filepath.Walk( currentDirectory, wordCountFile )
  if err != nil {
    log.Fatal(err)
  }

  // Display Debug Statistics
  debugStats( currentDirectory )

  // Combine all maps from goroutines via channel
  //done := make(chan bool)
  go combineChannelMaps()
  //wg.Wait()
  //fmt.Print( breakspace + "Waiting for combineChannelMaps() to finish" )
	//<- done
	//fmt.Print( breakspace + "combineChannelMaps() completed" )
  time.Sleep(time.Second)

  // Sort the map by highest int value first in descending order
  var ss []kv
  // Utilize the map combined from all maps via channel
  //testChannelMap := <- mapChannel
  for k, v := range totalMap { // debugMap / testChannelMap / totalMap
    ss = append(ss, kv{k, v})
  }
  sort.Slice(ss, func(i, j int) bool {
    return ss[i].Value > ss[j].Value
  })

  // Variable to ensure resultCount limit is followed
  currentResult := 1

  // Top (user defined amount) Most Used Words
  wordRankingList := ""
  for _, kv := range ss {
    // Getting a total of all words is just adding up all the kv.Value together
    totalWordCount += kv.Value

    // We're now sorted from most to least, so now honor the resultCount limit
    if ( currentResult < resultCount+1 ) {
      // Store the most frequent word list to utilize at-will later
      // TODO: count / total = the % a word was used out of the total number of words -> backlog wishlist feature
      wordRankingList += breakspace + "  " + strconv.Itoa(currentResult) + "     " + cBold + cYellow + strconv.Itoa(kv.Value) + cClr + "      " + "%" + "        " + kv.Key
      currentResult += 1
    }
  }

  // Display totals
  fmt.Print( breakspace + cBold + cCyan + "  Top " +
    cYellow + strconv.Itoa(resultCount) + cCyan + " Words in " +
    cYellow + strconv.Itoa(textFileCount) + cCyan + " .txt Files out of " +
    cYellow + strconv.Itoa(totalWordCount) + cCyan + " Words Total:" + cClr +
    breakspace + breakspace + cBold + cCyan + "  Rank  Count  % Total  Word" +
    breakspace + "  ----  -----  -------  ----" + cClr +
    wordRankingList )

  fmt.Print( breakspace + breakline )
}


// Combine all maps from goroutines via channel
// func combineChannelMaps( done chan bool ) {
func combineChannelMaps( ) {
  //go func() {
  for {
    m, more := <- mapChannel
    if more {
      //fmt.Println("received job", i)
      //fmt.Print( breakspace + "Received From Channel (mapChannel)" )
      //fmt.Print( "." )

      // Merge the map from the channel with the other map data
      currentWord := ``
      for k, v := range m {
        currentWord = k
        // First check if the word is in the map
        i, ok := totalMap[currentWord]
        if ok {
          // Add to count of word if it is already in the map
          totalMap[currentWord] = i + v
        } else {
          // Add word to map if doesn't exist
          totalMap[currentWord] = v
        }
      }

    } else {
      //fmt.Print( "Received All Channel Data" )
      //done <- true
      return
    }
  }
  //}()
}
