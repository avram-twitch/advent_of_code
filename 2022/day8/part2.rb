RAW = """30373
25512
65332
33549
35390"""

# First element --> Y axis
# Second element --> X axis

grid = File.open("input.txt").read.split("\n").map do |line|
  chars = line.chars
  chars.map(&:to_i)
end
# grid = RAW.split("\n").map do |line|
#   chars = line.chars
#   chars.map(&:to_i)
# end

max_y = grid.length - 1
max_x = grid[0].length - 1 

puts("Dimensions: x -- #{max_x} y -- #{max_y}")
scores = []
grid.each_with_index do |row, y|
  row.each_with_index do |tree, x|
    if x == 0 || y == 0 || x == max_x || y == max_y
      scores << 0
      next
    end

    # To Left
    to_left = row[0..x-1]
    left_score = 0
    to_left.reverse_each do |val|
      left_score += 1
      break if tree <= val
    end

    # To Right
    to_right = row[x+1..max_x]
    right_score = 0
    to_right.each do |val|
      right_score += 1
      break if tree <= val
    end

    column = grid.map { |row| row[x] }

    # To Up
    to_up = column[0..y-1]
    up_score = 0
    to_up.reverse_each do |val|
      up_score += 1
      break if tree <= val
    end

    # To Down
    to_down = column[y+1..max_y]
    down_score = 0
    to_down.each do |val|
      down_score += 1
      break if tree <= val
    end

    scores << left_score * right_score * up_score * down_score
  end
end

puts(scores.max)