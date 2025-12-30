ranges = []
with open("/home/vladan/Projects/Advent_of_Code/2025/day_2/input.txt") as f:
    ranges = f.read().split(",")
    ranges[len(ranges) - 1] = ranges[len(ranges) - 1].replace("\n", "")

sum = 0
for id_range in ranges:
    lower = int(id_range.split("-")[0])
    higher = int(id_range.split("-")[1])
    for id_num in range(lower, higher+1):
        id = str(id_num)
        splits = 2
        while splits <= len(id):
            part_size = len(id) // splits
            split_id = [id[i:i + part_size] for i in range(0, len(id), part_size)]
            if all(split == split_id[0] for split in split_id):
                sum += id_num
                break
            splits += 1

print(sum)
