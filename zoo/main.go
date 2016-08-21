package main

import (
  "fmt"
  "./animals"
  "math"
)

func main()  {
  fmt.Println(AppName())
  fmt.Println(animals.ElephantFeed())
  fmt.Println(animals.MonkeyFeed())
  fmt.Println(animals.RabbitFeed())
  fmt.Printf("uint32 max value = %d\n", math.MaxUint32)
}
