package main

import (
    "os"
    "log"
    "bufio"
    "strings"
    "strconv"
)

func readInput(fp string) Input {
    f, err := os.Open(fp)
    if err != nil {
        log.Fatalf("Error opening file: %v", err)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)

    var ranges []Range
    // Get ranges
    for scanner.Scan() {
        line := scanner.Text()
        // Reached end of ranges
        if line == "" {
            break
        }
        dashSplit := strings.Split(line, "-")
        start, _ := strconv.Atoi(dashSplit[0])
        end, _ := strconv.Atoi(dashSplit[1])
        newRange := Range{Start: start, End: end}
        ranges = append(ranges, newRange)
    }

    var ids []int
    // Get ids
    for scanner.Scan() {
        line := scanner.Text()
        num, _ := strconv.Atoi(line)
        ids = append(ids, num)
    }

    return Input{Ranges: ranges, IDs: ids}
}
