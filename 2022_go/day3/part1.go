package main

import (
    "bufio"
    "log"
    "fmt"
    "os"
    "strings"
)

func read_file(filePath string) [][]string {
    f, err := os.Open(filePath)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    var rucksacks [][]string
    var compartments []string

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        line := scanner.Text()
        length := len(line)

        midpoint := length / 2
        firstPart := line[0:midpoint]
        secondPart := line[midpoint:length]
        compartments = append(compartments, firstPart)
        compartments = append(compartments, secondPart)
        rucksacks = append(rucksacks, compartments)
        compartments = nil
    }

    return rucksacks
}

func init() {
}

func getPriority(rucksack []string) int {
    myChar := getMatchingChar(rucksack)
    myInt := int(myChar[0])

    if myInt >= int('a') && myInt <= ('z') {
        return myInt - int('a') + 1
    }
    return myInt - int('A') + 27
}

func getMatchingChar(rucksack []string) string {
    firstPart := rucksack[0]
    secondPart := rucksack[1]

    for i := 0; i < len(firstPart); i++ {
        myChar := string(firstPart[i])
        if strings.Contains(secondPart, myChar) {
            return myChar
        }
    }
    return "a"
}

func main() {
    // rucksacks := read_file("top_ten.txt")
    rucksacks := read_file("input.txt")

    sum := 0
    for _, rucksack := range rucksacks {
        sum += getPriority(rucksack)
        // fmt.Printf("First: %s Second: %s Matched char: %s Priority: %d\n", rucksack[0], rucksack[1], getMatchingChar(rucksack), getPriority(rucksack))
    }

    fmt.Printf("Sum: %d", sum)
}
