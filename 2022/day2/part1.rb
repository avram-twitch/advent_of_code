RAW = """A Y
B X
C Z
"""

ENEMY_ROCK     = "A"
ENEMY_PAPER    = "B"
ENEMY_SCISSORS = "C"

MY_ROCK     = "X"
MY_PAPER    = "Y"
MY_SCISSORS = "Z"

class Round
  def initialize(enemy_move, my_move)
    @enemy_move = enemy_move
    @my_move = my_move
  end

  def my_score
    score = 0
    score += move_score
    score += outcome_score
  end

  def move_score
    case @my_move
    when MY_ROCK
      return 1
    when MY_PAPER
      return 2
    when MY_SCISSORS
      return 3
    else
      fail
    end
  end

  def outcome_score
    # DRAWS
    if @enemy_move == ENEMY_ROCK && @my_move == MY_ROCK
      return 3
    elsif @enemy_move == ENEMY_PAPER && @my_move == MY_PAPER
      return 3
    elsif @enemy_move == ENEMY_SCISSORS && @my_move == MY_SCISSORS
      return 3
    # WINS
    elsif @enemy_move == ENEMY_ROCK && @my_move == MY_PAPER
      return 6
    elsif @enemy_move == ENEMY_PAPER && @my_move == MY_SCISSORS
      return 6
    elsif @enemy_move == ENEMY_SCISSORS && @my_move == MY_ROCK
      return 6
    end

    # Everything remaining is a loss
    0
  end
end

def parse(input_txt)
  rows = input_txt.split("\n")
  rows.map do |row|
    puts("parsing #{row}")
    enemy_move, my_move = row.split(" ")
    Round.new(enemy_move, my_move)
  end
end


file = File.open("input.txt")
output = parse(file.read)

puts(output.inject(0) { |sum, round| sum + round.my_score })
