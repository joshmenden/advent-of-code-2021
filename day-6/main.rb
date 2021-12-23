require "byebug"
fname = "input.txt"

fresh_fish = {
  8 => 0,
  7 => 0,
  6 => 0,
  5 => 0,
  4 => 0,
  3 => 0,
  2 => 0,
  1 => 0,
  0 => 0,
}

fishes = {
  8 => 0,
  7 => 0,
  6 => 0,
  5 => 0,
  4 => 0,
  3 => 0,
  2 => 0,
  1 => 0,
  0 => 0,
}

content = File.read(fname)
content.split(",").each do |age|
  fishes[age.to_i] += 1
end


for day in 1..256 do
  todays_fish = fresh_fish.dup
  fishes.each_key do |k|
    next if fishes[k] == 0
    if k > 0
      todays_fish[k - 1] += fishes[k]
    else
      todays_fish[6] += fishes[k]
      todays_fish[8] += fishes[k]
    end
  end

  fishes = todays_fish
end

puts fishes.values.reduce(:+)