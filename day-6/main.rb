require "byebug"
fname = "example.txt"


class LanternFish
  attr_accessor :timer

  def initialize(timer:)
    @timer = timer
  end
end

fishes = []

content = File.read(fname)
content.split(",").each do |age|
  fishes.push(LanternFish.new(timer: age.to_i))
end

for day in 1..256 do
  # puts fishes.map(&:timer).join(" ")

  new_fish = []
  fishes.each do |fish|
    if fish.timer == 0
      new_fish.push(LanternFish.new(timer: 8))
      fish.timer = 6
    else
      fish.timer -= 1
    end
  end

  fishes = fishes.concat(new_fish)
end

puts fishes.count