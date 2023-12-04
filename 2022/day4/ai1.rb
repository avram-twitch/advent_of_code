def count_overlapping_pairs(pairs)
  # Parse each range in the pairs, and create a list of pairs of integers representing the start and end of the range
  ranges = pairs.map do |pair|
    pair.split(",").map do |range|
      range.split("-").map(&:to_i)
    end
  end

  # Count the number of pairs where one range fully contains the other
  count = 0
  ranges.each do |range1, range2|
    count += 1 if (range1[0] <= range2[0] && range1[1] >= range2[1]) || (range1[0] >= range2[0] && range1[1] <= range2[1])
  end
  count
end

# Test the function
# pairs = ["2-4,6-8", "2-3,4-5", "5-7,7-9", "2-8,3-7", "6-6,4-6", "2-6,4-8"]
pairs = File.open("input.txt").read.split("\n")
p count_overlapping_pairs(pairs) # => 2
