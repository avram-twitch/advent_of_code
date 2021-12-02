package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)

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

func execute_instruction(instruction string, aim int, depth int, horizontal_position int) (int, int, int) {
    command, spaces := unpack_instruction(instruction)

    if command == "down" {
        return aim + spaces, depth, horizontal_position
    } else if command == "up" {
        return aim - spaces, depth, horizontal_position
    } else if command == "forward" {
        return aim, depth + (aim * spaces), horizontal_position + spaces
    }

    log.Fatalf("Invalid command: %s", command)
    return -1, -1, -1
}

func unpack_instruction(instruction string) (string, int) {
    split_instruction := strings.Split(instruction, " ")
    command := split_instruction[0]
    spaces, err := strconv.Atoi(split_instruction[1])
    check(err)
    return command, spaces
}

func main() {
    lines := read_file("input.txt")

    aim := 0
    depth := 0
    horizontal_position := 0

    for _, instruction := range lines {
        aim, depth, horizontal_position = execute_instruction(instruction, aim, depth, horizontal_position)
    }

    fmt.Printf("Horizontal position X Depth = %d", horizontal_position * depth)
}
