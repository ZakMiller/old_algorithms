import os


def get_stall(stalls):
    if len(stalls) == 1:
        return stalls[0]
    else:
        stalls_with_max_distance = []
        for stall in stalls:
            if len(stalls_with_max_distance) == 0 or stalls_with_max_distance[0][2] == stall[2]:
                stalls_with_max_distance.append(stall)
            elif stalls_with_max_distance[0][2] < stall[2]:
                stalls_with_max_distance = []
                stalls_with_max_distance.append(stall)

        return stalls_with_max_distance[0]


def get_ls(stalls, possible_spot):
    number_of_empty_stalls = 0
    for i in xrange(possible_spot - 1, 0, -1):
        if stalls[i] == False:
            number_of_empty_stalls += 1
        else:
            return number_of_empty_stalls
    return number_of_empty_stalls


def get_rs(stalls, possible_spot):
    number_of_empty_stalls = 0
    for i in xrange(possible_spot + 1, len(stalls)):
        if stalls[i] == False:
            number_of_empty_stalls += 1
        else:
            return number_of_empty_stalls
    return number_of_empty_stalls


def get_min_max(stalls, i):
    Ls = get_ls(stalls, i)
    Rs = get_rs(stalls, i)
    return (min(Ls, Rs), max(Ls, Rs))


def seat_everyone(stalls, number_people):
    choice_seats = []

    for i in xrange(0, len(stalls)):
        if stalls[i] == False:
            min_max = get_min_max(stalls, i)
            spot = (i, min_max[0], min_max[1])
            choice_seats.append(spot)
        # Get all seat values.

        # Sort values.
    choice_seats = sorted(
        choice_seats, key=lambda element: (-element[1], -element[2]))

    seat = choice_seats.pop(0)
    stalls[seat[0]] = True
    while len(choice_seats) > 0:
        top_seat = choice_seats[0]
        Ls = get_ls(stalls, top_seat[0])
        Rs = get_rs(stalls, top_seat[0])
        if min(Ls, Rs) == top_seat[1] and max(Ls, Rs) == top_seat[2]:
            seat = choice_seats.pop(0)
            stalls[seat[0]] = True
            last_seat = seat
        else:
            for i in xrange(0, len(choice_seats)):
                index = choice_seats[i][0]
                min_max = get_min_max(stalls, index)
                choice_seats[i] = (index, min_max[0], min_max[1])
            choice_seats = sorted(
                choice_seats, key=lambda element: (-element[1], -element[2]))
            seat = choice_seats.pop(0)
            stalls[seat[0]] = True
            last_seat = seat
    return (last_seat[1], last_seat[2])


with open('input.txt') as input_file:
    testCases = input_file.readlines()
numTestCases = testCases[0]
testCases.remove(numTestCases)
output_filename = 'output.txt'
if os.path.exists(output_filename):
    os.remove(output_filename)
min_max = []
for i in range(0, len(testCases)):
    x = testCases[i].split()
    number_stalls = x[0]
    number_people = x[1]
    stalls = [True] + int(number_stalls) * [False] + [True]
    minimum, maximum = seat_everyone(stalls, number_people)
    # for i in xrange(0, int(number_people)):
    #stalls, minimum, maximum = seat_person(stalls)
    min_max.append((minimum, maximum))

with open(output_filename, 'a') as output_file:
    for i in xrange(0, len(min_max)):
        output_file.write('Case #' + str(i + 1) + ': ' +
                          str(min_max[i][1]) + ' ' + str(min_max[i][0]) + '\n')
