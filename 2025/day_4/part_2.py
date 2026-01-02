map = []
with open("/home/vladan/Projects/Advent_of_Code/2025/day_4/input.txt") as f:
    map = f.read().split("\n")
    map = map[:len(map)-1]

counter = 0
removed = 0
new_map = []
new_row = ""
while True:
    for i in range(len(map)):
        for j in range(len(map[0])):
            checker = 0
            next = map[i][j]
            if map[i][j] == "@":
                for k in range(-1, 2):
                    for l in range(-1, 2):
                        if k == 0 and l == 0:
                            continue

                        if (i == 0 and k == -1) or (j == 0 and l == -1):
                            continue

                        try:
                            if map[i+k][j+l] == "@":
                                checker += 1
                        except IndexError:
                            continue
                if checker < 4:
                    removed += 1
                    next = "."
            new_row += next
        new_map.append(new_row)
        new_row = ""
    if removed == 0:
        break
    counter += removed
    removed = 0
    map = new_map
    new_map = []
    
print(counter)
