package main

import (
    "bufio"
    "log"
    "fmt"
    "os"
    "strings"
)

func read_file(filePath string) []string {
    f, err := os.Open(filePath)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    var rucksacks []string

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        line := scanner.Text()

        rucksacks = append(rucksacks, line)
    }

    return rucksacks
}

func init() {
}

func getPriority(rucksacks []string) int {
    myChar := getMatchingChar(rucksacks)
    myInt := int(myChar[0])

    if myInt >= int('a') && myInt <= ('z') {
        return myInt - int('a') + 1
    }
    return myInt - int('A') + 27
}

func getMatchingChar(rucksacks []string) string {
    first := rucksacks[0]
    second := rucksacks[1]
    third := rucksacks[2]

    for i := 0; i < len(first); i++ {
        myChar := string(first[i])

        if strings.Contains(second, myChar) && strings.Contains(third, myChar) {
            return myChar
        }
    }

    return "a"
}

func main() {
    // rucksacks := read_file("top_ten.txt")
    rucksacks := read_file("input.txt")

    sum := 0
    for i := 0; i < len(rucksacks); i = i + 3 {
        sliceRucksacks := rucksacks[i:i + 3]
        sum += getPriority(sliceRucksacks)
    }

    fmt.Printf("Sum: %d", sum)
}
