package main

import "time"
import "fmt"

func main() {
  fmt.Println("Timer 1 started")
  timer1 := time.NewTimer(2 * time.Second)

  <-timer1.C
  fmt.Println("Timer 1 expired")

  fmt.Println("Timer 2 started")
  timer2 := time.NewTimer(2 * time.Second)

  go func(){
    <-timer2.C
    fmt.Println("Timer 2 expired")
  }()
  fmt.Scanln()
}