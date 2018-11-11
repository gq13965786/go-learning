package main

import "fmt"
import "os"

func main(){
  defer fmt.Println("!")//no printing

  os.Exit(3)
}
