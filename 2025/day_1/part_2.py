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
    direction, steps = instruction
    if direction == "L":
        until_zero = current
        if until_zero == 0:
            until_zero = 100
        if steps >= until_zero:
            counter += (steps - until_zero) // 100 + 1
        current = (current - steps) % 100
    elif direction == "R":
        until_zero = 100 - current
        if steps >= until_zero:
            counter += (steps - until_zero) // 100 + 1
        current = (current + steps) % 100

print(counter)
