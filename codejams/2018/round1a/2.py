import sys


def get_optimal_cashier(cashiers, robots):
    optimal_cashier = (1000, 1000, 1000, 1000)
    min_time = 1000000
    startup_cost = 0
    for cashier in cashiers:
        bits_available = cashier[0]
        cost_per = cashier[1]
        cost_flat = cashier[2]
        used = cashier[3]
        cost = cashier[4]
        time = (used + 1) * cost_per + cost_flat

        valid_time = time < min_time or (time == min_time and cashier[3] == 0 and optimal_cashier[3] != 0) or (
            time == min_time and cashier[4] < startup_cost)

        if valid_time and (robots > 0 or used > 0):
            min_time = min(time, min_time)
            startup_cost = cashier[4]
            optimal_cashier = cashier
    return optimal_cashier, min_time


def problem_two(robots, bits, cashiers_count, cashiers_list_strings):
    cashiers = []
    times = []
    for cashier in cashiers_list_strings:
        cashier_pieces = cashier.split(" ")
        cashiers.append(
            (int(cashier_pieces[0]), int(cashier_pieces[1]), int(cashier_pieces[2]), 0, 0))
    while bits > 0:

        cashier, time = get_optimal_cashier(cashiers, robots)
        bits_available = cashier[0]
        cost_per = cashier[1]
        cost_flat = cashier[2]
        used = cashier[3]
        cost = cashier[4]
        if used == 0:
            robots -= 1
        bits_used = 1
        bits -= bits_used
        new_used = used + 1
        new_cost = new_used*cost_per + cost_flat
        if cashier[3] + 1 < cashier[0]:

            cashiers.append(
                (bits_available, cost_per, cost_flat, new_used, new_cost))
        else:
            times.append(new_cost)

        cashiers.remove(cashier)
    used_times = [x[4] for x in cashiers]
    return max(times + used_times, key=lambda item: item)


problem_number = 1
infile = sys.stdin
next(infile)
for line in infile:
    problem_pieces = line.split(" ")
    robots = int(problem_pieces[0])
    bits = int(problem_pieces[1])
    cashiers_count = int(problem_pieces[2])
    cashiers = []
    for i in xrange(0, cashiers_count):
        cashiers.append(next(infile).strip())
    print "Case #" + str(problem_number) + ": " + \
        str(problem_two(robots, bits, cashiers_count, cashiers))
    problem_number += 1
