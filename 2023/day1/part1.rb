TEXT_TO_NUMBER = {
  "zero": 0,
  "one": 1,
  "two": 2,
  "three": 3,
  "four": 4,
  "five": 5,
  "six": 6,
  "seven": 7,
  "eight": 8,
  "nine": 9
}

def file_to_array_of_lines(file_name)
  File.readlines(file_name)
end

def sub_letter_numbers(lines)
  lines.map do |line|
    current_line = line.dup
    counter = 0
    while true do
      first_number = find_first_number(current_line)
      break if first_number.nil?
      value = TEXT_TO_NUMBER[first_number]
      current_line.gsub!(/#{first_number}/, "#{first_number[0]}#{value.to_s}#{first_number[first_number.length - 1]}")
    end
    current_line
  end
end

def find_first_number(text)
  first_position = nil
  first_number = nil
  TEXT_TO_NUMBER.keys.each do |key|
    found_position = text.index(key.to_s)
    next if found_position.nil?

    if first_number.nil? || found_position < first_position
      first_position = found_position
      first_number = key
    end
  end

  return first_number
end

def filter_nonnumerics(lines)
  lines.map { |line| line.gsub(/[^\d]/, "") }
end

def extract_first_and_last_numbers(lines)
  lines.map do |line|
    (line[0] + line[line.length - 1]).to_i
  end
end

lines = file_to_array_of_lines("./input.txt")
# lines = file_to_array_of_lines("./test.txt")
subbed = sub_letter_numbers(lines)
filtered = filter_nonnumerics(subbed)
first_and_lasts = extract_first_and_last_numbers(filtered)

20.times do |i|
  puts(lines[i])
  puts(subbed[i])
  puts(filtered[i])
  puts(first_and_lasts[i])
  puts
end
puts(first_and_lasts.sum)
