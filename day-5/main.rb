require "byebug"
fname = "input.txt"

content = File.read(fname)

points = []

content.split("\n").each do |line|
  data = line.split(" -> ")
  points.push({
    start: {
      x: data[0].split(",")[0].to_i,
      y: data[0].split(",")[1].to_i
    },
    end: {
      x: data[1].split(",")[0].to_i,
      y: data[1].split(",")[1].to_i
    }
  })
end

covered_points = {}

points.each do |point|
  if point[:start][:y] != point[:end][:y] && point[:start][:x] != point[:end][:x] # diagonal line
    current_point = point[:start]
    x_dir_pos = point[:end][:x] > point[:start][:x]
    y_dir_pos = point[:end][:y] > point[:start][:y]

    loop do
      if covered_points.key?([current_point[:x], current_point[:y]])
        covered_points[[current_point[:x], current_point[:y]]] += 1
      else
        covered_points[[current_point[:x], current_point[:y]]] = 1
      end

      break if current_point == point[:end]

      if x_dir_pos
        current_point[:x] += 1
      else
        current_point[:x] -= 1
      end

      if y_dir_pos
        current_point[:y] += 1
      else
        current_point[:y] -= 1
      end
    end
  elsif point[:start][:x] == point[:end][:x] # vertical line
    y_min = [point[:start][:y], point[:end][:y]].min
    y_max = [point[:start][:y], point[:end][:y]].max
    x_const = point[:start][:x]
    counter = y_min
    while counter <= y_max do
      point = [x_const, counter]
      if covered_points.key?(point)
        covered_points[point] += 1
      else
        covered_points[point] = 1
      end

      counter += 1
    end
  else # horizontal line
    x_min = [point[:start][:x], point[:end][:x]].min
    x_max = [point[:start][:x], point[:end][:x]].max
    y_const = point[:start][:y]
    counter = x_min
    while counter <= x_max do
      point = [counter, y_const]

      if covered_points.key?(point)
        covered_points[point] += 1
      else
        covered_points[point] = 1
      end

      counter += 1
    end
  end
end

overlap = 0

covered_points.each_value do |v|
  overlap += 1 if v > 1
end

puts overlap
