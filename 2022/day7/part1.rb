RAW = """$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
"""

PROMPT = "$ "
MAX_SIZE = 70_000_000
STORAGE = 70_000_000
TARGET_FREE = 30_000_000
class Folder
  attr_accessor :children
  attr_reader :path, :parent

  def initialize(path, parent)
    @path = path
    @parent = parent
    @children = []
  end

  def to_s
    "{FOLDER Path: #{path}, Parent: #{parent&.path}, Children_count: #{children.length}, Size: #{size} }"
  end

  def size
     children.map(&:size).sum
  end

  def has_child?(name)
    children.map(&:path).include?(name)
  end

  def get_child(name)
    children.select { |child| child.path == name }.first
  end

  def eligible_for_deletion?
    size < MAX_SIZE
  end
end

class MyFile
  attr_reader :path, :size

  def initialize(path, size)
    @path = path
    @size = size
  end

  def to_s
    "{FILE Path: #{path}, Size #{size} }"
  end

end

class CommandParser
  attr_accessor :current_folder
  attr_reader :terminal_lines, :root

  def initialize(list_of_terminal_lines)
    @terminal_lines = list_of_terminal_lines
    @pwd = nil
  end

  def get_all_eligible_folders!
    _recursive_check(@root)
  end

  def _recursive_check(current)
    if current.class == MyFile
      return []
    end

    # if current.eligible_for_deletion?
    #   return [current]
    # end

    eligible_folders = []
    current.children.each do |child|
      eligible_folders << _recursive_check(child)
    end

    eligible_folders << current if current.eligible_for_deletion?
    eligible_folders.flatten
  end

  def execute_commands!
    terminal_lines.each do |line|
      puts("Running: #{line}")

      if is_cd?(line)
        cd(line)
        next
      end

      next if is_ls?(line)

      parse_ls_output!(line)
    end
  end

  def cd(line)
    dir = line.split(" ")[2]
    if dir == ".."
      @pwd = @pwd.parent
    elsif @pwd != nil
      if @pwd.path == "/"
        new_path = "#{@pwd.path}#{dir}"
      else
        new_path = "#{@pwd.path}/#{dir}"
      end

      if @pwd.has_child?(new_path)
        @pwd = @pwd.get_child(new_path)
      else
        new_folder = Folder.new(new_path, @pwd)
        puts("Creating new folder: #{new_folder}")
        @pwd.children << new_folder
        @pwd = new_folder
      end
    else
      # Should be root
      @root = Folder.new(dir, nil)
        puts("Creating new folder: #{@root}")
      @pwd = @root
    end
  end

  def parse_ls_output!(line)
    split = line.split(" ")
    first = split[0]
    second = split[1]

    if @pwd.path == "/"
      new_path = "#{@pwd.path}#{second}"
    else
      new_path = "#{@pwd.path}/#{second}"
    end

    if first == "dir"
      new_item = Folder.new(new_path, @pwd)
      puts("Creating new folder: #{new_item}")
    else
      new_item = MyFile.new(new_path, first.to_i)
      puts("Creating new file: #{new_item}")
    end
    @pwd.children << new_item
  end

  def is_cd?(line)
    line[0..3] == PROMPT + "cd"
  end

  def is_ls?(line)
    line[0..3] == PROMPT + "ls"
  end
end

# commands = RAW.split("\n")
commands = File.open("input.txt").read.split("\n")

parser = CommandParser.new(commands)
parser.execute_commands!

goal_size = TARGET_FREE - (STORAGE - parser.root.size)
candidate = nil
candidate_size = STORAGE

puts("\n" * 5)
puts("Goal size: #{goal_size}")
candidates = []
parser.get_all_eligible_folders!.each do |folder|
  puts("#{folder.path}: #{folder.size}")
  candidates << folder if folder.size > goal_size
end

puts(candidates)
puts(candidates.map(&:size).min)