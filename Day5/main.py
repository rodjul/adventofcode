
def print_matrix(matrix):
    for i in range(0, len(matrix)):
        print(matrix[i])

values = [
    # "772,989 -> 772,394"
]

with open("input") as f:
    for row in f:
        values.append(row.rstrip())

values.sort()

print(values)

def parse_values(value):
    splitted = value.split(" -> ")
    x1, x2 = int(splitted[0].split(",")[0]), int(splitted[1].split(",")[0])
    y1, y2 = int(splitted[0].split(",")[1]), int(splitted[1].split(",")[1])
    return x1, y1, x2, y2

def find_length_matrix(values):
    maxX = 0
    maxY = 0

    for value in values:
        x1, y1, x2, y2 = parse_values(value)

        if x1 > maxX:
            maxX = x1
        if x2 > maxX:
            maxX = x2

        if y1 > maxY:
            maxY = y1
        if y2 > maxY:
            maxY = y2

    return maxX, maxY


length_x, length_y = find_length_matrix(values)

size = length_x
if length_y > size:
    size = length_y

# length_x, length_y = 988, 987

print("length_x: ", length_x)
print("length_y: ", length_y)
print("Matrix size: ", size)

matrix = [[0 for _ in range(0, size + 1)] for _ in range(0, size  + 1)]


def count_overlap(matrix):
    count = 0
    for x in range(0, len(matrix)):
        for y in range(0, len(matrix[0])):
            if not matrix[y][x] >=2:
                continue
            count += 1
    return count

def part1(values):
    for value in values:
        x1, y1, x2, y2 = parse_values(value)

        if x1 > x2:
            temp = x1
            x1 = x2
            x2 = temp

        if y1 > y2:
            temp = y1
            y1 = y2
            y2 = temp

        # only consider horizontal and vertical lines: lines where either x1 = x2 or y1 = y2
        if (y1 != y2 and x1!=x2):
            continue

        for x in range(x1, x2 + 1):
            for y in range(y1, y2 + 1):
                matrix[y][x] = matrix[y][x] + 1

    # print_matrix(matrix)

    print("Points: ", count_overlap(matrix))


def part2(values):
    for value in values:
        x1, y1, x2, y2 = parse_values(value)

        if (y1 != y2 and x1!=x2):
            numbersX = []
            numbersY = []
            if x1>x2:
                numbersX = [a for a in range(x1, x2-1, -1)]
            elif x1<x2:
                numbersX = [a for a in range(x1, x2+1)]

            if y1>y2:
                numbersY = [a for a in range(y1, y2-1, -1)]
            elif y1<y2:
                numbersY = [a for a in range(y1, y2+1)]

            i = 0
            while i<len(numbersX) and i<len(numbersY):
                x = numbersX[i]
                y = numbersY[i]
                matrix[y][x] = matrix[y][x] + 1
                i += 1
            continue

        if x1 > x2:
            temp = x1
            x1 = x2
            x2 = temp

        if y1 > y2:
            temp = y1
            y1 = y2
            y2 = temp

        for x in range(x1, x2 + 1):
            for y in range(y1, y2 + 1):
                try:
                    matrix[y][x] = matrix[y][x] + 1
                except Exception as e:
                    print("VALUE: ", value)
                    print("X: ", x)
                    print("Y: ", y)
                    raise e


part2(values)

print_matrix(matrix)

print("Points: ", count_overlap(matrix))
