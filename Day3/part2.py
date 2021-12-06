
fileName = "Day3/input"


filterList = {
    "0": [],
    "1": [],
}


def get_dominant(index, array, most_dominant=1):
    count = {
        "0": 0,
        "1": 0,
    }
    for line in array:
        count[line[index]] = count[line[index]] + 1

    if count['0'] == count['1']:
        return -1

    if most_dominant:
        return 0 if count['0'] > count['1'] else 1
    return 0 if count['0'] < count['1'] else 1

def find_dominant(array, dominant=1):
    filter_list = []
    tempArray = array
    for index in range(1, len(array[0])):
        filter_number = get_dominant(index, tempArray, most_dominant=dominant)
        if filter_number == -1:
            filter_number = dominant

        for line in tempArray:
            if int(line[index]) == filter_number:
                filter_list.append(line)

        if len(filter_list) == 1:
            return filter_list

        tempArray = filter_list[:]
        filter_list = []

    return filter_list



count = 0
with open(fileName, "r") as f:
    for line in f:
        line = line.rstrip()
        filterList[ line[0] ].append(line)

final1 = find_dominant(filterList["1"], dominant=1)[0]
final2 = find_dominant(filterList["0"], dominant=0)[0]
print("Final 1: ", final1, int(final1, 2))
print("Final 0: ", final2, int(final2, 2))
print("Total: ", int(final1, 2) * int(final2, 2))