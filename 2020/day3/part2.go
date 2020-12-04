package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func read_file (fp string) [][]string {
    f, err := os.Open(fp)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    var grid [][]string

    scanner := bufio.NewScanner(f)
    row := 0

    for scanner.Scan() {
        grid = append(grid, []string{})
        for _, c := range scanner.Text() {
            grid[row] = append(grid[row], string(c))
        }
        row += 1
    }

    return grid
}

func print_grid (grid [][]string) {
    for _, row := range grid {
        for _, c := range row {
            fmt.Print(c)
        }
        fmt.Print("\n")
    }
}

func count_collisions(grid [][]string, dx int, dy int) int {
    curr_x := 0
    curr_y := 0
    hill_length := len(grid)
    max_x := len(grid[0])

    collision_count := 0

    for curr_y < hill_length {
        if grid[curr_y][curr_x] == "#" {
            collision_count += 1
        }
        curr_y += dy
        curr_x += dx
        if curr_x >= max_x {
            curr_x -= max_x
        }
    }
    return collision_count
}

func main() {
    grid := read_file("input1.txt")
    slopes := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
    product := 1
    for _, slope := range slopes {
        dx := slope[0]
        dy := slope[1]
        count := count_collisions(grid, dx, dy)
        product *= count_collisions(grid, dx, dy)
        fmt.Printf("For %dx%d the collision count is %d. Product is at %d\n", dx, dy, count, product)
    }
    fmt.Println(product)
}
