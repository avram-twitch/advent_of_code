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

func printTrees(trees [][]int, xHighlight int, yHighlight int) {
    for y, row := range trees {
        for x, tree := range row {
            if x == xHighlight && y == yHighlight {
                fmt.Printf("*%d*", tree)
            } else {
                fmt.Printf("%d ", tree)
            }
        }
        fmt.Printf("\n\n")
    }
}

func maxScenicScore(trees [][]int) int {
    maxScore := 0
    for y, row := range trees {
        for x := range row {
            currentScore := scenicScore(trees, x, y)
            if maxScore < currentScore {
                maxScore = currentScore
            }
        }
    }
    return maxScore
}

func scenicScore(trees [][]int, x int, y int) int {
    var upScore, downScore, leftScore, rightScore int
    treeHeight := trees[y][x]
    // Look Up
    for i := y - 1; i > -1; i-- {
        currentHeight := trees[i][x]
        upScore += 1
        if currentHeight >= treeHeight {
            break
        }
    }
    // Look Down
    for i := y + 1; i < len(trees); i++ {
        currentHeight := trees[i][x]
        downScore += 1
        if currentHeight >= treeHeight {
            break
        }
    }
    // Look Left
    for i := x - 1; i > -1; i-- {
        currentHeight := trees[y][i]
        leftScore += 1
        if currentHeight >= treeHeight {
            break
        }
    }
    // Look Right
    for i := x + 1; i < len(trees[0]); i++ {
        currentHeight := trees[y][i]
            rightScore += 1
        if currentHeight >= treeHeight {
            break
        }
    }

    return upScore * downScore * leftScore * rightScore
}

func main() {
    trees := loadTrees("input.txt")
    fmt.Printf("There highest scenic score is %d\n", maxScenicScore(trees))
}
