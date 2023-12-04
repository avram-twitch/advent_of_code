
RAW = """vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw"""

class Group
    def initialize(input_string1, input_string2, input_string3)
        @rucksack1 = Rucksack.new(input_string1)
        @rucksack2 = Rucksack.new(input_string2)
        @rucksack3 = Rucksack.new(input_string3)
    end

    def identify_badge_item
        @rucksack1.contents.chars & @rucksack2.contents.chars & @rucksack3.contents.chars
    end
end

class Rucksack
    def initialize(input_string)
        @contents = input_string
        @first_half, @second_half = split_in_halves(input_string)
    end

    def contents
        @contents
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

inputs.each_slice(3) do |input|
    group = Group.new(input[0], input[1], input[2])
    puts("Same item is #{group.identify_badge_item}")
    priorities += get_priority(group.identify_badge_item.first)
end
puts("Sum of priorities = #{priorities}")