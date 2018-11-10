package main

import (
  "fmt"
)

func main() {

  messages := make(chan string,2)

  messages <-"buffered"
  messages <-"channel" //order like queue first in first out

  fmt.Println(<-messages)
  fmt.Println(<-messages)

}
