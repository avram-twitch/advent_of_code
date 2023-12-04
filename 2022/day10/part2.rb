RAW = """addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop"""

ADD_CYCLES = 2
NOOP_CYCLES = 1
ROW_LENGTH = 40

class Cpu
  attr_accessor :x_values,
                :instructions,
                :number_of_cycles,
                :current_cycle

  def initialize(instructions)
    @x_values = [1]
    @instructions = instructions
    @current_cycle = 0
  end

  def perform_instructions
    instructions.each do |instruction|
      split = instruction.split(" ")
      if split[0] == "noop"
        NOOP_CYCLES.times do
          increment_cycle!
          x_values << x_values.last
        end
      end

      if split[0] == "addx"
        ADD_CYCLES.times do |i|
          increment_cycle!
          current_value = x_values.last
          current_value += split[1].to_i if (i + 1) == ADD_CYCLES
          x_values << current_value
        end
      end
    end

    @number_of_cycles = current_cycle
  end

  def increment_cycle!
    self.current_cycle += 1
    crt_print(current_cycle, x_values.last)
  end

  def signal_strength_at(pos)
    x_values[pos - 1] * pos
  end
end

def crt_print(current_cycle, sprite_position)
  current_position = (current_cycle - 1) % ROW_LENGTH
  is_lit = ((sprite_position - 1 ) == current_position) ||
    (sprite_position == current_position) ||
    ((sprite_position + 1 ) == current_position)

  drawn = is_lit ? "#" : "."
  # puts("Cycle #{current_cycle}, X at #{sprite_position}, draws at postion #{current_position}: #{drawn}")
  # puts("During cycle #{current_cycle}: CRT draws #{drawn} in position #{current_cycle - 1}")
  # puts()
  print(drawn)
  puts() if current_cycle % ROW_LENGTH == 0
end

# instructions = RAW.split("\n")
instructions = File.open("input.txt").read.split("\n")

computer = Cpu.new(instructions)
computer.perform_instructions

current_check = 20
all = []

while current_check <= computer.number_of_cycles
  all << computer.signal_strength_at(current_check)
  current_check += 40
end

puts(all.sum)