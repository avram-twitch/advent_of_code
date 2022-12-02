RAW = """A Y
B X
C Z
"""

ROCK     = "A"
PAPER    = "B"
SCISSORS = "C"

LOSE = "X"
DRAW = "Y"
WIN  = "Z"

class Round
  def initialize(enemy_move, outcome)
    @enemy_move = enemy_move
    @outcome = outcome
  end

  def my_score
    score = 0
    score += move_score
    score += outcome_score
  end

  def my_move
    case @outcome
    when DRAW
      return @enemy_move
    when LOSE
      return what_loses_to @enemy_move
    when WIN
      return what_beats @enemy_move
    end
  end

  def what_beats(move)
    case move
    when ROCK
      return PAPER
    when PAPER
      return SCISSORS
    when SCISSORS
      return ROCK
    end
  end

  def what_loses_to(move)
    case move
    when ROCK
      return SCISSORS
    when PAPER
      return ROCK
    when SCISSORS
      return PAPER
    end
  end

  def move_score
    case my_move
    when ROCK
      return 1
    when PAPER
      return 2
    when SCISSORS
      return 3
    else
      fail
    end
  end

  def outcome_score
    return 0 if @outcome == LOSE
    return 3 if @outcome == DRAW
    6
  end
end

def parse(input_txt)
  rows = input_txt.split("\n")
  rows.map do |row|
    enemy_move, my_move = row.split(" ")
    Round.new(enemy_move, my_move)
  end
end


# output = parse(RAW)
file = File.open("input.txt")
output = parse(file.read)

puts(output.inject(0) { |sum, round| sum + round.my_score })
