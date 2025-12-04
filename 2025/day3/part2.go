package main

import (
    "fmt"
    "os"
    "bufio"
    "log"
    "math"
)

type Battery []int

const NUM_BANKS = 12

func (b Battery) CalculateJoltage() int {
    // fmt.Printf("Calculating for %d\n", b)

    var banks [NUM_BANKS]int
    batterySize := len(b)

    for i, bank := range b {
        // fmt.Printf("\tChecking %d\n", bank)
        remainingSpotsToCheck := batterySize - i
        var checkStart int

        // Don't try to swap out banks when I don't have enough remaining banks
        // to meet the NUM_BANKS requirement
        if NUM_BANKS <= remainingSpotsToCheck {
            checkStart = 0
        } else {
            checkStart = NUM_BANKS - remainingSpotsToCheck
        }

        replacedPrevious := false
        for j := checkStart; j < len(banks); j++ {
            existingBank := banks[j]
            if existingBank < bank && !replacedPrevious {
                banks[j] = bank
                replacedPrevious = true
                // fmt.Printf("\tReplaced bank: %d\n", banks)
                continue
            }

            // If we already replaced a previous bank, we need to zero out the rest
            if replacedPrevious {
                banks[j] = 0
            }
        }
    }

    joltage := 0
    // eg 432 => (4 * 10 ^ 2 ) + (3 * 10 ^ 1) + (2 * 10 ^ 0)
    for i, bank := range banks {
        exp := float64(NUM_BANKS - i - 1)
        magnitude := int(math.Pow(10, exp))
        joltage += bank * magnitude
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
    // fp := "tmp.txt"
    batteries := readInput(fp)
    sum := 0

    for _, battery := range batteries {
        sum += battery.CalculateJoltage()
    }
    fmt.Printf("Total joltage: %d\n", sum)
}
