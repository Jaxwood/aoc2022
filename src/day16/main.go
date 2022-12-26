package main

import (
  "fmt"
  "strings"
  "strconv"
)

type ValveNetwork = []Valve

type Valve struct {
  name string
  rate int
  connections []string
}

func parse(input string) ValveNetwork {
  lines := strings.Split(input, "\n")
  result := ValveNetwork{}
  for _, line := range lines {
    if line == "" {
      continue
    }
    segments := strings.Split(line, " ")
    rateStr := strings.Split(segments[4], "=")
    rate, _ := strconv.Atoi(strings.Trim(rateStr[1], ";"))
    valves := []string{}
    for _, valve := range segments[9:] {
      valves = append(valves, strings.Trim(valve, ","))
    }
    result = append(result, Valve{segments[1], rate, valves})
  }
  return result
}

func toMap(valves ValveNetwork) map[string]Valve {
  result := map[string]Valve{}
  for _, valve := range valves {
    result[valve.name] = valve
  }
  return result
}

func day16(input string) int {
  valves := toMap(parse(input))
  totalMinutes := 30
  start := "AA"
  return 0
}
