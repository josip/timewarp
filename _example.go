package main

import (
  "fmt"
  "github.com/josip/timewarp"
  "time"
)

func main() {
  // convert to desired timezone before performing any transformations!
  now := timewarp.Warp(time.Now().UTC())
  fmt.Println("Hello there visitor! It's ", now)
  fmt.Println(fmt.Sprintf("(shh, it's acually %s)", now.Time))
  fmt.Println(timewarp.Yesterday())

  fmt.Println(now.FirstDayOfWeek(time.Monday))
  fmt.Println(now.Next(time.Tuesday).SubMinutes(2))
  fmt.Println(now.EndOfDay())
}
