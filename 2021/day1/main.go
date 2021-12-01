package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

func read_file(fp string) ([]int) {
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

    if scanner.Err() != nil {
        log.Fatalf("readlines: %s", err)
    }
    return lines
}

func increment_if_depth_increased(prev_depth int, curr_depth int) (int) {
    if prev_depth == 0 {
        return 0
    }

    if curr_depth > prev_depth {
        return 1
    } else {
        return 0
    }
}

func main() {
    lines := read_file("input.txt")

    prev_depth := 0
    increase_counter := 0

    for _, curr_depth := range lines {
        increase_counter += increment_if_depth_increased(prev_depth, curr_depth)
        prev_depth = curr_depth
    }

    fmt.Printf("Final count: %d", increase_counter)
}
