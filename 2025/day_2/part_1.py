ranges = []
with open("/home/vladan/Projects/Advent_of_Code/2025/day_2/input.txt") as f:
    ranges = f.read().split(",")
    ranges[len(ranges) - 1] = ranges[len(ranges) - 1].replace("\n", "")

sum = 0
for id_range in ranges:
    lower = int(id_range.split("-")[0])
    higher = int(id_range.split("-")[1])
    for id in range(lower, higher+1):
        str_id = str(id)
        if len(str_id) % 2 != 0:
            continue
        left_half = str_id[:len(str_id) // 2]
        right_half = str_id[len(str_id) // 2:]
        if left_half == right_half:
            sum += id

print(sum)
