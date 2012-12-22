# timewarp #
_A time travelling library for Go._


## Using timewarp ##
`TimeWarp` is a thin <strike>warpper</strike>wrapper around Go's built-in [time](http://golang.org/pkg/time) library.


    package main
    import (
      "fmt"
      "github.com/josip/timewarp"
      "time"
    )

    func main() {
      now := timewarp.Warp(time.Now())
      fmt.Println("Hello there visitor! It's ", now)
      fmt.Println(fmt.Sprintf("(shh, it's acually %s)", now.Time))
    }
   
Please remember at all times that `Time` property returns the underlying `time.Time` struct.

### Basics ###

“Adding” or ”subtracting“ hours/minutes/seconds to/from the warped date can be performed with `Add` and `Sub` methods, respectively.

    now.Add(h, m, s, ms, µs, ns)
    now.Sub(h, m, s, ms, µs, ns)
    
    now.Add(1).Sub(0, 3).Sub(0, 7, 0, 10)

Or with the accompanying helpers:

    now.AddHours(1).SubMinutes(3).SubMinutes(7)
    now.AddSeconds(30).SubMinutes(1).AddDays(3)

Please note that helper methods are available for weeks, days, hours, minutes and seconds.

### Time travel ###

Besides `Add` and `Sub`, timewarp possesses a few more methods for smoother time travel within the day and the week. 

    fmt.Println("It's a new day…", now.BeginningOfDay())
    fmt.Println("every 24 hours…", now.EndOfDay())

Methods `BeginningOfDay` and `EndOfDay` modify hh:mm:ss to 00:00:00 and 23:59:59, respectively. 

    now.FirstDayOfWeek(time.Monday)

`FirstDayOfWeek` returns date (at midnight) of the first day in the week - `time.Monday` if you're in the Europe, `time.Sunday` if you're in the US or `time.Friday` if you're a certain internet celebrity.

`This` is a generalised version of the previously mentioned method, it returns the date of a day in the current week:

    now.This(time.Saturday)
    
    if now.Next(time.Tuesday).IsChristmas() {
      fmt.Println("Christmas is on Tuesday! Woohoo!")
	} else {
      fmt.Println("Not yet :(")
      if now.IsFriday() {
        fmt.Println("…but at least it's Friday!")
      }
	}

`Next` and `Last` do almost the same thing - except that they return the dates in the next or the previous week. 

### One more thing… ###

Timewarp packs a few more methods to stop you from repeating yourself (a common symptom of time travel):

    timewarp.Now()           // => timewarp.Wrap(time.Now())
    timewarp.Today()         // => Now().BeginningOfDay()
    timewarp.Tomorrow()      // => Today().Tomorrow()
    timewarp.Yesterday()     // => Today().Yesterday()
    smtmYesterday.Since(now) // => string(yesterday - now)

Both `Tomorrow` and `Yesterday` are also available as methods on `TimeWrap` structs.

## Licence ##
Timewarp is freely available under MMIT licence.