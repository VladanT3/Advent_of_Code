instructions = []
instruction = []
number = 0
with open("/home/vladan/Projects/Advent_of_Code/2025/day_1/input.txt") as f:
    for ch in f.read():
        if ch == "\n":
            instruction.append(number)
            instructions.append(instruction)
            instruction = []
            number = 0
        elif ch == "L" or ch == "R":
            instruction.append(ch)
        else:
            number *= 10
            number += int(ch)

current = 50
counter = 0
for instruction in instructions:
    if instruction[0] == "L":
        current = (current - instruction[1]) % 100
    elif instruction[0] == "R":
        current = (current + instruction[1]) % 100
    if current == 0:
        counter += 1

print(counter)
