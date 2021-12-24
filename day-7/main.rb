require "byebug"
fname = "input.txt"
content = File.read(fname)
positions = content.split(",").map { |pos| pos.to_i }

@costs_map = {}

def calculate_fuel(pos, home)
  abs = (pos - home).abs()
  return @costs_map[abs] if @costs_map.key?(abs)

  stepper = 1
  fuel = 0
  for pos in 1..(abs) do
    fuel += stepper
    stepper += 1
  end

  @costs_map[abs] = fuel

  return fuel
end


min = positions.min
max = positions.max

fuel_costs = []

for homebase in min..max do
  cost = positions.map { |pos| calculate_fuel(pos, homebase) }.reduce(:+)
  fuel_costs.push(cost)
end

puts fuel_costs.min