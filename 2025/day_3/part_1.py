banks = []
with open("/home/vladan/Projects/Advent_of_Code/2025/day_3/input.txt") as f:
    banks = f.read().split("\n")
    banks = banks[:len(banks)-1]

sum = 0
for bank in banks:
    first_max = 0
    second_max = 1
    for i in range(len(bank)):
        if int(bank[i]) > int(bank[first_max]) and i != len(bank) - 1:
            first_max = i
            second_max = first_max + 1

    for i in range(first_max + 1, len(bank)):
        if int(bank[i]) > int(bank[second_max]):
            second_max = i

    batteries = bank[first_max] + bank[second_max]
    sum += int(batteries)

print(sum)
