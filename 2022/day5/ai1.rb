# # Parse the input and create the stacks and the moves
# stacks = [['Z', 'N'], ['M', 'C', 'D'], ['P']]
# moves = [[1, 2, 1], [3, 1, 3], [2, 2, 1], [1, 1, 2]]

# # Perform the moves in sequence
# moves.each do |from, n, to|
#   # Take the top n crates from the source stack and remove them
#   crates = stacks[from-1].pop(n)

#   # Add the crates to the destination stack
#   stacks[to-1].concat(crates)
# end

# # Output the top crates of each stack
# puts stacks.map { |stack| stack[-1] }.join

# TRY 2
# Parse the input and create the stacks and the moves
# stacks = [['Z', 'N'], ['M', 'C', 'D'], ['P']]
# moves = [[1, 2, 1], [3, 1, 3], [2, 2, 1], [1, 1, 2]]
# 
# # Perform the moves in sequence
# moves.each do |from, n, to|
#   # Take the top n crates from the source stack and remove them
#   crates = []
#   n.times { crates << stacks[from-1].shift }
# 
#   # Add the crates to the destination stack
#   stacks[to-1].concat(crates)
# end
# 
# # Output the top crates of each stack
# puts stacks.map { |stack| stack[-1] }.join

# ## TRY 3
# # Parse the input and create the stacks and the moves
# stacks = [['Z', 'N'], ['M', 'C', 'D'], ['P']]
# moves = [[1, 2, 1], [3, 1, 3], [2, 2, 1], [1, 1, 2]]

# # Perform the moves in sequence
# moves.each do |from, n, to|
#   # Take the top n crates from the source stack and remove them
#   crates = []
#   n.times { crates << stacks[from-1].shift }

#   # Add the crates to the destination stack
#   stacks[to-1].unshift(*crates)
# end

# # Output the top crates of each stack
# puts stacks.map { |stack| stack[-1] }.join

# ## TRY 4
# # Parse the input and create the stacks and the moves
# stacks = [['Z', 'N'], ['M', 'C', 'D'], ['P']]
# moves = [[2, 1, 1], [1, 3, 3], [2, 2, 1], [1, 1, 2]]

# # Perform the moves in sequence
# moves.each do |n, from, to|
#   # Take the top n crates from the source stack and remove them
#   crates = []
#   n.times { crates << stacks[from-1].shift }

#   # Add the crates to the destination stack
#   stacks[to-1].unshift(*crates)
# end

# # Output the top crates of each stack
# puts stacks.map { |stack| stack[-1] }.join

# ## Try 5
# # Parse the input and create the stacks and the moves
# stacks = [['Z', 'N'], ['M', 'C', 'D'], ['P']]
# moves = [[2, 1, 1], [1, 3, 3], [2, 2, 1], [1, 1, 2]]

# # Perform the moves in sequence
# moves.each do |n, from, to|
#   # Take the top n crates from the source stack and remove them
#   crates = []
#   n.times { crates << stacks[from-1].pop }

#   # Add the crates to the destination stack
#   stacks[to-1].append(*crates)
# end

# # Output the top crates of each stack
# puts stacks.map { |stack| stack[-1] }.join

## TRY 6
# Parse the input and create the stacks and the moves
stacks = [['Z', 'N'], ['M', 'C', 'D'], ['P']]
moves = [[1, 2, 1], [3, 1, 3], [2, 2, 1], [1, 1, 2]]

# Perform the moves in sequence
moves.each do |n, from, to|
  # Take the top n crates from the source stack and remove them
  crates = []
  n.times { crates << stacks[from-1].pop }

  # Add the crates to the destination stack
  stacks[to-1].append(*crates)
end

# Output the top crates of each stack
puts stacks.map { |stack| stack[-1] }.join
