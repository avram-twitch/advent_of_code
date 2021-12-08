package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)

const (
    MARKED_SPOT = -1
)

func main() {
    call_numbers, bingo_cards := read_file("input.txt")

    var last_winning_card [][]int
    var last_winning_number int
    won_count := 0
    total_cards := len(bingo_cards)
    won_cards := make([]bool, total_cards)
    last_card_won := false

    for _, call_number := range call_numbers {
        if last_card_won {
            break
        }

        for i, card := range bingo_cards {
            card = mark_call_number(call_number, card)
            card_has_won := check_if_card_has_won(card)
            if card_has_won && !won_cards[i] {
                won_count += 1
                won_cards[i] = true
            }

            if card_has_won && won_count == total_cards {
                last_winning_number = call_number
                last_winning_card = card
                last_card_won = true
                break
            }
        }
    }

    pretty_print_card(last_winning_card)
    fmt.Print(multiply_winning(last_winning_number, last_winning_card))
}

func multiply_winning(number int, card [][]int) (int) {
    sum := 0

    for _, line := range card {
        for _, num := range line {
            if num != MARKED_SPOT {
                sum += num
            }
        }
    }

    return number * sum
}

func mark_call_number(call_number int, card [][]int) ([][]int) {
    for i, line := range card {
        var new_line []int
        for j, card_number := range line {
            if card_number == call_number {
                new_line = line
                new_line[j] = MARKED_SPOT
                card[i] = new_line
            }
        }
    }
    return card
}

func check_if_card_has_won(card [][]int) (bool) {
    // horizontal check
    for _, line := range card {
        won := true
        for _, num := range line {
            won = won && num == MARKED_SPOT
        }

        if won {
            return won
        }
    }

    // vertical check
    for i, line := range card {
        won := true
        for j, _ := range line {
            won = won && card[j][i] == MARKED_SPOT
        }

        if won {
            return won
        }
    }
    return false
}

func pretty_print_card(card [][]int) {
    for _, line := range card {
        fmt.Print(line)
        fmt.Print("\n")
    }
    fmt.Print("\n")
}

func read_file(fp string) ([]int, [][][]int) {
    first_line := true
    f, err := os.Open(fp)
    check(err)
    defer f.Close()

    var call_numbers []int
    var bingo_cards [][][]int
    scanner := bufio.NewScanner(f)

    for scanner.Scan() {
        if first_line {
            call_numbers = parse_call_numbers_line(scanner.Text())
            first_line = false
        } else {
            if scanner.Text() == "" {
                var new_card [][]int
                bingo_cards = append(bingo_cards, new_card)
            } else {
                bingo_card_numbers := parse_bingo_card_line(scanner.Text())
                last_card_position := len(bingo_cards) - 1
                bingo_cards[last_card_position] = append(bingo_cards[last_card_position], bingo_card_numbers)
            }
        }
    }

    check(scanner.Err())
    return call_numbers, bingo_cards
}

func parse_call_numbers_line(text string) ([]int) {
    split_text := strings.Split(text, ",")
    return parse_line(split_text, ",")
}

func parse_bingo_card_line(text string) ([]int) {
    split_text := strings.Fields(text)
    return parse_line(split_text, " ")
}

func parse_line(split_text []string, delimeter string) ([]int) {
    var line_numbers []int

    for _, num := range split_text {
        line_numbers = append(line_numbers, atoi(num))
    }
    
    return line_numbers
}

func atoi(s string) (int) {
    val, err := strconv.Atoi(s)
    check(err)
    return val
}

func check(err error) {
    if err != nil {
        log.Fatalf("readlines: %s", err)
    }
}
