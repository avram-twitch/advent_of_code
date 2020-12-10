package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

type Seat struct {
    row int
    col int
}

func file_to_seats(fp string) []Seat{
    seat_strings := read_file(fp)
    var seats []Seat
    for _, seat_string := range seat_strings {
        seats = append(seats, string_to_seat(seat_string))
    }
    return seats
}

func read_file (fp string) []string {
    f, err := os.Open(fp)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    var seats []string

    scanner := bufio.NewScanner(f)

    for scanner.Scan() {
        seats = append(seats, string(scanner.Text()))
    }

    return seats
}

func string_to_seat(seat string) Seat {
    row := get_row(seat)
    col := get_col(seat)
    return Seat{row, col}
}

func get_row(seat string) int {
    bin := byte(0b0)
    for _, c := range seat {
        if string(c) == "F" {
            bin = fl_shift(bin)
        } else if string(c) == "B" {
            bin = br_shift(bin)
        }
    }
    return int(bin)
}

func get_col(seat string) int {
    bin := byte(0b0)
    for _, c := range seat {
        if string(c) == "L" {
            bin = fl_shift(bin)
        } else if string(c) == "R" {
            bin = br_shift(bin)
        }
    }
    return int(bin)
}

// Add 0 if char is "F" or "L"
func fl_shift(b byte) byte {
	out := b
	out = out << 1
	return out
}

// Add 1 if char is "B" or "R"
func br_shift(b byte) byte {
	out := b
	out = out << 1
	out = out | 0b1
	return out
}

func seat_id(seat Seat) int {
    return (seat.row * 8) + seat.col
}

func main() {
    seats := file_to_seats("input1.txt")
    max_seat := 0
    for _, seat := range seats {
        curr_id := seat_id(seat)
        if curr_id > max_seat {
            max_seat = curr_id
        }
    }
    fmt.Println(max_seat)
}
