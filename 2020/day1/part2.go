package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

func read_file(fp string) ([]int, error) {
    f, err := os.Open(fp)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    var lines []int
    scanner := bufio.NewScanner(f)

    for scanner.Scan() {
        val, err := strconv.Atoi(scanner.Text())
        if err != nil {
            log.Fatal(err)
        }
        lines = append(lines, val)
    }

    return lines, scanner.Err()
}

func main() {
    const TARGET int = 2020
    lines, err := read_file("input.txt")
    if err != nil {
        log.Fatalf("readlines: %s", err)
    }

    for i, a := range lines {
        for j, b := range lines[i:] {
            for _, c := range lines[j:] {
                if a + b + c == TARGET {
                    fmt.Println(a, b, c)
                    fmt.Println(a * b * c)
                }
            }
        }
    }
}
