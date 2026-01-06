manifold = []
with open("/home/vladan/Projects/Advent_of_Code/2025/day_7/input.txt") as f:
    manifold = f.read().split("\n")[:-1]

col = manifold[0].find("S")
count = 0
visited = {}

def traverse_manifold(manifold, row, col):
    if row >= len(manifold) or col >= len(manifold[0]) or col < 0:
        return

    if manifold[row][col] == "^":
        if not visited.get((row, col)):
            visited[(row, col)] = True
            global count
            count += 1
            traverse_manifold(manifold, row, col-1)
            traverse_manifold(manifold, row, col+1)
        else:
            return
    else:
        traverse_manifold(manifold, row+1, col)

traverse_manifold(manifold, 1, col)
print(count)
