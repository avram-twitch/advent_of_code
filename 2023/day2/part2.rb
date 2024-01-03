def file_to_array_of_lines(file_name)
  File.readlines(file_name)
end

def parse_lines_to_hash(lines)
  games = {}

  lines.each do |line|
    game, rounds = line.split(":")
    game = game.scan(/\d/).join("")
    games[game] = parse_round_to_list_of_hashes(rounds)
  end

  games
end

def parse_round_to_list_of_hashes(rounds)
  split_rounds = rounds.strip.split(";")
  parsed_rounds = []
  split_rounds.each do |round|
    round_hash = {}
    pulls = round.split(",")
    pulls.each do |pull|
      number, color = pull.strip.split(" ")
      color = color.strip
      number = number.strip.to_i
      round_hash[color] = number
    end

    parsed_rounds << round_hash
  end
  parsed_rounds
end

def generate_min_sets(games)
  min_sets = []
  games.each do |_, rounds|
    min_set = {}
    rounds.each do |round|
      round.each do |color, num|
        color = color.strip
        min_set.fetch(color, num)
        min_set[color] = num if num.to_i > min_set[color].to_i
      end
    end

    # puts("Round: #{rounds} converted to #{min_set}")
    min_sets << min_set
  end
  min_sets
end

lines = file_to_array_of_lines("input.txt")
games = parse_lines_to_hash(lines)
min_sets = generate_min_sets(games)
powers = min_sets.map do |min_set|
  power = 1
  puts(min_set)
  min_set.each do |color, num|
    puts("#{color}: #{num}")
    power *= num
  end
  power
end
puts(powers.sum)


