// package main

import (
    "fmt"
    "os"
    "bufio"
    "log"
)

type Battery []int

func (b Battery) CalculateJoltage() int {
    // fmt.Printf("Calculating joltage for battery %d\n", b)

    tensDigit := 0
    joltage := 0

    for i, bank := range b {
        // Don't consider the last digit as the tens digit!
        if bank > tensDigit && i < len(b) - 1 {
            // fmt.Printf("\tTens set to %d\n", bank)
            tensDigit = bank
            continue
        }

        candidateJoltage := (tensDigit * 10) + bank
        // fmt.Printf("\tCandidate jolt: %d, prev jolt: %d...", candidateJoltage, joltage)
        if candidateJoltage > joltage {
            // fmt.Printf("replacing!\n")
            joltage = candidateJoltage
        }
        // fmt.Printf("\n")
    }

    return joltage
}

func readInput(fp string) []Battery {
    file, err := os.Open(fp)
    if err != nil {
        log.Fatalf("Error opening file: %v", err)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)

    batteries := []Battery{}

    for scanner.Scan() {
        line := scanner.Text()
        battery := stringToDigits(line)
        batteries = append(batteries, battery)
    }

    return batteries
}

func stringToDigits(s string) []int {
    digits := make(Battery, len(s))
    for i, char := range s {
        digits[i] = int(char - '0')
    }
    return digits
}

func main() {
    // fp := "test.txt"
    fp := "input.txt"
    batteries := readInput(fp)
    sum := 0

    for _, battery := range batteries {
        sum += battery.CalculateJoltage()
    }
    fmt.Printf("Total joltage: %d\n", sum)
}
