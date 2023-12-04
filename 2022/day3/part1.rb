RAW = """vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw"""

class Rucksack
    def initialize(input_string)
        @first_half, @second_half = split_in_halves(input_string)
    end

    def split_in_halves(string)
        length = string.length
        half = length / 2
        first = string.slice(0, half)
        second = string.slice(half, length)
        return first, second
    end

    def identify_shared_items
        @first_half.chars & @second_half.chars
    end
end

def get_priority(character)
    if character == character.downcase
        return character.ord - 'a'.ord + 1
    else
        return character.ord - 'A'.ord + 27
    end
end

# inputs = RAW.split("\n")

file = File.open("input.txt")

inputs = file.read.split("\n")

priorities = 0

inputs.each do |input|
    rucksack = Rucksack.new(input)
    puts("Same item is #{rucksack.identify_shared_items}")
    priorities += get_priority(rucksack.identify_shared_items.first)
end
puts("Sum of priorities = #{priorities}")