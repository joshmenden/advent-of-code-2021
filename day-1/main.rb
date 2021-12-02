arr = []

depth_a = nil
depth_b = nil
depth_c = nil

sum_a = nil
sum_b = nil

increases = 0

arr.each_with_index do |depth, index|
  if index == 0 || index == 1
    next
  end


  depth_a = arr[index - 2]
  depth_b = arr[index - 1]
  depth_c = arr[index]
  sum = depth_a + depth_b + depth_c

  if sum_a == nil
    sum_a = sum
    sum_b = sum
    next
  end

  sum_a = sum_b
  sum_b = sum

  puts "#{sum_a} > #{sum_b} ?"
  if sum_b > sum_a
    increases = increases + 1
  end
end

puts "increases #{increases}"