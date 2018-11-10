package main

import (
  "fmt"
  "time"
)

func main(){

  timer1 := time.NewTimer(2*time.Second)

  <-timer1.C
  fmt.Println("Timer 1 expired")

  timer2 := time.NewTimer(time.Second)
  go func(){
    <-timer2.C
    fmt.Println("Timer 2 expired too")
  }()
  //time.Sleep(5*time.Second)
  stop2 := timer2.Stop()
  if stop2 {//bool
    fmt.Println("Timer 2 stopped")
  }
}
