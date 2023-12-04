RAW = """R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2"""

X_AXIS = 0
Y_AXIS = 1
NUMBER_OF_KNOTS = 10

def touching?(head_position, tail_position)
  head_x, head_y = head_position
  tail_x, tail_y = tail_position

  x_delta = (head_x - tail_x).abs
  y_delta = (head_y - tail_y).abs

  x_delta <= 1 && y_delta <= 1
end

def update_tail(head_position, tail_position)
  return tail_position if touching?(head_position, tail_position)

  head_x, head_y = head_position
  tail_x, tail_y = tail_position

  x_delta = head_x - tail_x
  y_delta = head_y - tail_y

  new_tail_x = tail_x
  new_tail_y = tail_y

  # Same x coord
  if x_delta == 0
    new_tail_y = head_y - 1 if head_y > tail_y
    new_tail_y = head_y + 1 if head_y < tail_y
  end

  # Same y coord
  if y_delta == 0
    new_tail_x = head_x - 1 if head_x > tail_x
    new_tail_x = head_x + 1 if head_x < tail_x
  end

  # Needs diagonal movement
  if (y_delta.abs + x_delta.abs) > 2
    # puts("moving Diagonally")
    # print("#{head_x}, #{head_y} -- #{tail_x}, #{tail_y}\n")
    new_tail_y = tail_y + 1 if head_y > tail_y
    new_tail_y = tail_y - 1 if head_y < tail_y

    new_tail_x = tail_x + 1 if head_x > tail_x
    new_tail_x = tail_x - 1 if head_x < tail_x
  end

  [new_tail_x, new_tail_y]
end

def get_axis_and_sign_for_movement(direction)
  axis = X_AXIS
  sign = 1
  case direction
  when "R"
    axis = X_AXIS
    sign = 1
  when "L"
    axis = X_AXIS
    sign = -1
  when "U"
    axis = Y_AXIS
    sign = 1
  when "D"
    axis = Y_AXIS
    sign = -1
  end

  [axis, sign]
end

def print_grid(max_x, max_y, all_knots)
  max_y.times do |_y|
    max_x.times do |x|
      y = max_y - 1 - _y
      coord = [x, y]

      something_matched = false
      all_knots.each_with_index do |knot, i|
        if knot == coord
          something_matched = true
          if i == 0
            print("H ")
            break
          end

          # Tail
          if i == NUMBER_OF_KNOTS - 1
            print("T ")
            break
          end

          print("#{i + 1} ")
          break
        end
      end

      print(". ") unless something_matched
    end
    puts()
  end
end

tail_visited = []

all_knots = []
NUMBER_OF_KNOTS.times do
  all_knots << [0, 0]
end
head_position = all_knots[0]
tail_position = all_knots[NUMBER_OF_KNOTS - 1]
tail_visited << tail_position

instructions = File.open("input.txt").read.split("\n").map do |line|
  direction, steps = line.split(" ")
  [direction, steps.to_i]
end
# instructions = RAW.split("\n").map do |line|
#   direction, steps = line.split(" ")
#   [direction, steps.to_i]
# end

# print_grid(6, 5, all_knots)
# puts()

instructions.each do |instruction|
  direction, steps = instruction
  axis, sign = get_axis_and_sign_for_movement(direction)

  steps.times do
    all_knots.each_with_index do |current_knot, i|
      # Head
      if i == 0
        current_knot[axis] += (1 * sign)
        all_knots[i] = current_knot
        next
      end

      if i == 1
        previous_moved_knot = all_knots[i - 1]
        new_current_knot = update_tail(previous_moved_knot, current_knot)
        all_knots[i] = new_current_knot
        next
      end
      previous_moved_knot = all_knots[i - 1]
      new_current_knot = update_tail(previous_moved_knot, current_knot)
      all_knots[i] = new_current_knot

      # Tail
      if i == NUMBER_OF_KNOTS - 1
        tail_visited << new_current_knot
      end
    end
    head_position = all_knots[0]
    tail_position = all_knots[NUMBER_OF_KNOTS - 1]
    # print_grid(6, 5, all_knots)
    # puts()
  end
end

tail_visited.each do |position|
  print("(#{position[0]}, #{position[1]})")
  puts()
end
puts(tail_visited.uniq.count)
