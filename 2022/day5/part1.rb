STACK = """    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3
"""

MOVES = """move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
"""

def parse_stack(stack)
  levels = stack.split("\n")
  base = levels.pop
  num_stacks = base.split(' ').map(&:to_i).max

  stacks = []
  for i in 1..num_stacks do
    current_stack = []
    levels.each do |level|
      current = level.slice!(0, 4).sub('[', '').sub(']', '').strip
      current_stack.insert(0, current) unless current == ''
    end

    stacks << current_stack.compact
  end

  stacks
end

def parse_moves(moves)
  moves = moves.split("\n")
  parsed_moves = []
  moves.each do |move|
    split = move.split(" ")
    number = split[1].to_i
    from = split[3].to_i
    to = split[5].to_i
    parsed_moves << [number, from, to]
  end

  parsed_moves
end

def carry_out_moves(stack, moveset)
  moveset.each do |move|
    number = move[0]
    from = move[1] - 1
    to = move[2] - 1

    # puts("Moving #{number} from #{move[1]} to #{move[2]}")
    number.times do
      stack[to].push(stack[from].pop)
    end
    # print_out_stacks(stack)
    # puts()
  end

  stack
end

def print_out_stacks(stacks)
  stacks.each do |stack|
    stack.each do |ele|
      print("#{ele}\t")
    end
    puts("\n")
  end
end

move_input = File.open("moves.txt").read
stacks_input = File.open("stacks.txt").read

moves = parse_moves(move_input)
stacks = parse_stack(stacks_input)

print_out_stacks(stacks)

# moves = parse_moves(MOVES)

moves.each do |move|
  print("#{move[0]} #{move[1]} #{move[2]}")
  puts()
end

results = carry_out_moves(stacks, moves)
print_out_stacks(results)
print(results.map(&:last).join)