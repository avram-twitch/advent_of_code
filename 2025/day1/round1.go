// package main

import (
    "fmt"
    "log"
    "bufio"
    "os"
    "strconv"
)

const PIN_START = 50

type Direction byte // 'L' or 'R'

type Movement struct {
    Dir Direction
    Count int
}

type LockState struct {
    PinPosition int
    ZeroCount int
}

func readInput(fp string) []Movement {
    file, err := os.Open(fp)
    if err != nil {
        log.Fatalf("Error opening file: %v", err)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)

    movements := []Movement{}

    for scanner.Scan() {
        line := scanner.Text()
        direction := Direction(line[0])
        count, err := strconv.Atoi(line[1:])

        if err != nil {
            fmt.Println("Error converting string to int:", err)
            return movements
        }

        newMovement := Movement{Dir: direction, Count: count}
        movements = append(movements, newMovement)

    }

    if err := scanner.Err(); err != nil {
        log.Fatalf("Error during scanning: %v", err)
    }

    return movements
}

func (lock *LockState) makeMovement(move Movement) {
    switch move.Dir {
    case 'L':
        lock.PinPosition -= move.Count
    case 'R':
        lock.PinPosition += move.Count
    default:
        panic("unknown direction")

    }

    // Enforce range of 0-99
    lock.PinPosition %= 100
    if lock.PinPosition < 0 {
        lock.PinPosition += 100
    }

    if lock.PinPosition == 0 {
        lock.ZeroCount++
    }
}

func main() {
    movements := readInput("input.txt")

    lock := LockState{PinPosition: PIN_START, ZeroCount: 0}

    for _, movement := range movements {
        fmt.Printf("Direction: %c, Count: %d\n", movement.Dir, movement.Count)
        lock.makeMovement(movement)
        // makeMovement(&lock, movement)
    }

    fmt.Printf("Zero count: %d\n", lock.ZeroCount)
}
