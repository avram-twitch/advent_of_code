package main

func (r Range) Has(n int) bool {
    return n >= r.Start && n <= r.End
}

func Part1(input Input) int {
    freshCount := 0

    for _, id := range input.IDs {
        for _, r := range input.Ranges {
            if r.Has(id) {
                freshCount++
                break
            }
        }
    }
    return freshCount
}
