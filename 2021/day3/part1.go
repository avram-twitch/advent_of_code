package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func main() {
    diagnostic_report := read_file("input.txt")

    one_bit_more_numerous_threshold := (len(diagnostic_report) / 2) + (len(diagnostic_report) % 2)
    bit_counter := make([]int, len(diagnostic_report[0]))

    for _, line := range diagnostic_report {
        bit_counter = count_binary_line(line, bit_counter)
    }

    gamma, epsilon := extract_gamma_and_epsilon(bit_counter, one_bit_more_numerous_threshold)
    fmt.Printf("gamma: %d, epsilon: %d\n", gamma, epsilon)
    fmt.Printf("Multiplied: %d", gamma * epsilon)

}

func count_binary_line(line string, bit_counter []int) ([]int) {
    for i, char := range line {
        bit_counter[i] += int(char - '0')
    }
    return bit_counter
}

func extract_gamma_and_epsilon(bit_counter []int, threshold int) (int, int) {
    gamma, epsilon := 0, 0

    for _, bit := range bit_counter {
        gamma = gamma << 1
        epsilon = epsilon << 1
        gamma += bit / threshold 
        epsilon += 1 - bit / threshold 
    }
    return gamma, epsilon
}

func check(err error) {
    if err != nil {
        log.Fatalf("readlines: %s", err)
    }
}

func read_file(fp string) ([]string) {
    f, err := os.Open(fp)
    check(err)
    defer f.Close()

    var lines []string
    scanner := bufio.NewScanner(f)

    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    check(scanner.Err())
    return lines
}
