RAW = [
  'mjqjpqmgbljsphdztnvjfqwrcgsmlb',
  'bvwbjplbgvbhsrlpgdmjqwftvncz',
  'nppdvjthqldpwncqszvftbrmjlhg',
  'nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg',
  'zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw'
]

class Message
  NEEDED_UNIQUE_CHARACTERS = 14

  def initialize(input)
    @raw_input = input
  end

  def characters_needed_to_process
    processed_chars = []
    current_count = 0

    @raw_input.chars.each do |char|
      current_count += 1
      processed_chars << char
      processed_chars = processed_chars.drop(1) if processed_chars.length > NEEDED_UNIQUE_CHARACTERS
      return current_count if processed_chars.uniq.length == NEEDED_UNIQUE_CHARACTERS
    end

    current_count
  end
end

puzzle_input = File.open("input.txt").read
message = Message.new(puzzle_input)
puts(message.characters_needed_to_process)

# RAW.each do |raw|
#   message = Message.new(raw)
#   puts(message.characters_needed_to_process)
# end