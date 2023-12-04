lines = "1\n2\n3\n\n4\n5\n\n6\n7\n8\n9"
RAW = """
1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
"""

# split the string into groups of numbers
# groups = lines.split("\n\n")
groups = RAW.split("\n\n")

# initialize a variable to keep track of the highest sum
highest_sum = 0

# iterate over the groups of numbers
groups.each do |group|
  # split each group into individual numbers
  numbers = group.split("\n")

  # sum the numbers in the group
  sum = numbers.map(&:to_i).reduce(:+)

  # if the sum is higher than the current highest sum, update the variable
  if sum > highest_sum
    highest_sum = sum
  end
end

# print the highest sum
puts highest_sum
