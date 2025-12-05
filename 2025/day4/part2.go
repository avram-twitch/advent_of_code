package main

import (
    "os"
    "log"
    "bufio"
    "fmt"
)

const ROLL = "@"
const EMPTY = "."
const ADJACENT_ROLLS_LIMIT = 4

type Diagram [][]string
type RollLoc struct {
    X int
    Y int
}

func (d Diagram) Height() int {
    return len(d)
}

func (d Diagram) Width() int {
    if len(d) == 0 {
        return 0
    }
    return len(d[0])
}

func (d Diagram) IsInBounds(x int, y int) bool {
    xInBounds := x >= 0 && x < d.Width()
    yInBounds := y >= 0 && y < d.Height()
    return xInBounds && yInBounds
}

func (d Diagram) Print() {
    for _, row := range d {
        for _, ch := range row {
            fmt.Print(ch)
        }
        fmt.Println()
    }
}

func (d Diagram) RemoveRolls(rolls []RollLoc) {
    for _, rollLoc := range rolls {
        d[rollLoc.Y][rollLoc.X] = EMPTY
    }
}

func (d Diagram) CalcRollsAccessible() int {
    accessibleCount := 0

    for {
        d.Print()
        var locs []RollLoc
        initialCount := accessibleCount
        for y, row := range d {
            for x := range row {
                if d.IsRollAccessible(x, y) {
                    accessibleCount++
                    locs = append(locs, RollLoc{X: x, Y: y})
                }
            }
        }

        d.RemoveRolls(locs)
        // No change, nothing left to do
        if initialCount == accessibleCount {
            break
        }
    }

    return accessibleCount
}

func (d Diagram) IsRollAccessible(x int, y int) bool {
    // Don't count non rolls!
    if d[y][x] != ROLL {
        return false
    }

    adjacentRolesCount := 0
    for x_check := x - 1; x_check <= x + 1; x_check++ {
        for y_check := y - 1; y_check <= y + 1; y_check++ {

            // In bounds and not the given x/y coordinates themselves
            if d.IsInBounds(x_check, y_check) && !(x_check == x && y_check == y) {
                if d[y_check][x_check] == ROLL {
                    adjacentRolesCount++
                }
            }

            if adjacentRolesCount >= ADJACENT_ROLLS_LIMIT {
                return false
            }
        }
    }

    return true
}

// ------- File Parsing ---------

func readInput(fp string) Diagram {
    f, err := os.Open(fp)
    if err != nil {
        log.Fatalf("Error opening file: %v", err)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    var out Diagram

    for scanner.Scan() {
        line := scanner.Text()
        row := make([]string, len(line))

        for i, ch := range line {
            row[i] = string(ch)
        }

        out = append(out, row)
    }

    return out
}

func main() {
    fp := "test.txt"
    // fp := "input.txt"
    diagram := readInput(fp)
    fmt.Printf("Accessible Count: %d", diagram.CalcRollsAccessible())
}
