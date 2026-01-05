homework = []
with open("/home/vladan/Projects/Advent_of_Code/2025/day_6/input.txt") as f:
    homework = f.read().split("\n")
    homework = homework[:len(homework)-1]

operations = homework[-1]
operations = operations.strip().split(" ")
operations = [op for op in operations if op]

numbers = homework[:-1]

longest = max(numbers, key=len)
op = 0

solutions = []
addition_result = 0
mult_result = 1
for i in range(len(longest)):
    formed_number = ""
    for num in numbers:
        if num[i] != " ":
            formed_number += num[i]
    if formed_number != "":
        if operations[op] == "+":
            addition_result += int(formed_number)
        else:
            mult_result *= int(formed_number)
    else:
        if operations[op] == "+":
            solutions.append(addition_result)
        else:
            solutions.append(mult_result)
        addition_result = 0
        mult_result = 1
        op += 1

if operations[op] == "+":
    solutions.append(addition_result)
else:
    solutions.append(mult_result)

print(sum(solutions))
