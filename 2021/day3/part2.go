package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

//
// Not passing part 2 yet, value too low.
// Arriving at 2 final values for co measurement
//

func main() {
    diagnostic_report := read_file("input.txt")

    most_common_bits := get_most_common_bits(diagnostic_report)

    oxygen_rating, co_rating:= extract_oxygen_and_co_rating(diagnostic_report, most_common_bits)

    fmt.Printf("oxygen_rating: %d, co_rating: %d\n", oxygen_rating, co_rating)
    fmt.Printf("Multiplied: %d", oxygen_rating * co_rating)
}

func get_most_common_bits(diagnostic_report []string) ([]int) {
    bit_counter := make([]int, len(diagnostic_report[0]))

    for _, line := range diagnostic_report {
        bit_counter = count_binary_line(line, bit_counter)
    }

    var most_common_bits []int
    one_bit_more_numerous_threshold := (len(diagnostic_report) / 2) + (len(diagnostic_report) % 2) - 1

    for _, counter := range bit_counter {
        most_common_bits = append(most_common_bits, counter / one_bit_more_numerous_threshold)
    }

    return most_common_bits
}

func count_binary_line(line string, bit_counter []int) ([]int) {
    for i, char := range line {
        bit_counter[i] += int(char - '0')
    }
    return bit_counter
}

func extract_oxygen_and_co_rating(diagnostic_report []string, most_common_bits []int) (int, int) {

    oxygen_rating := filter_oxygen(diagnostic_report, most_common_bits)
    co_rating := filter_co(diagnostic_report, most_common_bits)

    return string_to_int(oxygen_rating), string_to_int(co_rating)
}

func filter_oxygen(diagnostic_report []string, most_common_bits []int) (string) {
    return filter(diagnostic_report, most_common_bits)
}

func filter_co(diagnostic_report []string, most_common_bits []int) (string) {
    var flipped_bits []int
    for _, bit := range most_common_bits {
        switch bit {
        case 1:
            flipped_bits = append(flipped_bits, 0)
        case 0:
            flipped_bits = append(flipped_bits, 1)
        }
    }
    return filter(diagnostic_report, flipped_bits)
}

func filter(diagnostic_report []string, most_common_bits []int) (string) {
    fmt.Printf("Bits: ")
    fmt.Print(most_common_bits)
    fmt.Print("\n")
    fmt.Print("Report: \n")
    fmt.Print(diagnostic_report)
    fmt.Print("\n")


    if len(diagnostic_report) == 1 {
        return diagnostic_report[0]
    } 

    var filtered_report []string

    for _, line := range diagnostic_report {
        shifts := len(most_common_bits) - 1
        line_as_int := string_to_int(line)
        line_as_int = line_as_int >> shifts

        if (line_as_int & 01) == most_common_bits[0] {
            filtered_report = append(filtered_report, line)
        }
    }

    if len(filtered_report) == 0 {
        return diagnostic_report[1]
    }

    return filter(filtered_report, most_common_bits[1:])
}

func string_to_int(s string) (int) {
    line_as_int, err := strconv.ParseInt(s, 2, 32)
    check(err)
    return int(line_as_int)
}

func check(err error) {
    if err != nil {
        log.Fatalf("readlines: %s", err)
    }
}

func read_file(fp string) ([]string) {
    f, err := os.Open(fp)
    check(err)
    defer f.Close()

    var lines []string
    scanner := bufio.NewScanner(f)

    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    check(scanner.Err())
    return lines
}
