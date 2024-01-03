package main

import (
    "bufio"
    "log"
    "fmt"
    "strings"
    "os"
)

var enemyPlayMap map[string]string
var myPlayMap map[string]string
var playValueMap map[string]int
var outcomesMap map[string]map[string]int

func read_file(filePath string) [][]string {
    f, err := os.Open(filePath)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    var rounds [][]string
    var round []string

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        line := scanner.Text()
        splitLine := strings.Split(line, " ")

        for _, play := range splitLine {
            round = append(round, play)
        }
        rounds = append(rounds, round)
        round = nil
    }

    return rounds
}

func init() {
    enemyPlayMap = map[string]string {
        "A": "Rock",
        "B": "Paper",
        "C": "Scissors",
    }
    myPlayMap = map[string]string {
        "X": "Rock",
        "Y": "Paper",
        "Z": "Scissors",
    }
    playValueMap = map[string]int {
        "Rock": 1,
        "Paper": 2,
        "Scissors": 3,
    }
    outcomesMap = make(map[string]map[string]int)
    outcomesMap["Rock"] = map[string]int{"Rock": 3, "Paper": 6, "Scissors": 0}
    outcomesMap["Paper"] = map[string]int{"Rock": 0, "Paper": 3, "Scissors": 6}
    outcomesMap["Scissors"] = map[string]int{"Rock": 6, "Paper": 0, "Scissors": 3}
}

func get_round_points(round []string) int {
    enemyPlay := enemyPlayMap[round[0]]
    myPlay := myPlayMap[round[1]]
    return outcomesMap[enemyPlay][myPlay] + playValueMap[myPlay]
}

func main() {
    // rounds := read_file("top_ten.txt")
    rounds := read_file("input.txt")
    sum := 0

    for _, round := range rounds {
        points := get_round_points(round)
        // fmt.Printf("Result %s %s: %d\n", round[0], round[1], points)
        sum += points
    }
    fmt.Printf("%d", sum)
}
