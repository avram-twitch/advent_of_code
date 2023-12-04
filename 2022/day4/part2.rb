
RAW = """2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
"""

class Elf
  def initialize(input)
   @raw_input = input
    @start, @finish = input.split("-").map(&:to_i)
  end

  def raw_input
    @raw_input
  end

  def start
    @start
  end

  def finish
    @finish
  end

  def contains?(elf)
    elf.start >= start && elf.finish <= finish
  end

  def includes?(elf)
    return true if elf.start >= start && elf.start <= finish
    elf.finish >= start && elf.finish <= finish
  end
end

file = File.open("input.txt")
lines = file.read.split("\n")

included_count = 0

lines.each do |line|
  first, second = line.split(",")
  elf1 = Elf.new(first)
  elf2 = Elf.new(second)

  included_count += 1 if elf1.includes?(elf2) || elf2.includes?(elf1)
end

puts(included_count)