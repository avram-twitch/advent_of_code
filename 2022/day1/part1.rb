RAW = """
1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
"""

class Elf
  def initialize(calories)
    @calories = calories
  end

  def total_calories
    @total_calories ||= @calories.sum
  end
end


def parse(raw_string)
  out = []
  raw_string.split("\n\n").each do |group_of_calories|
    list_of_categories = group_of_calories.split("\n").map(&:to_i)
    out << Elf.new(list_of_categories)
  end
  out
end

file = File.open("input.txt")
# elves = parse(RAW)
elves = parse(file.read)
puts(elves.max { |a, b| a.total_calories <=> b.total_calories }.total_calories)

# Part 2

sum = 0
elves.max(3) { |a, b| a.total_calories <=> b.total_calories }.each do |elf|
  sum += elf.total_calories
end
puts(sum)
