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
visible_count = 0
grid.each_with_index do |row, y|
  row.each_with_index do |tree, x|
    if x == 0 || y == 0 || x == max_x || y == max_y
      visible_count += 1
      next
    end

    # To Left
    to_left = row[0..x-1].filter { |val| val >= tree }
    if to_left.count == 0
      visible_count += 1
      next
    end

    # To Right
    to_right = row[x+1..max_x].filter { |val| val >= tree }
    if to_right.count == 0
      visible_count += 1
      next
    end

    column = grid.map { |row| row[x] }

    # To Up
    to_up = column[0..y-1].filter { |val| val >= tree }
    if to_up.count == 0
      visible_count += 1
      next
    end

    # To Down
    to_down = column[y+1..max_y].filter { |val| val >= tree }
    if to_down.count == 0
      visible_count += 1
      next
    end
  end
end

puts(visible_count)