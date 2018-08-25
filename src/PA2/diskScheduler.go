/*
  “I Alex Phillips (al562633) affirm that this program is entirely
  my own work and that I have neither developed my code
  together with any another person,  nor copied any code
  from anyother person, nor permitted my code to be copied
  or otherwise used by any other person, nor have I copied,
  modified, or otherwise used programs created by others.
  I acknowledge that any violation of the above terms will
  be treated as academic dishonesty.
*/

package main

import (
  "bufio"
  "os"
  "fmt"
  "strconv"
  "log"
  "strings"
  "sort"
)

type schedule struct {
  process string
  lower int
  upper int
  initial int
  requests []int
}

func parse() *schedule{

  file, _ := os.Open(os.Args[1])
  defer file.Close()
  scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

  // Data Struct
  scheduler := new(schedule)

  for scanner.Scan() {
    var word = (scanner.Text())
    array := strings.Split(word, " ")

    if array[0] == "use" {
      scheduler.process = array[1]
    }

    if array[0] == "lowerCYL" {
      temp, err := strconv.Atoi(array[1])
      if err != nil {
        fmt.Println("Error")
      }
      scheduler.lower = temp
    }

    if array[0] == "upperCYL" {
      temp, err := strconv.Atoi(array[1])
      if err != nil {
        fmt.Println("Error")
      }
      scheduler.upper = temp
    }

    if array[0] == "initCYL" {
      temp, err := strconv.Atoi(array[1])
      if err != nil {
        fmt.Println("Error")
      }
      scheduler.initial = temp
    }

    if array[0] == "cylreq" {
      temp, err := strconv.Atoi(array[1])
      if err != nil {
        fmt.Println("Error")
      }
      scheduler.requests = append(scheduler.requests, temp)
    }
  }

  return scheduler
}

// First Come First Serve
func fcfs(scheduler *schedule) {

  fmt.Println("Seek algorithm: FCFS")
  fmt.Printf("\tLower cylinder: %5d\n", scheduler.lower)
  fmt.Printf("\tUpper cylinder: %5d\n", scheduler.upper)
  fmt.Printf("\tInit cylinder: %5d\n", scheduler.initial)
  fmt.Printf("\tCylinder requests:\n")

  length := len(scheduler.requests)

  spot := scheduler.initial
  total := 0

  for i := 0; i < length; i++ {
    fmt.Printf("\t\tCylinder %5d\n", scheduler.requests[i])
  }

  for i := 0; i < length; i++ {
    fmt.Printf("Servicing %5d\n", scheduler.requests[i])

    total += absoluteDiff(spot, scheduler.requests[i])

    spot = scheduler.requests[i]
  }

  fmt.Printf("FCFS traversal count = %d\n", total)

  return
}

// Give the absolute value of the difference
// between two integers
func absoluteDiff(a int, b int) int{
  if a > b {
    return a - b
  }

  return b - a
}


func sstf(scheduler *schedule) {
  fmt.Println("Seek algorithm: SSTF")
  fmt.Printf("\tLower cylinder: %5d\n", scheduler.lower)
  fmt.Printf("\tUpper cylinder: %5d\n", scheduler.upper)
  fmt.Printf("\tInit cylinder: %5d\n", scheduler.initial)
  fmt.Printf("\tCylinder requests:\n")

  length := len(scheduler.requests)

  for i := 0; i < length; i++ {
    fmt.Printf("\t\tCylinder %5d\n", scheduler.requests[i])
  }

  if length == 0 {
    return
  }

  spot := scheduler.initial
  total := 0
  used := make([]bool, length)
  request := scheduler.requests[0]
  distance := absoluteDiff(spot, request)
  index := 0

  for i := 0; i < length; i++ {
    for j := 0; j < length; j++ {
      if ( !used[j] && (absoluteDiff(spot, scheduler.requests[j]) < distance) ) {
        request = scheduler.requests[j]
        distance = absoluteDiff(spot, scheduler.requests[j])
        index = j
      }

      if j == length-1 {
        used[index] = true
        fmt.Printf("Servicing %d\n", request)
        total += distance
        spot = request
      }
    }

    for j := 0; j < length; j++ {
      if !used[j] {
        request = scheduler.requests[j]
        distance = absoluteDiff(spot, scheduler.requests[j])
        index = j
        j = length
      }
    }
  }

  fmt.Printf("SSTF traversal count = %d\n", total)

  return
}

func scan(scheduler *schedule) {
  fmt.Println("Seek algorithm: SCAN")
  fmt.Printf("\tLower cylinder: %5d\n", scheduler.lower)
  fmt.Printf("\tUpper cylinder: %5d\n", scheduler.upper)
  fmt.Printf("\tInit cylinder: %5d\n", scheduler.initial)
  fmt.Printf("\tCylinder requests:\n")

  length := len(scheduler.requests)

  for i := 0; i < length; i++ {
    fmt.Printf("\t\tCylinder %5d\n", scheduler.requests[i])
  }

  if length == 0 {
    return
  }

  sort.Ints(scheduler.requests)

  // false is down and true is up
  target := scheduler.initial
  index := 0
  total := 0

  // Find where initial is between two values
  for i := 0; i < length - 1; i++ {
    if ( scheduler.requests[i] >= target) {
      index = i
      i = length
    }
  }

  spot := scheduler.initial

  for i := index; i < length; i++ {
    total += absoluteDiff(spot, scheduler.requests[i])
    spot = scheduler.requests[i]
    fmt.Printf("Servicing %d\n", scheduler.requests[i])
  }

  if index != 0 {
    total += absoluteDiff(spot, scheduler.upper)
    spot = scheduler.upper
  }

  for i := index-1; i >= 0; i-- {
    total += absoluteDiff(spot, scheduler.requests[i])
    spot = scheduler.requests[i]
    fmt.Printf("Servicing %d\n", scheduler.requests[i])
  }

  fmt.Printf("SCAN traversal count = %d\n", total)
}

func cscan(scheduler *schedule) {
  fmt.Println("Seek algorithm: C-SCAN")
  fmt.Printf("\tLower cylinder: %5d\n", scheduler.lower)
  fmt.Printf("\tUpper cylinder: %5d\n", scheduler.upper)
  fmt.Printf("\tInit cylinder: %5d\n", scheduler.initial)
  fmt.Printf("\tCylinder requests:\n")

  length := len(scheduler.requests)

  for i := 0; i < length; i++ {
    fmt.Printf("\t\tCylinder %5d\n", scheduler.requests[i])
  }

  if length == 0 {
    return
  }

  sort.Ints(scheduler.requests)

  // false is down and true is up
  target := scheduler.initial
  index := 0
  total := 0

  // Find where initial is between two values
  for i := 0; i < length - 1; i++ {
    if ( scheduler.requests[i] >= target) {
      index = i
      i = length
    }
  }

  spot := scheduler.initial

  for i := index; i < length; i++ {
    total += absoluteDiff(spot, scheduler.requests[i])
    spot = scheduler.requests[i]
    fmt.Printf("Servicing %d\n", scheduler.requests[i])
  }

  if index != 0 {

    total += absoluteDiff(spot, scheduler.upper)
    spot = scheduler.upper

    total += absoluteDiff(spot, scheduler.lower)
    spot = scheduler.lower

  }

  for i := 0; i < index; i++ {
    total += absoluteDiff(spot, scheduler.requests[i])
    spot = scheduler.requests[i]
    fmt.Printf("Servicing %d\n", scheduler.requests[i])
  }

  fmt.Printf("C-SCAN traversal count = %d\n", total)
}

func look(scheduler *schedule) {
  fmt.Println("Seek algorithm: LOOK")
  fmt.Printf("\tLower cylinder: %5d\n", scheduler.lower)
  fmt.Printf("\tUpper cylinder: %5d\n", scheduler.upper)
  fmt.Printf("\tInit cylinder: %5d\n", scheduler.initial)
  fmt.Printf("\tCylinder requests:\n")

  length := len(scheduler.requests)

  for i := 0; i < length; i++ {
    fmt.Printf("\t\tCylinder %5d\n", scheduler.requests[i])
  }

  if length == 0 {
    return
  }

  sort.Ints(scheduler.requests)

  // false is down and true is up
  target := scheduler.initial
  index := 0
  total := 0

  // Find where initial is between two values
  for i := 0; i < length - 1; i++ {
    if ( scheduler.requests[i] >= target) {
      index = i
      i = length
    }
  }

  spot := scheduler.initial

  for i := index; i < length; i++ {
    total += absoluteDiff(spot, scheduler.requests[i])
    spot = scheduler.requests[i]
    fmt.Printf("Servicing %d\n", scheduler.requests[i])
  }

  for i := index-1; i >= 0; i-- {
    total += absoluteDiff(spot, scheduler.requests[i])
    spot = scheduler.requests[i]
    fmt.Printf("Servicing %d\n", scheduler.requests[i])
  }

  fmt.Printf("LOCK traversal count = %d\n", total)
}

func clook(scheduler * schedule) {
  fmt.Println("Seek algorithm: C-LOOK")
  fmt.Printf("\tLower cylinder: %5d\n", scheduler.lower)
  fmt.Printf("\tUpper cylinder: %5d\n", scheduler.upper)
  fmt.Printf("\tInit cylinder: %5d\n", scheduler.initial)
  fmt.Printf("\tCylinder requests:\n")

  length := len(scheduler.requests)

  for i := 0; i < length; i++ {
    fmt.Printf("\t\tCylinder %5d\n", scheduler.requests[i])
  }

  if length == 0 {
    return
  }

  sort.Ints(scheduler.requests)

  // false is down and true is up
  target := scheduler.initial
  index := 0
  total := 0

  // Find where initial is between two values
  for i := 0; i < length - 1; i++ {
    if ( scheduler.requests[i] >= target) {
      index = i
      i = length
    }
  }

  spot := scheduler.initial

  for i := index; i < length; i++ {
    total += absoluteDiff(spot, scheduler.requests[i])
    spot = scheduler.requests[i]
    fmt.Printf("Servicing %d\n", scheduler.requests[i])
  }

  if index != 0 {
    total += absoluteDiff(spot, scheduler.requests[0])
    spot = scheduler.requests[0]
    fmt.Printf("Servicing %d\n", scheduler.requests[0])
  }

  for i := 1; i < index; i++ {
    total += absoluteDiff(spot, scheduler.requests[i])
    spot = scheduler.requests[i]
    fmt.Printf("Servicing %d\n", scheduler.requests[i])
  }

  fmt.Printf("C-LOCK traversal count = %d\n", total)
}

func main() {

  if len(os.Args) < 2 {
    log.Println("Missing a file in command")
    return
  }

  // if _, err := os.Stat(os.Args[2]); err == nil {
  //   os.Remove(os.Args[2])
  // }

  // out, err := os.OpenFile(os.Args[2], os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
  // if err != nil {
  //   panic(err)
  // }
  // defer out.Close()
  // log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
  // log.SetOutput(out)

  // parse the read file
  // that creates a new schedule
  scheduler := parse()

  // decide what type of scheduler and
  // and send it to the correct function
  if scheduler.process == "fcfs" {
    fcfs(scheduler)
  }

  if scheduler.process == "sstf" {
    sstf(scheduler)
  }

  if scheduler.process == "scan" {
    scan(scheduler)
  }

  if scheduler.process == "c-scan" {
    cscan(scheduler)
  }

  if scheduler.process == "look" {
    look(scheduler)
  }

  if scheduler.process == "c-look" {
    clook(scheduler)
  }
}
