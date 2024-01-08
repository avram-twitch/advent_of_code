package main

import (
    "os"
    "log"
    "bufio"
    "strings"
    "strconv"
    "fmt"
)

const STACK = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3
` 

const MOVES = `move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`

type Instruction struct {
    quantity int
    from int
    to int
}

func parseMoves(fp string) []Instruction {
    file, err := os.Open(fp)

    var instructions []Instruction

    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        text := scanner.Text()
        splitText := strings.Split(text, " ")
        quantity, err := strconv.Atoi(splitText[1])
        if err != nil {
            log.Fatal(err)
        }
        from, err := strconv.Atoi(splitText[3])
        if err != nil {
            log.Fatal(err)
        }
        to, err := strconv.Atoi(splitText[5])
        if err != nil {
            log.Fatal(err)
        }
        instruction := Instruction{quantity: quantity, from: from, to: to}
        instructions = append(instructions, instruction)
    }

    return instructions
}

func parseStack(fp string) [][]string {
    content, err := os.ReadFile(fp)

    if err != nil {
        log.Fatal(err)
    }

    lines := strings.Split(string(content), "\n")
    numberedStacks := strings.Split(lines[len(lines) - 1], " ")
    numberOfStacks, err := strconv.Atoi(numberedStacks[len(numberedStacks) - 2])

    if err != nil {
        log.Fatal(err)
    }

    var stacks [][]string

    for i := 0; i < numberOfStacks; i++ {
        var currentStack []string

        for j, line := range lines {
            if j == len(lines) - 1 {
                continue
            }
            offset := (i * 4) + 1 // An item comes in a chunk of 3 ("[A]")
            currentStack = append([]string{string(line[offset])}, currentStack...)
        }
        stacks = append(stacks, currentStack)
    }
    return stacks
}

func printStacks(stacks [][]string) {
    for i, stack := range stacks {
        fmt.Printf("%d: ", i)
        for _, element := range stack {
            fmt.Printf(" %s ", element)
        }
        fmt.Printf("\n")
    }
}

func executeInstructions(stacks [][]string, instructions []Instruction) [][]string {
    for _, instruction := range instructions {
        fmt.Printf("Executing quantity: %d, From %d, To %d\n", instruction.quantity, instruction.from, instruction.to)
        for i := 0; i < instruction.quantity; i++ {
            toStack := stacks[instruction.to - 1]
            fromStack := stacks[instruction.from - 1]
            movedValue := fromStack[len(fromStack) - 1]
            toStack = append(toStack, movedValue)
            fromStack = fromStack[:len(fromStack) - 1]
        }
        printStacks(stacks)
    }

    return stacks
}

func main() {
    stacks := parseStack("./stacks.txt")
    instructions := parseMoves("./top_ten.txt")
    printStacks(stacks)
    stacks = executeInstructions(stacks, instructions)
    fmt.Printf("AFTER RUN\n")
    printStacks(stacks)
}
