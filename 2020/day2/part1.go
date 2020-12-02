package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
)

func check_error(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func read_file(fp string) ([][]string, error) {
    f, err := os.Open(fp)
    check_error(err)
    defer f.Close()

    var lines [][]string
    scanner := bufio.NewScanner(f)

    for scanner.Scan() {
        vals := strings.Split(scanner.Text(), " ")
        vals[1] = strings.Split(vals[1], ":")[0] // Get just the letter, not the : as well
        lines = append(lines, vals)
    }

    return lines, scanner.Err()
}

func is_valid(pass_info []string) bool {
    char_range := strings.Split(pass_info[0], "-")
    count_low, _ := strconv.Atoi(char_range[0])
    count_high, _ := strconv.Atoi(char_range[1])
    char := pass_info[1]
    password := pass_info[2]
    char_count := strings.Count(password, char)
    return char_count >= count_low && char_count <= count_high
}

func main() {
    lines, err := read_file("input1.txt")
    check_error(err)
    count := 0

    for _, password_info := range lines {
        if is_valid(password_info) {
            fmt.Println(password_info)
            count += 1
        }
    }
    fmt.Println(count)
}
