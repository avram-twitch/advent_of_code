package main

import (
    "os"
    "log"
    "bufio"
    "strings"
    "strconv"
    "fmt"
)

// Representing Trees as an array of arrays of ints
func loadTrees(fp string) [][]int {
    file, err := os.Open(fp)


    if err != nil {
        log.Fatal("Could not open file")
    }

    scanner := bufio.NewScanner(file)

    var trees [][]int

    for scanner.Scan() {
        currentLine := scanner.Text()
        splitLine := strings.Split(currentLine, "")
        var row []int

        for _, char := range splitLine {
            num, err := strconv.Atoi(char)
            if err != nil {
                log.Fatalf("Could not convert string %s to int", splitLine)
            }
            row = append(row, num)
        }

        trees = append(trees, row)
    }

    return trees
}

func printTrees(trees [][]int) {
    for _, row := range trees {
        for _, tree := range row {
            fmt.Printf("%d ", tree)
        }
        fmt.Printf("\n\n")
    }
}

func countVisibleTrees(trees [][]int) int {
    sum := 0
    for y, row := range trees {
        for x := range row {
            if isVisible(trees, x, y) {
                sum += 1
            }
        }
    }

    return sum
}

func isVisible(trees [][]int, x int, y int) bool {
    selectedTreeHeight := trees[y][x]

    // Border
    if x == 0 || y == 0 {
        return true
    }

    // Also Border
    if y == len(trees) || x == len(trees[0]) {
        return true
    }


    allLower := true
    // Left X-axis traversal. If _all_ elements are lower, it is visible
    for i := 0; i < x; i++ {
        if trees[y][i] >= selectedTreeHeight {
            allLower = false
        }
    }
    if allLower {
        return true
    }

    allLower = true
    // Right X-axis traversal. If _all_ elements are lower, it is visible
    for i := x + 1; i < len(trees[0]); i++ {
        if trees[y][i] >= selectedTreeHeight {
            allLower = false
        }
    }
    if allLower {
        return true
    }

    allLower = true
    // Top Y-axis traversal. If _all_ elements are lower, it is visible
    for i := 0; i < y; i++ {
        if trees[i][x] >= selectedTreeHeight {
            allLower = false
        }
    }
    if allLower {
        return true
    }

    allLower = true
    // Top Y-axis traversal. If _all_ elements are lower, it is visible
    for i := y + 1; i < len(trees); i++ {
        if trees[i][x] >= selectedTreeHeight {
            allLower = false
        }
    }
    if allLower {
        return true
    }
    // If it misses all visibility checks, then it is invisible
    return false
}

func main() {
    trees := loadTrees("input.txt")
    printTrees(trees)
    fmt.Printf("There are a total of %d trees visible", countVisibleTrees(trees))
}
