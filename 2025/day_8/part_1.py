from math import sqrt, pow

coords = []
with open("/home/vladan/Projects/Advent_of_Code/2025/day_8/input.txt") as f:
    coords = f.read().split("\n")[:-1]

distances = {}
for i in range(len(coords)):
    for j in range(len(coords)):
        if coords[i] == coords[j]:
            continue
        
        coords_1 = coords[i].split(",")
        coords_2 = coords[j].split(",")
        distance = 0
        for k in range(len(coords_1)):
            distance += pow((int(coords_1[k]) - int(coords_2[k])), 2)
        distance = sqrt(distance)

        if distances.get((coords[j], coords[i])) is not None:
            continue
        else:
            distances[(coords[i], coords[j])] = distance

distances = {key: value for key, value in sorted(distances.items(), key=lambda item: item[1])}
for key, val in distances.items():
    print(f"{key}: {val}")
print()

circuits = []
i = 1
for key in distances.keys():
    if i > 10:
        break
    if not circuits:
        circuit = set()
        circuit.update(key)
        circuits.append(circuit)
        continue

    added_to_circuit = False
    for circuit in circuits:
        for coord in circuit:
            if key[0] == coord:
                circuit.add(key[1])
                added_to_circuit = True
                break
            if key[1] == coord:
                circuit.add(key[0])
                added_to_circuit = True
                break
        if added_to_circuit:
            break
    if not added_to_circuit:
        new_circuit = set()
        new_circuit.update(key)
        circuits.append(new_circuit)

    i += 1

circuits.sort(key=len, reverse=True)
for circuit in circuits:
    print(circuit)

result = 1
for i in range(3):
    result *= len(circuits[i])

print()
print(result)
