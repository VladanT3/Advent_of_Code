fresh = []
available = []
with open("/home/vladan/Projects/Advent_of_Code/2025/day_5/input.txt") as f:
    ids = f.read().split("\n\n")
    fresh = ids[0].split("\n")
    available = ids[1].split("\n")
    available = available[:len(available)-1]

counter = 0
for id in available:
    for id_range in fresh:
        lower, higher = id_range.split("-")
        if int(id) >= int(lower) and int(id) <= int(higher):
            counter += 1
            break

print(counter)
