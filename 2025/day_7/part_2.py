manifold = []
with open("/home/vladan/Projects/Advent_of_Code/2025/day_7/input.txt") as f:
    manifold = f.read().split("\n")[:-1]

def traverse_manifold(manifold, row, col):
    if row >= len(manifold) or col >= len(manifold[0]) or col < 0:
        return 1

    if manifold[row][col] == "^":
        return traverse_manifold(manifold, row, col-1) + traverse_manifold(manifold, row, col+1)
    else:
        return traverse_manifold(manifold, row+1, col)

print(traverse_manifold(manifold, 1, manifold[0].find("S")))
