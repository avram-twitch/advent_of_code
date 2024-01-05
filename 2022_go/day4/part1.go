package main

import (
    "fmt"
    "os"
    "log"
    "bufio"
    "strings"
    "strconv"
)

type Elf struct {
    start int
    finish int
}

func read_file(filePath string) [][]Elf {
    file, err := os.Open(filePath)

    if err != nil {
        log.Fatal(err)
    }

    var pairs [][]Elf
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
        splitLine := strings.Split(line, ",")

        var elves []Elf

        for _, line := range splitLine {
            thisRange := strings.Split(line, "-")
            first, ok := strconv.Atoi(thisRange[0])
            if ok != nil {
                log.Fatal(ok)
            }

            second, ok := strconv.Atoi(thisRange[1])

            if ok != nil {
                log.Fatal(ok)
            }

            elves = append(elves, Elf{start: first, finish: second})
        }

        pairs = append(pairs, elves)
        elves = nil
    }

    return pairs
}

func (e Elf) Contains(otherElf Elf) bool {
    return e.start <= otherElf.start && e.finish >= otherElf.finish
}

func main() {
    elfPairs := read_file("input.txt")

    sum := 0

    for _, pair := range elfPairs {
        containsOther := pair[0].Contains(pair[1]) || pair[1].Contains(pair[0])
        if containsOther {
            sum++
        }
    }

    fmt.Printf("Result: %d", sum)
}
