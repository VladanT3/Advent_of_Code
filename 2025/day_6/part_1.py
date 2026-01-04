homework = []
with open("/home/vladan/Projects/Advent_of_Code/2025/day_6/input.txt") as f:
    homework = f.read().split("\n")
    homework = homework[:len(homework)-1]
    homework = [elem.strip() for elem in homework]
    homework = [elem.split(" ") for elem in homework]
    homework = [[elem for elem in sublist if elem] for sublist in homework]

operations = homework[-1]
numbers = homework[:-1]
solutions = []
for i in range(len(operations)):
    if operations[i] == "+":
        result = 0
        for num in numbers:
            result += int(num[i])
        solutions.append(result)
    else:
        result = 1
        for num in numbers:
            result *= int(num[i])
        solutions.append(result)

print(sum(solutions))
