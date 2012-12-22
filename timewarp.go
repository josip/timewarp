package timewarp

import (
  "fmt"
  "strings"
  "time"
)

type TimeWarp struct {
  Time time.Time
}

// Warp a time.Time struct into a TimeWarp.
func Warp(t time.Time) TimeWarp {
  return TimeWarp{t}
}

// Now warps current time.
func Now() TimeWarp {
  return Warp(time.Now())
}

// Today sets hh:mm to 00:00 for today's date.
func Today() TimeWarp {
  return Now().BeginningOfDay()
}

// Tomorrow returns today+1 at 00:00.
func Tomorrow() TimeWarp {
  return Today().Tomorrow()
}

// Yesterday returns today-1 at 00:00.
func Yesterday() TimeWarp {
  return Today().Yesterday()
}

// Add adds a duration relative to t. Arguments are defined as as follows:
// hours, minutes, seconds, milliseconds, microseconds, nanoseconds
// Use 0 to leave the property unmodifed.
func (t TimeWarp) Add(args ...int) TimeWarp {
  t.Time = add(t.Time, args)
  return t
}

func (t TimeWarp) AddWeeks(w int) TimeWarp {
  return t.Add(w * 7 * 24)
}

func (t TimeWarp) AddDays(d int) TimeWarp {
  return t.Add(d * 24)
}

func (t TimeWarp) AddHours(h int) TimeWarp {
  return t.Add(0, h)
}

func (t TimeWarp) AddMinutes(m int) TimeWarp {
  return t.Add(0, 0, m)
}

func (t TimeWarp) AddSeconds(s int) TimeWarp {
  return t.Add(0, 0, 0, s)
}

// Subtracts duration from t.
func (t TimeWarp) Sub(args ...int) TimeWarp {
  t.Time = sub(t.Time, args)
  return t
}

func (t TimeWarp) SubWeeks(w int) TimeWarp {
  return t.Sub(w * 7 * 24)
}

func (t TimeWarp) SubDays(d int) TimeWarp {
  return t.Sub(d * 24)
}

func (t TimeWarp) SubHours(h int) TimeWarp {
  return t.Sub(h)
}

func (t TimeWarp) SubMinutes(m int) TimeWarp {
  return t.Sub(0, m)
}

func (t TimeWarp) SubSeconds(s int) TimeWarp {
  return t.Sub(0, 0, s)
}

// Tomorrow returns the day after t.
func (t TimeWarp) Tomorrow() TimeWarp {
  return t.AddDays(1)
}

// Yesterday returns the day before t.
func (t TimeWarp) Yesterday() TimeWarp {
  return t.SubDays(1)
}

// Date of this weeks' specified day.
func (t TimeWarp) This(day time.Weekday) TimeWarp {
  var res TimeWarp = t
  var d int = int(res.Time.Weekday()) - int(day)
  if d == 0 {
    return res
  }

  if d > 0 {
    res = res.SubDays(d)
  } else {
    res = res.AddDays(-d)
  }

  return res
}

// Next occurance of the specified weekday.
func (t TimeWarp) Next(day time.Weekday) TimeWarp {
  return t.This(day).AddDays(7)
}

// Last occurance of the specified weekday.
func (t TimeWarp) Last(day time.Weekday) TimeWarp {
  return t.This(day).SubDays(7)
}

// BeginningOfDay sets hours:minutes to 00:00
func (t TimeWarp) BeginningOfDay() TimeWarp {
  d, _ := time.ParseDuration(toDurationString(t.Time, "hmsn"))
  t.Time = t.Time.Add(-d)
  return t
}

// EndOfDay sets hours:minutes to 23:59
func (t TimeWarp) EndOfDay() TimeWarp {
  h, m, s := t.Time.Clock()
  // TODO: set milliseconds to 0000
  return t.Add(23-h, 59-m, 59-s)
}

// FirstDayOfWeek returns date of the first day in the week
// First day (ie. time.Monday or time.Sunday) has to be manually specified.
func (t TimeWarp) FirstDayOfWeek(startsWith time.Weekday) TimeWarp {
  // TODO: determine startsWith through locale or something
  return t.This(startsWith).BeginningOfDay()
  //return t.SubDays(int(t.Time.Weekday()) - int(startsWith)).BeginningOfDay()
}

// Hours, seconds and minutes that have 
func (t TimeWarp) Since(now TimeWarp) string {
  var d time.Duration = now.Time.Sub(t.Time)
  return d.String()
}

// IsFriday returns true on Friday, false otherwise
func (t TimeWarp) IsFriday() bool {
  return t.Time.Weekday() == time.Friday
}

// IsChristmas returns true on Christmas day, false otherwise
func (t TimeWarp) IsChristmas() bool {
  _, m, d := t.BeginningOfDay().Time.Date()
  return m == time.December && d == 25
}

func add(t time.Time, args []int) time.Time {
  var suffixes []string = []string{"h", "m", "s", "ms", "us", "ns"}
  var dur string

  for i, arg := range args {
    if arg != 0 {
      dur += fmt.Sprintf("%d%s", arg, suffixes[i])
    }
  }

  d, err := time.ParseDuration(dur)
  if err != nil {
    panic(err)
  }

  return t.Add(d)
}

func sub(t time.Time, args []int) time.Time {
  for i, arg := range args {
    if arg != 0 {
      args[i] = -args[i]
      break
    }
  }

  return add(t, args)
}

func toDurationString(t time.Time, parts string) (d string) {
  switch {
  case strings.Contains(parts, "h"):
    d += fmt.Sprintf("%dh", t.Hour())
    fallthrough
  case strings.Contains(parts, "m"):
    d += fmt.Sprintf("%dm", t.Minute())
    fallthrough
  case strings.Contains(parts, "s"):
    d += fmt.Sprintf("%ds", t.Second())
    fallthrough
  case strings.Contains(parts, "n"):
    d += fmt.Sprintf("%dns", t.Nanosecond())
  }
  return
}
