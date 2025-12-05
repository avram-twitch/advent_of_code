package main

func (r Range) Size() int {
    return r.End - r.Start + 1
}

func (r1 Range) Overlaps(r2 Range) bool {
    return r1.Start <= r2.End && r2.Start <= r1.End
}

func (r1 Range) Combine(r2 Range) Range {
    start := min(r1.Start, r2.Start)
    end := max(r1.End, r2.End)

    return Range{Start: start, End: end}
}

// Combines range if possible
func mergeRange(ranges []Range, newR Range) []Range {
    merged := newR
    kept := make([]Range, 0, len(ranges))

    for _, r := range ranges {
        // Need to keep a running merge to ensure if there are multiple
        // range overlaps we combine those too
        if r.Overlaps(merged) {
            merged = merged.Combine(r)
        } else {
            kept = append(kept, r)
        }
    }

    return append(kept, merged)
}

func Part2(input Input) int {
    var ranges []Range

    for _, r := range input.Ranges {
        ranges = mergeRange(ranges, r)
    }

    sum := 0

    for _, r := range ranges {
        sum += r.Size()
    }


    return sum
}
