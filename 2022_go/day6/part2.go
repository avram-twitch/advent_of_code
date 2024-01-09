package main

import (
    "os"
    "log"
)

func read_file(fp string) string {
    content, err := os.ReadFile(fp)

    if err != nil {
        log.Fatal(err)
    }

    return string(content)
}

func find_marker(stream string) int {
    const MARKER_LENGTH = 14
    for i := 0; i < len(stream); i++ {
        if i < MARKER_LENGTH {
            continue
        }

        if all_unique(string(stream[i - MARKER_LENGTH:i])) {
            return i
        }
    }

    return MARKER_LENGTH
}

func all_unique(marker string) bool {
    // log.Printf("Checking if %s is all unique", marker)
    for i := 0; i < len(marker) - 1; i++ {
        for j := i + 1; j < len(marker); j++ {
            if marker[i] == marker[j] {
                return false
            }
        }
    }
    return true
}

func main() {
    stream := read_file("input.txt")
    // stream := "mjqjpqmgbljsphdztnvjfqwrcgsmlb"
    markerStart := find_marker(stream)
    log.Printf("%d\n", markerStart)
}
