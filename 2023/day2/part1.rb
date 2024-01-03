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

def find_games_that_match_constraints(games, constraint)
  filtered_games = {}
  games.each do |game, rounds|
    filtered_games[game] = rounds if game_within_constraints?(rounds, constraint)
  end
  filtered_games
end

def game_within_constraints?(rounds, constraint)
  rounds.each do |round|
    round.each do |color, num|
      return false if num > constraint[color.to_sym] 
    end
  end

  true
end

lines = file_to_array_of_lines("input.txt")
games = parse_lines_to_hash(lines)
constraint = {"red": 12, "green": 13, "blue": 14}

filtered_games = find_games_that_match_constraints(games, constraint)
puts(filtered_games.keys.map(&:to_i).sum)


