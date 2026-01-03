ranges = []
with open("/home/vladan/Projects/Advent_of_Code/2025/day_5/input.txt") as f:
    ids = f.read().split("\n\n")
    ranges = ids[0].split("\n")

unique_ranges = []
while True:
    curr_lower, curr_higher = map(int, ranges[0].split("-"))
    new_ranges = ranges.copy()

    for id_range in ranges:
        lower, higher = map(int, id_range.split("-"))
        if curr_lower >= lower and curr_lower <= higher:
            curr_lower = lower
            try:
                new_ranges.remove(id_range)
            except ValueError:
                pass
        if curr_higher >= lower and curr_higher <= higher:
            curr_higher = higher
            try:
                new_ranges.remove(id_range)
            except ValueError:
                pass

    unique_ranges.append([curr_lower, curr_higher])

    ranges = new_ranges.copy()
    if len(ranges) == 0:
        break

print(unique_ranges)
