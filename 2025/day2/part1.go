// package main

import (
    "fmt"
    "log"
    "strings"
    "os"
    "bufio"
    "strconv"
)

type Range struct {
    Start int
    End int
}


func readInput(fp string) []Range {
    file, err := os.Open(fp)
    if err != nil {
        log.Fatalf("Error opening file: %v", err)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)

    ranges := []Range{}

    for scanner.Scan() {
        line := scanner.Text()

        for val := range strings.SplitSeq(line, ",") {
            dashSplit := strings.Split(val, "-")
            start, _ := strconv.Atoi(dashSplit[0])
            end, _ := strconv.Atoi(dashSplit[1])
            newRange := Range{Start: start, End: end}
            ranges = append(ranges, newRange)
        }
    }

    return ranges
}

func (r *Range) invalidSum() int {
    fmt.Printf("Calculating sum of range %d - %d\n", r.Start, r.End)
    sum := 0
    for i := r.Start; i <= r.End; i++ {
        if isNumInvalid(i) {
            sum += i
        }
    }

    return sum
}

func isNumInvalid(num int) bool {
    numString := strconv.Itoa(num)
    fmt.Printf("\tChecking %d (%s)...", num, numString)

    // Numbers with odd digits can't be invalid
    if len(numString) % 2 != 0 {
        fmt.Printf("valid!\n")
        return false
    }

    midpoint := len(numString) / 2

    leftPattern := numString[:midpoint]
    rightPattern := numString[midpoint:]

    if leftPattern == rightPattern {
        fmt.Printf("Invalid!\n")
        return true
    }
    fmt.Printf("valid!\n")
    return false
}

func main() {
    // fp := "test.txt"
     fp := "input.txt"
    ranges := readInput(fp)

    sum := 0
    for _, r := range ranges {
        sum += r.invalidSum()
    }
    fmt.Printf("Sum: %d\n", sum)
}
