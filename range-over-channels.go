package main

import "fmt"

func main() {
  queue := make(chan string,2)
  queue <- "one"
  queue <- "two"
  close(queue)//interesting defer close(queue) deadlock

  for elem:=range queue{
    fmt.Println(elem)
  }

}
