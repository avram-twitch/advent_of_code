package main

import (
    "fmt"
)

func main() {
    // fp := "test.txt"
    fp := "input.txt"
    input := readInput(fp)

    fmt.Printf("Part1 Fresh Count: %d\n", Part1(input))
    fmt.Printf("Part2 Fresh Count: %d\n", Part2(input))
}
