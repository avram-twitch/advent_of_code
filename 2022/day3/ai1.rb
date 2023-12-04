# TRY 1
# # Read rucksack list from input
# rucksacks = File.readlines("input.txt")
# 
# # Initialize total priority to 0
# total_priority = 0
# 
# # Iterate through each rucksack
# rucksacks.each do |rucksack|
#   # Get the first and second compartment of the rucksack
#   comp1 = rucksack[0, rucksack.length/2]
#   comp2 = rucksack[rucksack.length/2, rucksack.length]
# 
#   # Check if any characters appear in both compartments
#   common_chars = comp1.chars.select { |c| comp2.include?(c) }
# 
#   # If any characters are common, add their priority to the total
#   if common_chars.any?
#     total_priority += common_chars.first.downcase.ord - 'a'.ord + 1
#   end
# end
# 
# # Print the total priority
# puts total_priority

# # Read rucksack list from input
# rucksacks = File.readlines("input.txt")
# 
# # Initialize total priority to 0
# total_priority = 0
# 
# # Iterate through each rucksack
# rucksacks.each do |rucksack|
#   # Get the first and second compartment of the rucksack
#   comp1 = rucksack[0, rucksack.length/2]
#   comp2 = rucksack[rucksack.length/2, rucksack.length]
# 
#   # Check if any characters appear in both compartments
#   common_chars = comp1.chars.select { |c| comp2.include?(c) }
# 
#   # If any characters are common, add their priority to the total
#   if common_chars.any?
#     # Convert the character to lowercase
#     char = common_chars.first.downcase
# 
#     # Determine the priority using the formula
#     # Lowercase letters are 1-26, uppercase are 27-52
#     priority = char.ord - 'a'.ord + 1
#     if char.ord <= 'Z'.ord
#       priority += 26
#     end
# 
#     total_priority += priority
#   end
# end
# 
# # Print the total priority
# puts total_priority

# # TRY 3
# 
# # Read rucksack list from input
# rucksacks = File.readlines("input.txt")
# 
# # Initialize total priority to 0
# total_priority = 0
# 
# # Iterate through each rucksack
# rucksacks.each do |rucksack|
#   # Get the first and second compartment of the rucksack
#   comp1 = rucksack[0, rucksack.length/2]
#   comp2 = rucksack[rucksack.length/2, rucksack.length]
# 
#   # Check if any characters appear in both compartments
#   common_chars = comp1.chars.select { |c| comp2.include?(c) }
# 
#   # If any characters are common, add their priority to the total
#   if common_chars.any?
#     # Check if the character is uppercase or lowercase
#     char = common_chars.first
#     if char.upcase == char
#       char = char.downcase
#     end
# 
#     # Determine the priority using the formula
#     # Lowercase letters are 1-26, uppercase are 27-52
#     priority = char.ord - 'a'.ord + 1
#     if char.ord <= 'Z'.ord
#       priority += 26
#     end
# 
#     total_priority += priority
#   end
# end
# 
# # Print the total priority
# puts total_priority

# # TRY 4
# # Read rucksack list from input
# rucksacks = File.readlines("input.txt")
# 
# # Initialize total priority to 0
# total_priority = 0
# 
# # Iterate through each rucksack
# rucksacks.each do |rucksack|
#   # Get the first and second compartment of the rucksack
#   comp1 = rucksack[0, rucksack.length/2]
#   comp2 = rucksack[rucksack.length/2, rucksack.length]
# 
#   # Check if any characters appear in both compartments
#   common_chars = comp1.chars.select { |c| comp2.include?(c) }
# 
#   # If any characters are common, add their priority to the total
#   if common_chars.any?
#     # Check if the character is uppercase or lowercase
#     char = common_chars.first
#     if char.upcase == char
#       # Use the lowercase version of the character to determine the priority
#       char = char.downcase
#     end
# 
#     # Determine the priority using the formula
#     # Lowercase letters are 1-26, uppercase are 27-52
#     priority = char.ord - 'a'.ord + 1
#     if char.ord <= 'Z'.ord
#       priority += 26
#     end
# 
#     total_priority += priority
#   end
# end
# 
# # Print the total priority
# puts total_priority

# Read rucksack list from input
rucksacks = File.readlines("input.txt")

# Initialize total priority to 0
total_priority = 0

# Iterate through each rucksack
rucksacks.each do |rucksack|
  # Get the first and second compartment of the rucksack
  comp1 = rucksack[0, rucksack.length/2]
  comp2 = rucksack[rucksack.length/2, rucksack.length]

  # Check if any characters appear in both compartments
  common_chars = comp1.chars.select { |c| comp2.include?(c) }

  # If any characters are common, add their priority to the total
  if common_chars.any?
    # Determine the priority using the formula
    # Lowercase letters are 1-26, uppercase are 27-52
    priority = common_chars.first.ord - 'a'.ord + 1
    if common_chars.first.ord <= 'Z'.ord
      priority += 26
    end

    total_priority += priority
  end
end

# Print the total priority
puts total_priority