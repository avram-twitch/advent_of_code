
RAW = """Monkey 0:
Starting items: 79, 98
Operation: new = old * 19
Test: divisible by 23
  If true: throw to monkey 2
  If false: throw to monkey 3

Monkey 1:
Starting items: 54, 65, 75, 74
Operation: new = old + 6
Test: divisible by 19
  If true: throw to monkey 2
  If false: throw to monkey 0

Monkey 2:
Starting items: 79, 60, 97
Operation: new = old * old
Test: divisible by 13
  If true: throw to monkey 1
  If false: throw to monkey 3

Monkey 3:
Starting items: 74
Operation: new = old + 3
Test: divisible by 17
  If true: throw to monkey 0
  If false: throw to monkey 1"""

DIVIDING_WORRY_LEVEL = 1
MODULO = 1

class Game
  attr_accessor :rounds, :monkeys

  def initialize(rounds, monkeys)
    @rounds = rounds
    @monkeys = monkeys
  end

  def run_game
    rounds.times do |i|
      puts("Round #{i + 1}")
      monkeys.each do |monkey|
        monkey.run_turn(monkeys)
      end
      # print_monkeys
      # puts()
    end
  end

  def print_monkeys
    monkeys.each_with_index do |monkey, i|
      puts("Monkey #{i}: #{monkey.items}")
    end
  end

  def top_n_monkey_business(n)
    monkeys.map(&:times_inspected).max(n)
  end

  def print_inspections
    monkeys.each_with_index do |monkey, i|
      puts("Monkey #{i}: #{monkey.times_inspected}")
    end
  end
end

class Monkey
  attr_accessor :items,
                :operator,
                :operand,
                :test,
                :true_monkey,
                :false_monkey,
                :times_inspected

  def initialize(instructions)
    @times_inspected = 0
    # Starting Items
    @items = instructions[1].sub("Starting items: ", "").split(", ").map(&:to_i)

    # Operation
    operation = instructions[2].sub("Operation: new = ", "").split(" ")
    @operator = operation[1]
    @operand = operation[2]
    ## There is a case where the operation is old * old, so
    ## I just use the exponential
    if @operand == "old"
      @operator = "**"
      @operand = 2
    else
      @operand = @operand.to_i
    end

    # Test
    @test = instructions[3].gsub("Test: divisible by ", "").to_i

    # True Monkey
    @true_monkey = instructions[4].gsub("If true: throw to monkey ", "").to_i

    # False Monkey
    @false_monkey = instructions[5].gsub("If false: throw to monkey ", "").to_i
  end

  def to_s
    puts("Items: #{items}, Operator: #{operator}, Operand: #{operand}, test: #{test}, True Monkey: #{true_monkey}, False Monkey: #{false_monkey}")
  end

  def run_turn(monkeys)
    items.each do |item|
      @times_inspected += 1
      new_worry_level = item.public_send operator, operand
      new_worry_level /= DIVIDING_WORRY_LEVEL
      new_worry_level = new_worry_level % MODULO
      monkeys[true_monkey].items << new_worry_level if new_worry_level % test == 0
      monkeys[false_monkey].items << new_worry_level unless new_worry_level % test == 0
    end
    @items = []
  end
end

# monkey_instructions = RAW.split("\n\n")
monkey_instructions = File.open("input.txt").read.split("\n\n")

monkeys = monkey_instructions.map do |instructions|
  Monkey.new(instructions.split("\n"))
end

monkeys.each do |monkey|
  MODULO *= monkey.test
end

rounds = 10000
game = Game.new(rounds, monkeys)
game.run_game
puts(game.print_inspections)
puts(game.top_n_monkey_business(2).inject(:*))