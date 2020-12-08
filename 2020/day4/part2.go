package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
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

// Validations

func byr_valid(val string) bool {
    if val == "" {
        return false
    }
    byr, _ := strconv.Atoi(val)
    return byr >= 1920 && byr <= 2002
}

func iyr_valid(val string) bool {
    if val == "" {
        return false
    }
    iyr, _ := strconv.Atoi(val)
    return iyr >= 2010 && iyr <= 2020
}

func eyr_valid(val string) bool {
    if val == "" {
        return false
    }
    eyr, _ := strconv.Atoi(val)
    return eyr >= 2020 && eyr <= 2030
}

func hgt_valid(val string) bool {
    if val == "" {
        return false
    }
    hgt, _ := strconv.Atoi(val[:len(val) - 2])
    if val[len(val) - 2:] == "in" {
        return hgt >= 59 && hgt <= 76
    }
    if val[len(val) - 2:] == "cm" {
        return hgt >= 150 && hgt <= 193
    }
    return false
}

func hcl_valid(val string) bool {
    if val == "" {
        return false
    }
    // may need to check that these values are 0-9 and a-f
    return string(val[0]) == "#" && len(val) == 7
}

func ecl_valid(val string) bool {
    if val == "" {
        return false
    }
    valid_ecl := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
    for _, ecl := range valid_ecl {
        if ecl == val {
            return true
        }
    }
    return false
}

func pid_valid(val string) bool {
    if val == "" {
        return false
    }
    _, err := strconv.Atoi(val)
    return err == nil && len(val) == 9
}

func is_valid(passport map[string]string) bool {
    validations := map[string]func(string)bool {
        "byr": byr_valid,
        "iyr": iyr_valid,
        "eyr": eyr_valid,
        "hgt": hgt_valid,
        "hcl": hcl_valid,
        "ecl": ecl_valid,
        "pid": pid_valid,
    }

    for _, key := range REQUIRED_FIELDS {
        val, ok := passport[key]
        if !ok {
            return false
        }
        if !validations[key](val) {
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
