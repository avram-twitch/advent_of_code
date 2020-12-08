package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

var REQUIRED_FIELDS = []string{
    "byr", 
    "iyr",
    "eyr",
    "hgt",
    "hcl",
    "ecl",
    "pid",
}

func parse_file(fp string) []map[string]string {
    f, err := os.Open(fp)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    var passports []map[string]string
    scanner := bufio.NewScanner(f)

    buffer_string := ""
    for scanner.Scan() {
        curr_line := scanner.Text()
        if string(curr_line) != "" {
            buffer_string = buffer_string + " " + scanner.Text()
        } else {
            passport := parse_passport(buffer_string)
            passports = append(passports, passport)
            buffer_string = ""
        }
    }
    passport := parse_passport(buffer_string)
    passports = append(passports, passport)
    return passports
}

func parse_passport (raw string) map[string]string{
    passport := make(map[string]string)
    attributes := strings.Split(raw, " ")
    for _, attr := range attributes {
        key_value := strings.Split(attr, ":")
        if len(key_value) > 1 {
            passport[key_value[0]] = key_value[1]
        }
    }
    return passport
}

func get_valid_passports(passports []map[string]string) []map[string]string{
    var valid_passports []map[string]string
    for _, passport := range passports {
        if is_valid(passport) {
            valid_passports = append(valid_passports, passport)
        }
    }

    return valid_passports
}

func is_valid(passport map[string]string) bool {
    for _, key := range REQUIRED_FIELDS {
        _, ok := passport[key]
        if !ok {
            return false
        }
    }
    return true
}

func main() {
    passports := parse_file("input1.txt")
    valid_passports := get_valid_passports(passports)
    fmt.Println(len(valid_passports))
}
