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
  "fmt"
  "bufio"
  "os"
  "strconv"
  "log"
)

type single struct {
  name string
  arrival int
  burst int
}

type schedule struct {
  process int
  time int
  use string
  quantum int
  processes []single
}

func parse() *schedule {

  // FILE STRUCTURE: Reading In
  /*
    1) String 'Processcount' -> Declares what we get (Not held)
    2) Integer '#' -> Number of Processes (Held)
    3) String 'runfor' -> Declares what we get (Not held)
    4) Integer '#' -> Number of units in "time" (Held)
    5) String 'use' -> What scheduler? (Not held)
    6) String 'fcfs/sjf/rr' -> Type of scheduler to run (Held/ Can be a number)
        6.5)"rr" supplies quantum in next line
    7) Cycle "process", "name", "P#" (the process),
        "arrival", "#" (unit of time), "burst", "#" (units of time to complete)
    8) "end" (done with inputs)
  */

  file, _ := os.Open(os.Args[1])
  defer file.Close()
  scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

  // Data Struct
  scheduler := new(schedule)
  scheduler.processes = []single{}

  // Variables as flags
  counter := false
  time := false
  use := false
  quantum := false
  process := false
  name := false
  arrival := false
  burst := false
  quantumCount := 0
  tempName := ""
  tempArrival := ""

  for scanner.Scan() {
      var word = scanner.Text()

      if counter {
        temp, err := strconv.Atoi(word)
        if err != nil {
          fmt.Println("Error")
        }
        scheduler.process = temp
        counter = false
      }

      if time {
        temp, err := strconv.Atoi(word)
        if err != nil {
          fmt.Println("Error")
        }
        scheduler.time = temp
        time = false
      }

      if use {
        scheduler.use = word
        use = false
      }

      if (quantum && quantumCount == 0) {
        temp, err := strconv.Atoi(word)
        if err != nil {
          fmt.Println("Error")
        }
        scheduler.quantum = temp
        quantum = false
        quantumCount++
      }

      if (process && name) {
        tempName = word
        name = false
      }

      if (process && arrival) {
        tempArrival = word
        arrival = false
      }

      if(process && burst) {
        tempBurst, err := strconv.Atoi(word)
        if err != nil {
          fmt.Println("Error")
        }

        myInt, err := strconv.Atoi(tempArrival)
        if err != nil {
          fmt.Println("Error")
        }

        scheduler.processes = append(scheduler.processes,
              single{name: tempName, arrival: myInt, burst: tempBurst,})
        burst = false
      }


      if word == "processcount" {
        counter = true
      }

      if word == "runfor" {
        time = true
      }

      if word == "use" {
        use = true
      }

      if word == "quantum" {
        quantum = true
      }

      if word == "name" {
        process = true
        name = true
      }

      if word == "arrival" {
        arrival = true
      }

      if word == "burst" {
        burst = true
      }
  }
  // return Struct
  return scheduler
}

// Shortest Job First
func sjf(scheduler *schedule) {
  log.Printf("%3d processes\n", scheduler.process)
  log.Printf("Using preemptive Shortest Job First\n")

  time := scheduler.time
  processes := scheduler.process

  currentProcess := 0

  ready := []single{}
  temp := single{}
  tempIndex := 0
  selected := false
  lastTemp := single{}
  wait := make([]int, processes)

  for i := 0; i < time; i++ {

    // Check for arrivals
    for j := 0; j < processes; j++ {
      if scheduler.processes[j].arrival == i {
        log.Printf("Time %3d : " + scheduler.processes[j].name + " arrived\n", i)
        ready = append(ready, scheduler.processes[j])
        currentProcess++
      }
    }

    // Get rid of the bursts that have 0
    for w := 0; w < currentProcess; w++ {
      if(ready[w].burst == 0) {
        currentProcess = currentProcess - 1
        log.Printf("Time %3d : " + ready[w].name + " finished\n", i)
        slice := ready
        ready =  append(slice[:w], slice[w+1:]...)
        if len(ready) > 0 {
          temp = ready[0]
          tempIndex = 0
          selected = true
        } else {
          selected = false
        }
      }
    }

    if currentProcess == 0 {
      selected = false
    }

    //  Can Select and will be selected
    if currentProcess == 1 {
      temp = ready[0]
      tempIndex = 0
      if temp != lastTemp {
        log.Printf("Time %3d : " + temp.name + " selected (burst %3d)\n", i, temp.burst)
      }
      lastTemp = ready[0]
      selected = true
    }

    // Who gets to go
    for k := 0; k < currentProcess-1; k++ {
      if ready[k].burst != 0 {
        if ready[k].burst > ready[k+1].burst {
          temp = ready[k+1]
          tempIndex = k+1
        }
        selected = true
      }

      if k == currentProcess-2 {
        for loop := 0; loop < len(ready); loop++ {
          for inside := 0; inside < processes; inside++ {
            if (ready[loop].name == scheduler.processes[inside].name && loop != tempIndex) {
              wait[inside]++
            }
          }
        }
        if temp != lastTemp {
          log.Printf("Time %3d : " + temp.name + " selected (burst %3d)\n", i, temp.burst)
          lastTemp = ready[tempIndex]
        }
      }
    }

    // Use it
    if selected {
      ready[tempIndex] = single{name: temp.name, arrival: temp.arrival, burst: temp.burst-1}
      temp = ready[tempIndex]
      lastTemp = ready[tempIndex]
    } else {
      log.Printf("Time %3d : Idle\n", i)
    }
  }

  if len(ready) > 0 {
    log.Printf("Did not Finish\n")
  } else {
    log.Printf("Finished at time %3d\n\n", time)
  }

  for i := 0; i < processes; i++ {
    log.Printf(scheduler.processes[i].name + " wait %3d turnaround %3d\n", wait[i], scheduler.processes[i].burst + wait[i])
  }
}

// First Come First Serve
func fcfs(scheduler *schedule) {
  log.Printf("%3d processes\n", scheduler.process)
  log.Printf("Using First-Come First-Served\n")

  time := scheduler.time
  processes := scheduler.process
  selected := false
  temp := single{}

  ready := []single{}
  wait := make([]int, processes)

  for i := 0; i < time; i++ {
    // Check for arrivals
    for j := 0; j < processes; j++ {
      if scheduler.processes[j].arrival == i {
        log.Printf("Time %3d : " + scheduler.processes[j].name + " arrived\n", i)
        ready = append(ready, scheduler.processes[j])
      }
    }

    // Get rid of the burst that have 0
    if len(ready) >= 1 {
      if ready[0].burst == 0 {
        log.Printf("Time %3d : " + ready[0].name + " finished\n", i)
        slice := ready
        ready =  append(slice[:0], slice[1:]...)
        selected = false
      }
    }

    // Select it
    if (len(ready) >= 1 && !selected){
      log.Printf("Time %3d : " + ready[0].name + " selected (burst %3d)\n", i, ready[0].burst)

      selected = true
    }

    if selected {
      temp = ready[0]
      ready[0] = single{name: temp.name, arrival: temp.arrival, burst: temp.burst-1}

      for loop := 1; loop < len(ready); loop++ {
        for inside := 0; inside < processes; inside++ {
          if (ready[loop].name == scheduler.processes[inside].name) {
            wait[inside]++
          }
        }
      }
    } else {
      log.Printf("Time %3d : Idle\n", i)
    }
  }

  if len(ready) > 0 {
    log.Printf("Did not Finish\n")
  } else {
    log.Printf("Finished at time %3d\n\n", time)
  }

  for i := 0; i < processes; i++ {
    log.Printf(scheduler.processes[i].name + " wait %3d turnaround %3d\n", wait[i], scheduler.processes[i].burst + wait[i])
  }

}

// Round Robin
func rr(scheduler *schedule) {
  log.Printf("%3d processes\n", scheduler.process)
  log.Printf("Using Round-Robin\n")
  log.Printf("Quantum %3d\n\n", scheduler.quantum)

  time := scheduler.time
  processes := scheduler.process
  quantum := scheduler.quantum

  selected := false
  count := 0
  ready := []single{}
  temp := single{}
  index := 0
  length := 0
  rotate := 0
  wait := make([]int, processes)

  for i := 0; i < time; i++ {
    // Check for arrivals
    for j := 0; j < processes; j++ {
      if scheduler.processes[j].arrival == i {
        log.Printf("Time %3d : " + scheduler.processes[j].name + " arrived\n", i)
        ready = append(ready, scheduler.processes[j])
        length++
      }
    }

    // Get rid of the bursts that have 0
    for w := 0; w < length; w++ {
      if(ready[w].burst == 0) {
        length = length - 1
        count = 0

        log.Printf("Time %3d : " + ready[w].name + " finished\n", i)
        slice := ready
        ready =  append(slice[:w], slice[w+1:]...)

        if length != 0 {
          index = rotate % length
          temp = ready[index]
        }

        selected = false
      }
    }

    if len(ready) > 0 {
      if !selected && count == 0 {
        for loop := 0; loop < len(ready); loop++ {
          for inside := 0; inside < processes; inside++ {
            if (ready[loop].name == scheduler.processes[inside].name && loop != index) {
              wait[inside]++
            }
          }
        }

        count += 1
        log.Printf("Time %3d : " + ready[index].name + " selected (burst %3d)\n", i, ready[index].burst)
        temp = ready[index]
        ready[index] = single{name: temp.name, arrival: temp.arrival, burst: temp.burst-1}
        selected = true
      } else {
        for loop := 0; loop < len(ready); loop++ {
          for inside := 0; inside < processes; inside++ {
            if (ready[loop].name == scheduler.processes[inside].name && loop != index) {
              wait[inside]++
            }
          }
        }
        if count < quantum {
          count +=1
          temp = ready[index]
          ready[index] = single{name: temp.name, arrival: temp.arrival, burst: temp.burst-1}
        }
        if count == quantum {

          if len(ready) > 1 {
            rotate++
            index = rotate % length
            temp = ready[index]
          }

          count = 0
          selected = false
        }
      }
    } else {
      log.Printf("Time %3d : Idle\n", i)
    }
  }

  if len(ready) > 0 {
    log.Printf("Did not Finish\n")
  } else {
    log.Printf("Finished at time %3d\n\n", time)
  }

  for i := 0; i < processes; i++ {
    log.Printf(scheduler.processes[i].name + " wait %3d turnaround %3d\n", wait[i], scheduler.processes[i].burst + wait[i])
  }
}

func main() {

  if len(os.Args) < 3 {
    log.Println("Missing a file in command")
    return
  }

  if _, err := os.Stat(os.Args[2]); err == nil {
    os.Remove(os.Args[2])
  }

  out, err := os.OpenFile(os.Args[2], os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
  if err != nil {
    panic(err)
  }
  defer out.Close()
  log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
  log.SetOutput(out)

  // parse the read file
  // that creates a new schedule
  scheduler := parse()

  // decide what type of scheduler and
  // and send it to the correct function
  if scheduler.use == "rr" {
    rr(scheduler)
  }

  if scheduler.use == "sjf" {
    sjf(scheduler)
  }

  if scheduler.use == "fcfs" {
    fcfs(scheduler)
  }
}
