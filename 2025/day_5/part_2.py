ranges = []
with open("/home/vladan/Projects/Advent_of_Code/2025/day_5/input.txt") as f:
    ids = f.read().split("\n\n")
    ids = ids[0].split("\n")
    for rng in ids:
        lower, higher = rng.split("-")
        ranges.append((int(lower), int(higher)))
    ranges.sort()

unique_ranges = []
for lower, higher in ranges:
    if not unique_ranges:
        unique_ranges.append([lower, higher])
    else:
        last_lower, last_higher = unique_ranges[-1]
        if lower <= last_higher + 1:
            unique_ranges[-1][1] = max(higher, last_higher)
        else:
            unique_ranges.append([lower, higher])

count = sum(higher - lower + 1 for lower, higher in unique_ranges)
print(count)
