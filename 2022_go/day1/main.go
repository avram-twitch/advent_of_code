package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)

func read_file(fp string) ([][]int) {

    f, err := os.Open(fp)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    var elves [][]int
    var currentElf []int

    scanner := bufio.NewScanner(f)

    for scanner.Scan() {
        line := scanner.Text()
        if strings.TrimSpace(line) == "" {
            if len(currentElf) > 0 {
                elves = append(elves, currentElf)
                currentElf = nil
            }
        } else {
            num, err := strconv.Atoi(line)
            if err != nil {
                log.Fatal(err)
            }
            currentElf = append(currentElf, num)
        }
    }

    if len(currentElf) > 0 {
        elves = append(elves, currentElf)
    }

    if scanner.Err() != nil {
        log.Fatalf("readlines: %s", err)
    }

    return elves
}

func main() {
    elves := read_file("input.txt")

    maxCalories := 0

    for _, elf := range elves {
        currentCalories := 0
        for _, calorie := range elf {
            currentCalories += calorie
        }

        if currentCalories > maxCalories {
            maxCalories = currentCalories
        }
    }

    fmt.Printf("Max Calories is: %d", maxCalories)

}
