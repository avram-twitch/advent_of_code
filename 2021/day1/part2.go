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

func shift_window_to_include_new_value(new_value int, window []int) ([]int) {
    window = window[1:]
    window = append(window, new_value)
    return window
}

func depth_of_window(window []int) (int) {
    sum_of_depths := 0

    for _, depth := range window {

        // We don't have a complete window
        if depth == 0 {
            return 0
        }

        sum_of_depths += depth
    }
    return sum_of_depths

}

func main() {
    const WINDOW_SIZE int = 3

    lines := read_file("input.txt")

    var window = []int{0, 0, 0}

    prev_depth := 0
    increase_counter := 0

    for _, curr_depth := range lines {
        window = shift_window_to_include_new_value(curr_depth, window)
        curr_depth = depth_of_window(window)
        increase_counter += increment_if_depth_increased(prev_depth, curr_depth)

        prev_depth = curr_depth
    }

    fmt.Printf("Final count: %d", increase_counter)
}
