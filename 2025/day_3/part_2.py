banks = []
with open("/home/vladan/Projects/Advent_of_Code/2025/day_3/input.txt") as f:
    banks = f.read().split("\n")
    banks = banks[:len(banks)-1]

sum = 0
for bank in banks:
    curr_max = 0
    batteries = ""
    for i in range(12, 0, -1):
        for j in range(curr_max, len(bank)):
            if int(bank[j]) > int(bank[curr_max]) and len(bank) - j >= i:
                curr_max = j
        batteries += bank[curr_max]
        curr_max += 1
    sum += int(batteries)

print(sum)
