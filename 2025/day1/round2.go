package main

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
    for i := 0; i < move.Count; i++ {
        step := 1

        if move.Dir == 'L' {
            step = -1
        }

        lock.PinPosition += step

        if lock.PinPosition < 0 {
            lock.PinPosition += 100
        }
        if lock.PinPosition > 99 {
            lock.PinPosition -= 100
        }

        // increment the 0 counter after wrapping has happened
        if lock.PinPosition == 0 {
            lock.ZeroCount++
        }
    }
}

func main() {
    movements := readInput("input.txt")
    // movements := readInput("test.txt")

    lock := LockState{PinPosition: PIN_START, ZeroCount: 0}
    previous := lock.PinPosition

    fmt.Println("Begin")
    fmt.Printf("Move: n/a,\tPrev: %d,\tPos: %d,\tZeroCount: %d\n", previous, lock.PinPosition, lock.ZeroCount)
    for _, movement := range movements {
        previous = lock.PinPosition
        lock.makeMovement(movement)
        fmt.Printf("Move: %c%d,\tPrev: %d,\tPos: %d,\tZeroCount: %d\n", movement.Dir, movement.Count, previous, lock.PinPosition, lock.ZeroCount)
    }

    fmt.Printf("Zero count: %d\n", lock.ZeroCount)
}
